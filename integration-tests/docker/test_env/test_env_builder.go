package test_env

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/seth"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	"github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/logstream"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/testreporters"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/osutil"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
)

type CleanUpType string

const (
	CleanUpTypeNone     CleanUpType = "none"
	CleanUpTypeStandard CleanUpType = "standard"
	CleanUpTypeCustom   CleanUpType = "custom"
)

type WithoutOldEVMClient struct{}

func (WithoutOldEVMClient) WaitForEvents() error {
	return nil
}

func (WithoutOldEVMClient) ParallelTransactions(_ bool) {}

func (WithoutOldEVMClient) GetChainID() *big.Int {
	return big.NewInt(0)
}

func (WithoutOldEVMClient) ReturnFunds(_ *ecdsa.PrivateKey) error {
	return nil
}

func (WithoutOldEVMClient) EstimateGas(_ ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (WithoutOldEVMClient) Fund(_ string, _ *big.Float, _ uint64) error { return nil }

func (WithoutOldEVMClient) Close() error {
	return nil
}

type OldEVMClient interface {
	WaitForEvents() error
	ParallelTransactions(bool)
	GetChainID() *big.Int
	ReturnFunds(*ecdsa.PrivateKey) error
	EstimateGas(msg ethereum.CallMsg) (uint64, error)
	Fund(string, *big.Float, uint64) error
	Close() error
}

type ChainlinkNodeLogScannerSettings struct {
	FailingLogLevel zapcore.Level
	Threshold       uint
	AllowedMessages []testreporters.AllowedLogMessage
}

type CLTestEnvBuilder[OldEVMClientType OldEVMClient] struct {
	hasLogStream                    bool
	hasKillgrave                    bool
	hasSeth                         bool
	hasEVMClient                    bool
	clNodeConfig                    *chainlink.Config
	secretsConfig                   string
	clNodesCount                    int
	clNodesOpts                     []func(*ClNode)
	customNodeCsaKeys               []string
	defaultNodeCsaKeys              []string
	l                               zerolog.Logger
	t                               *testing.T
	te                              *CLClusterTestEnv[OldEVMClientType]
	isEVM                           bool
	cleanUpType                     CleanUpType
	cleanUpCustomFn                 func()
	evmNetworkOption                []EVMNetworkOption
	privateEthereumNetworks         []*ctf_config.EthereumNetworkConfig
	testConfig                      ctf_config.GlobalTestConfig
	chainlinkNodeLogScannerSettings *ChainlinkNodeLogScannerSettings
	oldEVMClientCreatorFn           func(blockchain.EVMNetwork, zerolog.Logger) (OldEVMClientType, error)

	/* funding */
	ETHFunds *big.Float
}

var DefaultAllowedMessages = []testreporters.AllowedLogMessage{
	testreporters.NewAllowedLogMessage("Failed to get LINK balance", "Happens only when we deploy LINK token for test purposes. Harmless.", zapcore.ErrorLevel, testreporters.WarnAboutAllowedMsgs_No),
	testreporters.NewAllowedLogMessage("Error stopping job service", "It's a known issue with lifecycle. There's ongoing work that will fix it.", zapcore.DPanicLevel, testreporters.WarnAboutAllowedMsgs_No),
}

var DefaultChainlinkNodeLogScannerSettings = ChainlinkNodeLogScannerSettings{
	FailingLogLevel: zapcore.DPanicLevel,
	Threshold:       1, // we want to fail on the first concerning log
	AllowedMessages: DefaultAllowedMessages,
}

func GetDefaultChainlinkNodeLogScannerSettingsWithExtraAllowedMessages(extraAllowedMessages ...testreporters.AllowedLogMessage) ChainlinkNodeLogScannerSettings {
	allowedMessages := append(DefaultAllowedMessages, extraAllowedMessages...)
	return ChainlinkNodeLogScannerSettings{
		FailingLogLevel: DefaultChainlinkNodeLogScannerSettings.FailingLogLevel,
		Threshold:       DefaultChainlinkNodeLogScannerSettings.Threshold,
		AllowedMessages: allowedMessages,
	}
}

func NewCLTestEnvBuilder[OldEVMClientType OldEVMClient]() *CLTestEnvBuilder[OldEVMClientType] {
	return &CLTestEnvBuilder[OldEVMClientType]{
		l:                               log.Logger,
		hasLogStream:                    true,
		hasEVMClient:                    true,
		isEVM:                           true,
		chainlinkNodeLogScannerSettings: &DefaultChainlinkNodeLogScannerSettings,
	}
}

// WithTestEnv sets the test environment to use for the test.
// If nil, a new test environment is created.
// If not nil, the test environment is used as-is.
// If TEST_ENV_CONFIG_PATH is set, the test environment is created with the config at that path.
func (b *CLTestEnvBuilder[OldEVMClientType]) WithTestEnv(te *CLClusterTestEnv[OldEVMClientType]) (*CLTestEnvBuilder[OldEVMClientType], error) {
	envConfigPath, isSet := os.LookupEnv("TEST_ENV_CONFIG_PATH")
	var cfg *TestEnvConfig
	var err error
	if isSet {
		cfg, err = NewTestEnvConfigFromFile(envConfigPath)
		if err != nil {
			return nil, err
		}
	}

	if te != nil {
		b.te = te
	} else {
		b.te, err = NewTestEnv[OldEVMClientType]()
		if err != nil {
			return nil, err
		}
	}

	if cfg != nil {
		b.te = b.te.WithTestEnvConfig(cfg)
	}
	return b, nil
}

// WithTestLogger sets the test logger to use for the test.
// Useful for parallel tests so the logging will be separated correctly in the results views.
//
//revive:disable:confusing-naming
func (b *CLTestEnvBuilder[OldEVMClientType]) WithTestInstance(t *testing.T) *CLTestEnvBuilder[OldEVMClientType] {
	b.t = t
	b.l = logging.GetTestLogger(t)
	return b
}

// WithoutLogStream disables LogStream logging component
func (b *CLTestEnvBuilder[OldEVMClientType]) WithoutLogStream() *CLTestEnvBuilder[OldEVMClientType] {
	b.hasLogStream = false
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithoutChainlinkNodeLogScanner() *CLTestEnvBuilder[OldEVMClientType] {
	b.chainlinkNodeLogScannerSettings = &ChainlinkNodeLogScannerSettings{}
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithChainlinkNodeLogScanner(settings ChainlinkNodeLogScannerSettings) *CLTestEnvBuilder[OldEVMClientType] {
	b.chainlinkNodeLogScannerSettings = &settings
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithCLNodes(clNodesCount int) *CLTestEnvBuilder[OldEVMClientType] {
	b.clNodesCount = clNodesCount
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithTestConfig(cfg ctf_config.GlobalTestConfig) *CLTestEnvBuilder[OldEVMClientType] {
	b.testConfig = cfg
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithCLNodeOptions(opt ...ClNodeOption) *CLTestEnvBuilder[OldEVMClientType] {
	b.clNodesOpts = append(b.clNodesOpts, opt...)
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithFunding(eth *big.Float) *CLTestEnvBuilder[OldEVMClientType] {
	b.ETHFunds = eth
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithSeth() *CLTestEnvBuilder[OldEVMClientType] {
	b.hasSeth = true
	b.hasEVMClient = false
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithoutEvmClients() *CLTestEnvBuilder[OldEVMClientType] {
	b.hasSeth = false
	b.hasEVMClient = false
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithPrivateEthereumNetwork(en ctf_config.EthereumNetworkConfig) *CLTestEnvBuilder[OldEVMClientType] {
	b.privateEthereumNetworks = append(b.privateEthereumNetworks, &en)
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithPrivateEthereumNetworks(ens []*ctf_config.EthereumNetworkConfig) *CLTestEnvBuilder[OldEVMClientType] {
	b.privateEthereumNetworks = ens
	return b
}

// Deprecated: Use TOML instead
func (b *CLTestEnvBuilder[OldEVMClientType]) WithCLNodeConfig(cfg *chainlink.Config) *CLTestEnvBuilder[OldEVMClientType] {
	b.clNodeConfig = cfg
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithSecretsConfig(secrets string) *CLTestEnvBuilder[OldEVMClientType] {
	b.secretsConfig = secrets
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithMockAdapter() *CLTestEnvBuilder[OldEVMClientType] {
	b.hasKillgrave = true
	return b
}

// WithNonEVM sets the test environment to not use EVM when built.
func (b *CLTestEnvBuilder[OldEVMClientType]) WithNonEVM() *CLTestEnvBuilder[OldEVMClientType] {
	b.isEVM = false
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithStandardCleanup() *CLTestEnvBuilder[OldEVMClientType] {
	b.cleanUpType = CleanUpTypeStandard
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithoutCleanup() *CLTestEnvBuilder[OldEVMClientType] {
	b.cleanUpType = CleanUpTypeNone
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithCustomCleanup(customFn func()) *CLTestEnvBuilder[OldEVMClientType] {
	b.cleanUpType = CleanUpTypeCustom
	b.cleanUpCustomFn = customFn
	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) WithOldEVMClientCreatorFn(creatorFn func(blockchain.EVMNetwork, zerolog.Logger) (OldEVMClientType, error)) *CLTestEnvBuilder[OldEVMClientType] {
	b.oldEVMClientCreatorFn = creatorFn
	return b
}

type EVMNetworkOption = func(*blockchain.EVMNetwork) *blockchain.EVMNetwork

// WithEVMNetworkOptions sets the options for the EVM network. This is especially useful for simulated networks, which
// by usually use default options, so if we want to change any of them before the configuration is passed to evm client
// or Chainlnik node, we can do it here.
func (b *CLTestEnvBuilder[OldEVMClientType]) WithEVMNetworkOptions(opts ...EVMNetworkOption) *CLTestEnvBuilder[OldEVMClientType] {
	b.evmNetworkOption = make([]EVMNetworkOption, 0)
	b.evmNetworkOption = append(b.evmNetworkOption, opts...)

	return b
}

func (b *CLTestEnvBuilder[OldEVMClientType]) Build() (*CLClusterTestEnv[OldEVMClientType], error) {
	if b.testConfig == nil {
		return nil, fmt.Errorf("test config must be set")
	}

	if b.te == nil {
		var err error
		b, err = b.WithTestEnv(nil)
		if err != nil {
			return nil, err
		}
	}

	b.te.TestConfig = b.testConfig

	var err error
	if b.t != nil {
		b.te.WithTestInstance(b.t)
	}

	if b.hasLogStream {
		loggingConfig := b.testConfig.GetLoggingConfig()
		// we need to enable logging to file if we want to scan logs
		if b.chainlinkNodeLogScannerSettings != nil && !slices.Contains(loggingConfig.LogStream.LogTargets, string(logstream.File)) {
			b.l.Debug().Msg("Enabling logging to file in order to support Chainlink node log scanning")
			loggingConfig.LogStream.LogTargets = append(loggingConfig.LogStream.LogTargets, string(logstream.File))
		}
		b.te.LogStream, err = logstream.NewLogStream(b.te.t, b.testConfig.GetLoggingConfig())
		if err != nil {
			return nil, err
		}

		// this clean up has to be added as the FIRST one, because cleanup functions are executed in reverse order (LIFO)
		if b.t != nil && b.cleanUpType != CleanUpTypeNone {
			b.t.Cleanup(func() {
				b.l.Info().Msg("Shutting down LogStream")
				logPath, err := osutil.GetAbsoluteFolderPath("logs")
				if err != nil {
					b.l.Info().Str("Absolute path", logPath).Msg("LogStream logs folder location")
				}

				// flush logs when test failed or when we are explicitly told to collect logs
				flushLogStream := b.t.Failed() || *b.testConfig.GetLoggingConfig().TestLogCollect

				// run even if test has failed, as we might be able to catch additional problems without running the test again
				if b.chainlinkNodeLogScannerSettings != nil {
					logProcessor := logstream.NewLogProcessor[int](b.te.LogStream)

					processFn := func(log logstream.LogContent, count *int) error {
						countSoFar := count
						newCount, err := testreporters.ScanLogLine(b.l, string(log.Content), b.chainlinkNodeLogScannerSettings.FailingLogLevel, uint(*countSoFar), b.chainlinkNodeLogScannerSettings.Threshold, b.chainlinkNodeLogScannerSettings.AllowedMessages)
						if err != nil {
							return err
						}
						*count = int(newCount)
						return nil
					}

					// we cannot do parallel processing here, because ProcessContainerLogs() locks a mutex that controls whether
					// new logs can be added to the log stream, so parallel processing would get stuck on waiting for it to be unlocked
				LogScanningLoop:
					for i := 0; i < b.clNodesCount; i++ {
						// ignore count return, because we are only interested in the error
						_, err := logProcessor.ProcessContainerLogs(b.te.ClCluster.Nodes[i].ContainerName, processFn)
						if err != nil && !strings.Contains(err.Error(), testreporters.MultipleLogsAtLogLevelErr) && !strings.Contains(err.Error(), testreporters.OneLogAtLogLevelErr) {
							b.l.Error().Err(err).Msg("Error processing CL node logs")
							continue
						} else if err != nil && (strings.Contains(err.Error(), testreporters.MultipleLogsAtLogLevelErr) || strings.Contains(err.Error(), testreporters.OneLogAtLogLevelErr)) {
							flushLogStream = true
							b.t.Errorf("Found a concerning log in Chainklink Node logs: %v", err)
							break LogScanningLoop
						}
					}
					b.l.Info().Msg("Finished scanning Chainlink Node logs for concerning errors")
				}

				if flushLogStream {
					b.l.Info().Msg("Flushing LogStream logs")
					// we can't do much if this fails, so we just log the error in LogStream
					if err := b.te.LogStream.FlushAndShutdown(); err != nil {
						b.l.Error().Err(err).Msg("Error flushing and shutting down LogStream")
					}
					b.te.LogStream.PrintLogTargetsLocations()
					b.te.LogStream.SaveLogLocationInTestSummary()
				}
				b.l.Info().Msg("Finished shutting down LogStream")
			})
		} else {
			b.l.Warn().Msg("LogStream won't be cleaned up, because either test instance is not set or cleanup type is set to none")
		}
	}

	if b.hasKillgrave {
		if b.te.DockerNetwork == nil {
			return nil, fmt.Errorf("test environment builder failed: %w", fmt.Errorf("cannot start mock adapter without a network"))
		}

		b.te.MockAdapter = test_env.NewKillgrave([]string{b.te.DockerNetwork.Name}, "", test_env.WithLogStream(b.te.LogStream))

		err = b.te.StartMockAdapter()
		if err != nil {
			return nil, err
		}
	}

	if b.t != nil {
		b.te.WithTestInstance(b.t)
	}

	switch b.cleanUpType {
	case CleanUpTypeStandard:
		b.t.Cleanup(func() {
			// Cleanup test environment
			if err := b.te.Cleanup(CleanupOpts{TestName: b.t.Name()}); err != nil {
				b.l.Error().Err(err).Msg("Error cleaning up test environment")
			}
		})
	case CleanUpTypeCustom:
		b.t.Cleanup(b.cleanUpCustomFn)
	case CleanUpTypeNone:
		b.l.Warn().Msg("test environment won't be cleaned up")
	case "":
		return b.te, fmt.Errorf("test environment builder failed: %w", fmt.Errorf("explicit cleanup type must be set when building test environment"))
	}

	if b.te.LogStream == nil && b.chainlinkNodeLogScannerSettings != nil {
		log.Warn().Msg("Chainlink node log scanner settings provided, but LogStream is not enabled. Ignoring Chainlink node log scanner settings, as no logs will be available.")
	}

	// in this case we will use the builder only to start chains, not the cluster, because currently we support only 1 network config per cluster
	if len(b.privateEthereumNetworks) > 1 {
		b.te.rpcProviders = make(map[int64]*test_env.RpcProvider)
		b.te.EVMNetworks = make([]*blockchain.EVMNetwork, 0)
		b.te.evmClients = make(map[int64]OldEVMClientType)
		for _, en := range b.privateEthereumNetworks {
			en.DockerNetworkNames = []string{b.te.DockerNetwork.Name}
			networkConfig, rpcProvider, err := b.te.StartEthereumNetwork(en)
			if err != nil {
				return nil, err
			}

			if b.hasEVMClient {
				evmClient, err := b.oldEVMClientCreatorFn(networkConfig, b.l)
				if err != nil {
					return nil, err
				}
				b.te.evmClients[networkConfig.ChainID] = evmClient

				//if err != nil {
				//	return nil, err
				//}
				//evmClient, err := blockchain.NewEVMClientFromNetwork(networkConfig, b.l)
			}

			if b.hasSeth {
				sethClient, err := actions.GetChainClient(b.testConfig, networkConfig)
				if err != nil {
					return nil, err
				}

				b.te.sethClients[networkConfig.ChainID] = sethClient
			}

			b.te.rpcProviders[networkConfig.ChainID] = &rpcProvider
			b.te.EVMNetworks = append(b.te.EVMNetworks, &networkConfig)
		}

		dereferrencedEvms := make([]blockchain.EVMNetwork, 0)
		for _, en := range b.te.EVMNetworks {
			dereferrencedEvms = append(dereferrencedEvms, *en)
		}

		nodeConfigInToml := b.testConfig.GetNodeConfig()

		nodeConfig, _, err := node.BuildChainlinkNodeConfig(
			dereferrencedEvms,
			nodeConfigInToml.BaseConfigTOML,
			nodeConfigInToml.CommonChainConfigTOML,
			nodeConfigInToml.ChainConfigTOMLByChainID,
		)
		if err != nil {
			return nil, err
		}

		err = b.te.StartClCluster(nodeConfig, b.clNodesCount, b.secretsConfig, b.testConfig, b.clNodesOpts...)
		if err != nil {
			return nil, err
		}

		b.te.isSimulatedNetwork = true

		return b.te, nil
	}

	b.te.rpcProviders = make(map[int64]*test_env.RpcProvider)
	networkConfig := networks.MustGetSelectedNetworkConfig(b.testConfig.GetNetworkConfig())[0]
	// This has some hidden behavior so I'm not the biggest fan, but it matches expected behavior.
	// That is, when we specify we want to run on a live network in our config, we will run on the live network and not bother with a private network.
	// Even if we explicitly declare that we want to run on a private network in the test.
	// Keeping this a Kludge for now as SETH transition should change all of this anyway.
	if len(b.privateEthereumNetworks) == 1 {
		if networkConfig.Simulated {
			// TODO here we should save the ethereum network config to te.Cfg, but it doesn't exist at this point
			// in general it seems we have no methods for saving config to file and we only load it from file
			// but I don't know how that config file is to be created or whether anyone ever done that
			var rpcProvider test_env.RpcProvider
			b.privateEthereumNetworks[0].DockerNetworkNames = []string{b.te.DockerNetwork.Name}
			networkConfig, rpcProvider, err = b.te.StartEthereumNetwork(b.privateEthereumNetworks[0])
			if err != nil {
				return nil, err
			}
			b.te.rpcProviders[networkConfig.ChainID] = &rpcProvider
			b.te.PrivateEthereumConfigs = b.privateEthereumNetworks

			b.te.isSimulatedNetwork = true
		} else { // Only start and connect to a private network if we are using a private simulated network
			b.te.l.Warn().
				Str("Network", networkConfig.Name).
				Int64("Chain ID", networkConfig.ChainID).
				Msg("Private network config provided, but we are running on a live network. Ignoring private network config.")
			rpcProvider := test_env.NewRPCProvider(networkConfig.HTTPURLs, networkConfig.URLs, networkConfig.HTTPURLs, networkConfig.URLs)
			b.te.rpcProviders[networkConfig.ChainID] = &rpcProvider
			b.te.isSimulatedNetwork = false
		}
		b.te.EVMNetworks = append(b.te.EVMNetworks, &networkConfig)

	}

	if !b.hasSeth && !b.hasEVMClient {
		log.Debug().Msg("No EVM client or SETH client specified, not starting any clients")
	}

	if b.hasSeth && b.hasEVMClient {
		return nil, errors.New("you can't use both Seth and EMVClient at the same time")
	}

	if b.isEVM {
		if b.evmNetworkOption != nil && len(b.evmNetworkOption) > 0 {
			for _, fn := range b.evmNetworkOption {
				fn(&networkConfig)
			}
		}
		if b.hasEVMClient {
			evmClient, err := b.oldEVMClientCreatorFn(networkConfig, b.l)
			if err != nil {
				return nil, err
			}

			b.te.evmClients = make(map[int64]OldEVMClientType)
			b.te.evmClients[networkConfig.ChainID] = evmClient
		}

		if b.hasSeth {
			b.te.sethClients = make(map[int64]*seth.Client)
			sethClient, err := actions.GetChainClient(b.testConfig, networkConfig)
			if err != nil {
				return nil, err
			}

			b.te.sethClients[networkConfig.ChainID] = sethClient
		}
	}

	var nodeCsaKeys []string

	// Start Chainlink Nodes
	if b.clNodesCount > 0 {
		// needed for live networks
		if len(b.te.EVMNetworks) == 0 {
			b.te.EVMNetworks = append(b.te.EVMNetworks, &networkConfig)
		}

		dereferrencedEvms := make([]blockchain.EVMNetwork, 0)
		for _, en := range b.te.EVMNetworks {
			network := *en
			if en.Simulated {
				if rpcs, ok := b.te.rpcProviders[network.ChainID]; ok {
					network.HTTPURLs = rpcs.PrivateHttpUrls()
					network.URLs = rpcs.PrivateWsUrsl()
				} else {
					return nil, fmt.Errorf("rpc provider for chain %d not found", network.ChainID)
				}
			}
			dereferrencedEvms = append(dereferrencedEvms, network)
		}

		nodeConfigInToml := b.testConfig.GetNodeConfig()

		nodeConfig, _, err := node.BuildChainlinkNodeConfig(
			dereferrencedEvms,
			nodeConfigInToml.BaseConfigTOML,
			nodeConfigInToml.CommonChainConfigTOML,
			nodeConfigInToml.ChainConfigTOMLByChainID,
		)
		if err != nil {
			return nil, err
		}

		err = b.te.StartClCluster(nodeConfig, b.clNodesCount, b.secretsConfig, b.testConfig, b.clNodesOpts...)
		if err != nil {
			return nil, err
		}

		nodeCsaKeys, err = b.te.ClCluster.NodeCSAKeys()
		if err != nil {
			return nil, err
		}
		b.defaultNodeCsaKeys = nodeCsaKeys
	}

	if b.clNodesCount > 0 && b.ETHFunds != nil {
		if b.hasEVMClient {
			b.te.ParallelTransactions(true)
			defer b.te.ParallelTransactions(false)
			if err := b.te.FundChainlinkNodes(b.ETHFunds); err != nil {
				return nil, err
			}
		}
		if b.hasSeth {
			for _, sethClient := range b.te.sethClients {
				if err := actions.FundChainlinkNodesFromRootAddress(b.l, sethClient, contracts.ChainlinkClientToChainlinkNodeWithKeysAndAddress(b.te.ClCluster.NodeAPIs()), b.ETHFunds); err != nil {
					return nil, err
				}
			}
		}
	}

	var enDesc string
	if len(b.te.PrivateEthereumConfigs) > 0 {
		for _, en := range b.te.PrivateEthereumConfigs {
			enDesc += en.Describe()
		}
	} else {
		enDesc = "none"
	}

	b.l.Info().
		Str("privateEthereumNetwork", enDesc).
		Bool("hasKillgrave", b.hasKillgrave).
		Int("clNodesCount", b.clNodesCount).
		Strs("customNodeCsaKeys", b.customNodeCsaKeys).
		Strs("defaultNodeCsaKeys", b.defaultNodeCsaKeys).
		Msg("Building CL cluster test environment..")

	return b.te, nil
}

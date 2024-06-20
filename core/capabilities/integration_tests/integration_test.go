package integration_tests

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/target"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/trigger"
	remotetypes "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

type triggerFactory = func(t *testing.T) commoncap.TriggerCapability
type targetFactory = func(t *testing.T, reportsSink chan commoncap.CapabilityResponse) commoncap.TargetCapability

type consensusFactory = func(t *testing.T) commoncap.ConsensusCapability

func Test_HardcodedWorkflow_DonTopologies(t *testing.T) {
	ctx := testutils.Context(t)

	reportsSink := make(chan commoncap.CapabilityResponse, 1000)

	numWorkflowPeers := 3
	workflowDonF := uint8(1)
	numCapabilityPeers := 3
	capabilityDonF := uint8(1)

	workflowDonNodes := createDons(ctx, t, []triggerFactory{mockMercuryTrigger},
		[]targetFactory{mockEthereumTestnetSepoliaTarget},
		[]consensusFactory{mockConsensus}, numWorkflowPeers, workflowDonF, numCapabilityPeers, capabilityDonF,
		reportsSink)
	for _, node := range workflowDonNodes {
		AddWorkflowJob(t, node)
	}

	reportCount := 0
	for range reportsSink {
		reportCount++
		if reportCount == numWorkflowPeers {
			break
		}
	}

}

// TODO replace the capability setup in this function with syncer
func createDons(ctx context.Context, t *testing.T, triggerFactories []triggerFactory,
	targetFactories []targetFactory,
	consensusFactories []consensusFactory,
	numWorkflowPeers int, workflowDonF uint8,
	numCapabilityPeers int, capabilityDonF uint8,
	reportsSink chan commoncap.CapabilityResponse) []*cltest.TestApplication {
	lggr := logger.TestLogger(t)

	capabilityPeers := make([]p2ptypes.PeerID, numCapabilityPeers)
	for i := 0; i < numCapabilityPeers; i++ {
		capabilityPeerID := p2ptypes.PeerID{}
		require.NoError(t, capabilityPeerID.UnmarshalText([]byte(NewPeerID())))
		capabilityPeers[i] = capabilityPeerID
	}

	capabilityPeerID := p2ptypes.PeerID{}
	require.NoError(t, capabilityPeerID.UnmarshalText([]byte(NewPeerID())))

	capDonInfo := commoncap.DON{
		ID:      "capability-don",
		Members: capabilityPeers,
		F:       capabilityDonF,
	}

	workflowPeers := make([]p2ptypes.PeerID, numWorkflowPeers)
	for i := 0; i < numWorkflowPeers; i++ {
		workflowPeerID := p2ptypes.PeerID{}
		require.NoError(t, workflowPeerID.UnmarshalText([]byte(NewPeerID())))
		workflowPeers[i] = workflowPeerID
	}

	workflowDonInfo := commoncap.DON{
		Members: workflowPeers,
		ID:      "workflow-don",
		F:       workflowDonF,
	}

	broker := newTestAsyncMessageBroker(t, 1000)

	workflowDONs := map[string]commoncap.DON{
		workflowDonInfo.ID: workflowDonInfo,
	}

	for i := 0; i < numCapabilityPeers; i++ {
		capabilityPeer := capabilityPeers[i]
		capabilityDispatcher := broker.NewDispatcherForNode(capabilityPeer)

		for _, factory := range triggerFactories {
			trig := factory(t)
			capInfo, err := trig.Info(ctx)
			require.NoError(t, err)
			capInfo.DON = &capDonInfo

			cfg := &remotetypes.RemoteTriggerConfig{}
			cfg.ApplyDefaults()
			cfg.MinResponsesToAggregate = uint32(workflowDonInfo.F + 1)
			triggerPublisher := trigger.NewTriggerPublisher(cfg, trig, capInfo, capDonInfo, workflowDONs, capabilityDispatcher, lggr)
			servicetest.Run(t, triggerPublisher)
			broker.RegisterReceiverNode(capabilityPeer, capInfo.ID, capInfo.DON.ID, triggerPublisher)
		}

		for _, factory := range targetFactories {
			cb := factory(t, reportsSink)
			capInfo, err := cb.Info(ctx)
			require.NoError(t, err)
			capInfo.DON = &capDonInfo

			capabilityNode := target.NewServer(capabilityPeer, cb, capInfo, capDonInfo, workflowDONs, capabilityDispatcher,
				1*time.Minute, lggr)
			servicetest.Run(t, capabilityNode)
			broker.RegisterReceiverNode(capabilityPeer, capInfo.ID, capInfo.DON.ID, capabilityNode)
		}
	}

	workflowNodes := make([]*cltest.TestApplication, numWorkflowPeers)
	for i := 0; i < numWorkflowPeers; i++ {
		workflowPeerDispatcher := broker.NewDispatcherForNode(workflowPeers[i])
		capabilityRegistry := capabilities.NewRegistry(lggr)

		for _, triggerFactory := range triggerFactories {
			trig := triggerFactory(t)
			capInfo, err := trig.Info(ctx)
			require.NoError(t, err)
			capInfo.DON = &capDonInfo

			cfg := &remotetypes.RemoteTriggerConfig{}
			cfg.ApplyDefaults()
			cfg.MinResponsesToAggregate = uint32(capDonInfo.F + 1)

			triggerSubscriber := trigger.NewTriggerSubscriber(cfg, capInfo, capDonInfo, workflowDonInfo, workflowPeerDispatcher, nil, lggr)
			servicetest.Run(t, triggerSubscriber)
			broker.RegisterReceiverNode(workflowPeers[i], capInfo.ID, capInfo.DON.ID, triggerSubscriber)
			err = capabilityRegistry.Add(ctx, triggerSubscriber)
			require.NoError(t, err)
		}

		for _, targetFactory := range targetFactories {
			targ := targetFactory(t, reportsSink)
			capInfo, err := targ.Info(ctx)
			require.NoError(t, err)
			capInfo.DON = &capDonInfo

			targetClient := target.NewClient(capInfo, workflowDonInfo, workflowPeerDispatcher, 1*time.Minute, lggr)
			servicetest.Run(t, targetClient)
			broker.RegisterReceiverNode(workflowPeers[i], capInfo.ID, capInfo.DON.ID, targetClient)
			err = capabilityRegistry.Add(ctx, targetClient)
			require.NoError(t, err)
		}

		// Consensus capability is local to the workflow node
		for _, consensusFactory := range consensusFactories {
			consensus := consensusFactory(t)
			err := capabilityRegistry.Add(ctx, consensus)
			require.NoError(t, err)
		}

		nodeInfo := commoncap.Node{
			PeerID:         &workflowPeers[i],
			WorkflowDON:    workflowDonInfo,
			CapabilityDONs: []commoncap.DON{capDonInfo},
		}

		workflowNode := StartNewNode(t, capabilityRegistry, nodeInfo)
		require.NoError(t, workflowNode.Start(testutils.Context(t)))
		workflowNodes[i] = workflowNode
	}

	servicetest.Run(t, broker)

	return workflowNodes
}

func StartNewNode(
	t *testing.T, capabilityRegistry *capabilities.Registry, nodeInfo commoncap.Node) *cltest.TestApplication {
	config, _ := heavyweight.FullTestDBV2(t, func(c *chainlink.Config, s *chainlink.Secrets) {})

	app := cltest.NewApplicationWithConfig(t, config, capabilityRegistry, nodeInfo)

	return app
}

func NewPeerID() string {
	var privKey [32]byte
	_, err := rand.Read(privKey[:])
	if err != nil {
		panic(err)
	}

	peerID := append(libp2pMagic(), privKey[:]...)

	return base58.Encode(peerID[:])
}

func libp2pMagic() []byte {
	return []byte{0x00, 0x24, 0x08, 0x01, 0x12, 0x20}
}

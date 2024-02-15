package encoding

import (
	"net/http"

	ocr2keepers "github.com/smartcontractkit/chainlink-common/pkg/types/automation"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/automation_utils_2_1"
	iregistry21 "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/i_keeper_registry_master_wrapper_2_1"
)

type UpkeepFailureReason uint8
type PipelineExecutionState uint8

const (
	// upkeep failure onchain reasons
	UpkeepFailureReasonNone                    UpkeepFailureReason = 0
	UpkeepFailureReasonUpkeepCancelled         UpkeepFailureReason = 1
	UpkeepFailureReasonUpkeepPaused            UpkeepFailureReason = 2
	UpkeepFailureReasonTargetCheckReverted     UpkeepFailureReason = 3
	UpkeepFailureReasonUpkeepNotNeeded         UpkeepFailureReason = 4
	UpkeepFailureReasonPerformDataExceedsLimit UpkeepFailureReason = 5
	UpkeepFailureReasonInsufficientBalance     UpkeepFailureReason = 6
	UpkeepFailureReasonMercuryCallbackReverted UpkeepFailureReason = 7
	UpkeepFailureReasonRevertDataExceedsLimit  UpkeepFailureReason = 8
	UpkeepFailureReasonRegistryPaused          UpkeepFailureReason = 9
	// leaving a gap here for more onchain failure reasons in the future
	// upkeep failure offchain reasons
	UpkeepFailureReasonMercuryAccessNotAllowed UpkeepFailureReason = 32
	UpkeepFailureReasonTxHashNoLongerExists    UpkeepFailureReason = 33
	UpkeepFailureReasonInvalidRevertDataInput  UpkeepFailureReason = 34
	UpkeepFailureReasonSimulationFailed        UpkeepFailureReason = 35
	UpkeepFailureReasonTxHashReorged           UpkeepFailureReason = 36

	// pipeline execution error
	NoPipelineError        PipelineExecutionState = 0
	CheckBlockTooOld       PipelineExecutionState = 1
	CheckBlockInvalid      PipelineExecutionState = 2
	RpcFlakyFailure        PipelineExecutionState = 3
	MercuryFlakyFailure    PipelineExecutionState = 4
	PackUnpackDecodeFailed PipelineExecutionState = 5
	MercuryUnmarshalError  PipelineExecutionState = 6
	InvalidMercuryRequest  PipelineExecutionState = 7
	InvalidMercuryResponse PipelineExecutionState = 8 // this will only happen if Mercury server sends bad responses
	UpkeepNotAuthorized    PipelineExecutionState = 9
)

// ErrCode is used for invoking an error handler with a specific error code.
type ErrCode uint32

const (
	ErrCodeNil                 ErrCode = 0
	ErrCodePartialContent      ErrCode = 808206
	ErrCodeDataStreamsError    ErrCode = 808500
	ErrCodeBadRequest          ErrCode = 808400
	ErrCodeUnauthorized        ErrCode = 808401
	ErrCodeEncodingError       ErrCode = 808600
	ErrCodeStreamLookupTimeout ErrCode = 808603
)

func HttpToErrCode(statusCode int) ErrCode {
	switch statusCode {
	case http.StatusOK:
		return ErrCodeNil
	case http.StatusPartialContent:
		return ErrCodePartialContent
	case http.StatusBadRequest:
		return ErrCodeBadRequest
	case http.StatusUnauthorized:
		return ErrCodeUnauthorized
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return ErrCodeDataStreamsError
	default:
		return 0
	}
}

type UpkeepInfo = iregistry21.KeeperRegistryBase21UpkeepInfo

type Packer interface {
	UnpackCheckResult(payload ocr2keepers.UpkeepPayload, raw string) (ocr2keepers.CheckResult, error)
	UnpackPerformResult(raw string) (PipelineExecutionState, bool, error)
	UnpackLogTriggerConfig(raw []byte) (automation_utils_2_1.LogTriggerConfig, error)
	PackReport(report automation_utils_2_1.KeeperRegistryBase21Report) ([]byte, error)
	UnpackReport(raw []byte) (automation_utils_2_1.KeeperRegistryBase21Report, error)
}

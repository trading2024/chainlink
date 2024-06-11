package store

import (
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

// Note: this corresponds to a database enum; any update here
// needs to be reflected in an update to the `workflow_status` enum
const (
	StatusStarted            = "started"
	StatusErrored            = "errored"
	StatusTimeout            = "timeout"
	StatusCompleted          = "completed"
	StatusCompletedEarlyExit = "completed_early_exit"
)

type StepOutput struct {
	Err   error
	Value values.Value
}

type WorkflowExecutionStep struct {
	ExecutionID string
	Ref         string
	Status      string

	Inputs  *values.Map
	Outputs StepOutput

	UpdatedAt *time.Time
}

type WorkflowExecution struct {
	Steps       map[string]*WorkflowExecutionStep
	ExecutionID string
	WorkflowID  string

	Status     string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	FinishedAt *time.Time
}

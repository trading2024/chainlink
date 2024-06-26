package integration_tests

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/datastreams"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

const triggerID = "streams-trigger@1.0.0"

/*
var (
	feedOne   = "0x1111111111111111111100000000000000000000000000000000000000000000"
	feedTwo   = "0x2222222222222222222200000000000000000000000000000000000000000000"
	feedThree = "0x3333333333333333333300000000000000000000000000000000000000000000"
)*/

func mockMercuryTrigger(t *testing.T, reports []datastreams.FeedReport) capabilities.TriggerCapability {
	mt := &mockTriggerCapability{
		CapabilityInfo: capabilities.MustNewCapabilityInfo(
			triggerID,
			capabilities.CapabilityTypeTrigger,
			"issues a trigger when a mercury report is received.",
		),
		ch: make(chan capabilities.CapabilityResponse, 10),
	}
	/*	resp, err := values.NewMap(map[string]any{
			"123": decimal.NewFromFloat(1.00),
			"456": decimal.NewFromFloat(1.25),
			"789": decimal.NewFromFloat(1.50),
		})
		require.NoError(t, err)
	*/

	resp, err := wrapReports(reports, "1", 12, datastreams.SignersMetadata{})
	require.NoError(t, err)

	/*
		meta := values.NewString("some metadata")

		triggerEvent := capabilities.TriggerEvent{
			TriggerType: "mercury-report",
			ID:          "1",
			Timestamp:   strconv.FormatInt(time.Now().UnixMilli(), 10),
			Metadata:    meta,
			Payload:     resp,
		}

		eventVal, err := values.Wrap(triggerEvent)

		cr := capabilities.CapabilityResponse{
			Value: eventVal,
		} */
	mt.triggerEvent = &resp
	return mt
}

type mockTriggerCapability struct {
	capabilities.CapabilityInfo
	triggerEvent *capabilities.CapabilityResponse
	ch           chan capabilities.CapabilityResponse
}

var _ capabilities.TriggerCapability = (*mockTriggerCapability)(nil)

func (m *mockTriggerCapability) RegisterTrigger(ctx context.Context, req capabilities.CapabilityRequest) (<-chan capabilities.CapabilityResponse, error) {
	if m.triggerEvent != nil {
		m.ch <- *m.triggerEvent
	}
	return m.ch, nil
}

func (m *mockTriggerCapability) UnregisterTrigger(ctx context.Context, req capabilities.CapabilityRequest) error {
	return nil
}

func wrapReports(reportList []datastreams.FeedReport, eventID string, timestamp int64, meta datastreams.SignersMetadata) (capabilities.CapabilityResponse, error) {
	val, err := values.Wrap(reportList)
	if err != nil {
		return capabilities.CapabilityResponse{}, err
	}

	metaVal, err := values.Wrap(meta)
	if err != nil {
		return capabilities.CapabilityResponse{}, err
	}

	triggerEvent := capabilities.TriggerEvent{
		TriggerType: triggerID,
		ID:          eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
		Metadata:    metaVal,
		Payload:     val,
	}

	eventVal, err := values.Wrap(triggerEvent)
	if err != nil {
		return capabilities.CapabilityResponse{}, err
	}

	// Create a new CapabilityResponse with the MercuryTriggerEvent
	return capabilities.CapabilityResponse{
		Value: eventVal,
	}, nil
}

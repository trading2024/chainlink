// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: networkName, networkID, chainID, oracleName, sender, feedName, feedPath, symbol, contractType, contractStatus, contractAddress, feedID
func (_m *Metrics) Cleanup(networkName string, networkID string, chainID string, oracleName string, sender string, feedName string, feedPath string, symbol string, contractType string, contractStatus string, contractAddress string, feedID string) {
	_m.Called(networkName, networkID, chainID, oracleName, sender, feedName, feedPath, symbol, contractType, contractStatus, contractAddress, feedID)
}

// HTTPHandler provides a mock function with given fields:
func (_m *Metrics) HTTPHandler() http.Handler {
	ret := _m.Called()

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// IncOffchainAggregatorAnswersTotal provides a mock function with given fields: contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName
func (_m *Metrics) IncOffchainAggregatorAnswersTotal(contractAddress string, feedID string, chainID string, contractStatus string, contractType string, feedName string, feedPath string, networkID string, networkName string) {
	_m.Called(contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName)
}

// SetFeedContractMetadata provides a mock function with given fields: chainID, contractAddress, feedID, contractStatus, contractType, feedName, feedPath, networkID, networkName, symbol
func (_m *Metrics) SetFeedContractMetadata(chainID string, contractAddress string, feedID string, contractStatus string, contractType string, feedName string, feedPath string, networkID string, networkName string, symbol string) {
	_m.Called(chainID, contractAddress, feedID, contractStatus, contractType, feedName, feedPath, networkID, networkName, symbol)
}

// SetHeadTrackerCurrentHead provides a mock function with given fields: blockNumber, networkName, chainID, networkID
func (_m *Metrics) SetHeadTrackerCurrentHead(blockNumber uint64, networkName string, chainID string, networkID string) {
	_m.Called(blockNumber, networkName, chainID, networkID)
}

// SetNodeMetadata provides a mock function with given fields: chainID, networkID, networkName, oracleName, sender
func (_m *Metrics) SetNodeMetadata(chainID string, networkID string, networkName string, oracleName string, sender string) {
	_m.Called(chainID, networkID, networkName, oracleName, sender)
}

// SetOffchainAggregatorAnswerStalled provides a mock function with given fields: isSet, contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName
func (_m *Metrics) SetOffchainAggregatorAnswerStalled(isSet bool, contractAddress string, feedID string, chainID string, contractStatus string, contractType string, feedName string, feedPath string, networkID string, networkName string) {
	_m.Called(isSet, contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName)
}

// SetOffchainAggregatorAnswers provides a mock function with given fields: answer, contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName
func (_m *Metrics) SetOffchainAggregatorAnswers(answer *big.Int, contractAddress string, feedID string, chainID string, contractStatus string, contractType string, feedName string, feedPath string, networkID string, networkName string) {
	_m.Called(answer, contractAddress, feedID, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName)
}

// SetOffchainAggregatorSubmissionReceivedValues provides a mock function with given fields: value, contractAddress, feedID, sender, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName
func (_m *Metrics) SetOffchainAggregatorSubmissionReceivedValues(value *big.Int, contractAddress string, feedID string, sender string, chainID string, contractStatus string, contractType string, feedName string, feedPath string, networkID string, networkName string) {
	_m.Called(value, contractAddress, feedID, sender, chainID, contractStatus, contractType, feedName, feedPath, networkID, networkName)
}

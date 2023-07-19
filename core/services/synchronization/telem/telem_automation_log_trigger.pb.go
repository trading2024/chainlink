// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: telem_automation_log_trigger.proto

package telem

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MercuryLookup contains the metadata about a mercury request
type MercuryLookup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpkeepId    string   `protobuf:"bytes,1,opt,name=upkeep_id,json=upkeepId,proto3" json:"upkeep_id,omitempty"`
	BlockNumber uint32   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"` // block number provided by ocr2keepers plugin
	Timestamp   uint32   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                        // current timestamp
	TimeParam   uint32   `protobuf:"varint,4,opt,name=time_param,json=timeParam,proto3" json:"time_param,omitempty"`       // time param key is block number for v0.2 and timestamp for v0.3, time param is the corresponding value
	Feeds       []string `protobuf:"bytes,5,rep,name=feeds,proto3" json:"feeds,omitempty"`                                 // array of feed names
}

func (x *MercuryLookup) Reset() {
	*x = MercuryLookup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryLookup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryLookup) ProtoMessage() {}

func (x *MercuryLookup) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryLookup.ProtoReflect.Descriptor instead.
func (*MercuryLookup) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{0}
}

func (x *MercuryLookup) GetUpkeepId() string {
	if x != nil {
		return x.UpkeepId
	}
	return ""
}

func (x *MercuryLookup) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MercuryLookup) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MercuryLookup) GetTimeParam() uint32 {
	if x != nil {
		return x.TimeParam
	}
	return 0
}

func (x *MercuryLookup) GetFeeds() []string {
	if x != nil {
		return x.Feeds
	}
	return nil
}

// MercuryResponse contains the metadata about mercury response
type MercuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpkeepId        string   `protobuf:"bytes,1,opt,name=upkeep_id,json=upkeepId,proto3" json:"upkeep_id,omitempty"`
	BlockNumber     uint32   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`                      // block number provided by ocr2keepers plugin
	Timestamp       uint32   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                             // current timestamp
	Feeds           []string `protobuf:"bytes,4,rep,name=feeds,proto3" json:"feeds,omitempty"`                                                      // array of feed names
	HttpStatusCodes []uint32 `protobuf:"varint,5,rep,packed,name=http_status_codes,json=httpStatusCodes,proto3" json:"http_status_codes,omitempty"` // Mercury server response code
	Success         bool     `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty"`                                                 // True if all feeds gave successful response
	Retryable       bool     `protobuf:"varint,7,opt,name=retryable,proto3" json:"retryable,omitempty"`                                             // whether feedLookup should be retried if request fails
	FailureReason   uint32   `protobuf:"varint,8,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`                // failure enum defined in abi.go (UPKEEP_FAILURE_REASON_MERCURY_ACCESS_NOT_ALLOWED or some on chain reasons)
}

func (x *MercuryResponse) Reset() {
	*x = MercuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryResponse) ProtoMessage() {}

func (x *MercuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryResponse.ProtoReflect.Descriptor instead.
func (*MercuryResponse) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{1}
}

func (x *MercuryResponse) GetUpkeepId() string {
	if x != nil {
		return x.UpkeepId
	}
	return ""
}

func (x *MercuryResponse) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MercuryResponse) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MercuryResponse) GetFeeds() []string {
	if x != nil {
		return x.Feeds
	}
	return nil
}

func (x *MercuryResponse) GetHttpStatusCodes() []uint32 {
	if x != nil {
		return x.HttpStatusCodes
	}
	return nil
}

func (x *MercuryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MercuryResponse) GetRetryable() bool {
	if x != nil {
		return x.Retryable
	}
	return false
}

func (x *MercuryResponse) GetFailureReason() uint32 {
	if x != nil {
		return x.FailureReason
	}
	return 0
}

// MercuryCheckCallback contains whether customer's checkCallBack returns true with mercury data as input
type MercuryCheckCallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpkeepId      string `protobuf:"bytes,1,opt,name=upkeep_id,json=upkeepId,proto3" json:"upkeep_id,omitempty"`
	BlockNumber   uint32 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`       // block number provided by ocr2keepers plugin
	Timestamp     uint32 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                              // current timestamp
	FailureReason uint32 `protobuf:"varint,4,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"` // failure enum defined in abi.go (on chain reason)
	UpkeepNeeded  bool   `protobuf:"varint,5,opt,name=upkeep_needed,json=upkeepNeeded,proto3" json:"upkeep_needed,omitempty"`    // result of checkCallBack eth call, whether upkeep needs to be performed
}

func (x *MercuryCheckCallback) Reset() {
	*x = MercuryCheckCallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryCheckCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryCheckCallback) ProtoMessage() {}

func (x *MercuryCheckCallback) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryCheckCallback.ProtoReflect.Descriptor instead.
func (*MercuryCheckCallback) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{2}
}

func (x *MercuryCheckCallback) GetUpkeepId() string {
	if x != nil {
		return x.UpkeepId
	}
	return ""
}

func (x *MercuryCheckCallback) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *MercuryCheckCallback) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MercuryCheckCallback) GetFailureReason() uint32 {
	if x != nil {
		return x.FailureReason
	}
	return 0
}

func (x *MercuryCheckCallback) GetUpkeepNeeded() bool {
	if x != nil {
		return x.UpkeepNeeded
	}
	return false
}

// LogTrigger contains log trigger upkeep's information
type LogTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpkeepId       string `protobuf:"bytes,1,opt,name=upkeep_id,json=upkeepId,proto3" json:"upkeep_id,omitempty"`
	BlockNumber    uint32 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`            // block number provided by ocr2keepers plugin
	Timestamp      uint32 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                   // current timestamp
	LogBlockNumber uint32 `protobuf:"varint,4,opt,name=log_block_number,json=logBlockNumber,proto3" json:"log_block_number,omitempty"` // block number of log we are checking in pipeline
	LogBlockHash   string `protobuf:"bytes,5,opt,name=log_block_hash,json=logBlockHash,proto3" json:"log_block_hash,omitempty"`        // block has of log we are checking in pipeline
}

func (x *LogTrigger) Reset() {
	*x = LogTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogTrigger) ProtoMessage() {}

func (x *LogTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogTrigger.ProtoReflect.Descriptor instead.
func (*LogTrigger) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{3}
}

func (x *LogTrigger) GetUpkeepId() string {
	if x != nil {
		return x.UpkeepId
	}
	return ""
}

func (x *LogTrigger) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *LogTrigger) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogTrigger) GetLogBlockNumber() uint32 {
	if x != nil {
		return x.LogBlockNumber
	}
	return 0
}

func (x *LogTrigger) GetLogBlockHash() string {
	if x != nil {
		return x.LogBlockHash
	}
	return ""
}

// LogTriggerSuccess contains whether checkLog/checkUpkeep eth call returns true for a LogTriggered Upkeep
type LogTriggerSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpkeepId    string `protobuf:"bytes,1,opt,name=upkeep_id,json=upkeepId,proto3" json:"upkeep_id,omitempty"`
	BlockNumber uint32 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"` // block number provided by ocr2keepers plugin
	Timestamp   uint32 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                        // current timestamp
	Success     bool   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`                            // result of checkLog/checkUpkeep eth call, whether upkeep needs to be performed
}

func (x *LogTriggerSuccess) Reset() {
	*x = LogTriggerSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogTriggerSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogTriggerSuccess) ProtoMessage() {}

func (x *LogTriggerSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogTriggerSuccess.ProtoReflect.Descriptor instead.
func (*LogTriggerSuccess) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{4}
}

func (x *LogTriggerSuccess) GetUpkeepId() string {
	if x != nil {
		return x.UpkeepId
	}
	return ""
}

func (x *LogTriggerSuccess) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *LogTriggerSuccess) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogTriggerSuccess) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AutomationTelemWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*AutomationTelemWrapper_MercuryLookup
	//	*AutomationTelemWrapper_MercuryResponse
	//	*AutomationTelemWrapper_MercuryCheckcallback
	//	*AutomationTelemWrapper_LogTrigger
	//	*AutomationTelemWrapper_LogTriggerSuccess
	Msg isAutomationTelemWrapper_Msg `protobuf_oneof:"msg"`
}

func (x *AutomationTelemWrapper) Reset() {
	*x = AutomationTelemWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_automation_log_trigger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomationTelemWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomationTelemWrapper) ProtoMessage() {}

func (x *AutomationTelemWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_telem_automation_log_trigger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomationTelemWrapper.ProtoReflect.Descriptor instead.
func (*AutomationTelemWrapper) Descriptor() ([]byte, []int) {
	return file_telem_automation_log_trigger_proto_rawDescGZIP(), []int{5}
}

func (m *AutomationTelemWrapper) GetMsg() isAutomationTelemWrapper_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *AutomationTelemWrapper) GetMercuryLookup() *MercuryLookup {
	if x, ok := x.GetMsg().(*AutomationTelemWrapper_MercuryLookup); ok {
		return x.MercuryLookup
	}
	return nil
}

func (x *AutomationTelemWrapper) GetMercuryResponse() *MercuryResponse {
	if x, ok := x.GetMsg().(*AutomationTelemWrapper_MercuryResponse); ok {
		return x.MercuryResponse
	}
	return nil
}

func (x *AutomationTelemWrapper) GetMercuryCheckcallback() *MercuryCheckCallback {
	if x, ok := x.GetMsg().(*AutomationTelemWrapper_MercuryCheckcallback); ok {
		return x.MercuryCheckcallback
	}
	return nil
}

func (x *AutomationTelemWrapper) GetLogTrigger() *LogTrigger {
	if x, ok := x.GetMsg().(*AutomationTelemWrapper_LogTrigger); ok {
		return x.LogTrigger
	}
	return nil
}

func (x *AutomationTelemWrapper) GetLogTriggerSuccess() *LogTriggerSuccess {
	if x, ok := x.GetMsg().(*AutomationTelemWrapper_LogTriggerSuccess); ok {
		return x.LogTriggerSuccess
	}
	return nil
}

type isAutomationTelemWrapper_Msg interface {
	isAutomationTelemWrapper_Msg()
}

type AutomationTelemWrapper_MercuryLookup struct {
	MercuryLookup *MercuryLookup `protobuf:"bytes,1,opt,name=mercury_lookup,json=mercuryLookup,proto3,oneof"`
}

type AutomationTelemWrapper_MercuryResponse struct {
	MercuryResponse *MercuryResponse `protobuf:"bytes,2,opt,name=mercury_response,json=mercuryResponse,proto3,oneof"`
}

type AutomationTelemWrapper_MercuryCheckcallback struct {
	MercuryCheckcallback *MercuryCheckCallback `protobuf:"bytes,3,opt,name=mercury_checkcallback,json=mercuryCheckcallback,proto3,oneof"`
}

type AutomationTelemWrapper_LogTrigger struct {
	LogTrigger *LogTrigger `protobuf:"bytes,4,opt,name=log_trigger,json=logTrigger,proto3,oneof"`
}

type AutomationTelemWrapper_LogTriggerSuccess struct {
	LogTriggerSuccess *LogTriggerSuccess `protobuf:"bytes,5,opt,name=log_trigger_success,json=logTriggerSuccess,proto3,oneof"`
}

func (*AutomationTelemWrapper_MercuryLookup) isAutomationTelemWrapper_Msg() {}

func (*AutomationTelemWrapper_MercuryResponse) isAutomationTelemWrapper_Msg() {}

func (*AutomationTelemWrapper_MercuryCheckcallback) isAutomationTelemWrapper_Msg() {}

func (*AutomationTelemWrapper_LogTrigger) isAutomationTelemWrapper_Msg() {}

func (*AutomationTelemWrapper_LogTriggerSuccess) isAutomationTelemWrapper_Msg() {}

var File_telem_automation_log_trigger_proto protoreflect.FileDescriptor

var file_telem_automation_log_trigger_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x22, 0xa2, 0x01, 0x0a, 0x0d,
	0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x22, 0x90, 0x02, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x6e, 0x65, 0x65, 0x64,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70,
	0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6b, 0x65, 0x65, 0x70,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c,
	0x6f, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6b,
	0x65, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x6b, 0x65, 0x65, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0xf9, 0x02, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x43, 0x0a, 0x10, 0x6d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x4d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x0f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x15, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x14,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x34, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a,
	0x6c, 0x6f, 0x67, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x13, 0x6c, 0x6f,
	0x67, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e,
	0x4c, 0x6f, 0x67, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x48, 0x00, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x4e, 0x5a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telem_automation_log_trigger_proto_rawDescOnce sync.Once
	file_telem_automation_log_trigger_proto_rawDescData = file_telem_automation_log_trigger_proto_rawDesc
)

func file_telem_automation_log_trigger_proto_rawDescGZIP() []byte {
	file_telem_automation_log_trigger_proto_rawDescOnce.Do(func() {
		file_telem_automation_log_trigger_proto_rawDescData = protoimpl.X.CompressGZIP(file_telem_automation_log_trigger_proto_rawDescData)
	})
	return file_telem_automation_log_trigger_proto_rawDescData
}

var file_telem_automation_log_trigger_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_telem_automation_log_trigger_proto_goTypes = []interface{}{
	(*MercuryLookup)(nil),          // 0: telem.MercuryLookup
	(*MercuryResponse)(nil),        // 1: telem.MercuryResponse
	(*MercuryCheckCallback)(nil),   // 2: telem.MercuryCheckCallback
	(*LogTrigger)(nil),             // 3: telem.LogTrigger
	(*LogTriggerSuccess)(nil),      // 4: telem.LogTriggerSuccess
	(*AutomationTelemWrapper)(nil), // 5: telem.AutomationTelemWrapper
}
var file_telem_automation_log_trigger_proto_depIdxs = []int32{
	0, // 0: telem.AutomationTelemWrapper.mercury_lookup:type_name -> telem.MercuryLookup
	1, // 1: telem.AutomationTelemWrapper.mercury_response:type_name -> telem.MercuryResponse
	2, // 2: telem.AutomationTelemWrapper.mercury_checkcallback:type_name -> telem.MercuryCheckCallback
	3, // 3: telem.AutomationTelemWrapper.log_trigger:type_name -> telem.LogTrigger
	4, // 4: telem.AutomationTelemWrapper.log_trigger_success:type_name -> telem.LogTriggerSuccess
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_telem_automation_log_trigger_proto_init() }
func file_telem_automation_log_trigger_proto_init() {
	if File_telem_automation_log_trigger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telem_automation_log_trigger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryLookup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telem_automation_log_trigger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telem_automation_log_trigger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryCheckCallback); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telem_automation_log_trigger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogTrigger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telem_automation_log_trigger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogTriggerSuccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telem_automation_log_trigger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomationTelemWrapper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_telem_automation_log_trigger_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*AutomationTelemWrapper_MercuryLookup)(nil),
		(*AutomationTelemWrapper_MercuryResponse)(nil),
		(*AutomationTelemWrapper_MercuryCheckcallback)(nil),
		(*AutomationTelemWrapper_LogTrigger)(nil),
		(*AutomationTelemWrapper_LogTriggerSuccess)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telem_automation_log_trigger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telem_automation_log_trigger_proto_goTypes,
		DependencyIndexes: file_telem_automation_log_trigger_proto_depIdxs,
		MessageInfos:      file_telem_automation_log_trigger_proto_msgTypes,
	}.Build()
	File_telem_automation_log_trigger_proto = out.File
	file_telem_automation_log_trigger_proto_rawDesc = nil
	file_telem_automation_log_trigger_proto_goTypes = nil
	file_telem_automation_log_trigger_proto_depIdxs = nil
}

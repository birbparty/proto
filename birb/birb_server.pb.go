// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: birb/birb_server.proto

package birb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_StartGame EventType = 0
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "StartGame",
	}
	EventType_value = map[string]int32{
		"StartGame": 0,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_birb_birb_server_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_birb_birb_server_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{0}
}

type PlayerStatus int32

const (
	PlayerStatus_Idle    PlayerStatus = 0
	PlayerStatus_Waiting PlayerStatus = 1
	PlayerStatus_Active  PlayerStatus = 10
)

// Enum value maps for PlayerStatus.
var (
	PlayerStatus_name = map[int32]string{
		0:  "Idle",
		1:  "Waiting",
		10: "Active",
	}
	PlayerStatus_value = map[string]int32{
		"Idle":    0,
		"Waiting": 1,
		"Active":  10,
	}
)

func (x PlayerStatus) Enum() *PlayerStatus {
	p := new(PlayerStatus)
	*p = x
	return p
}

func (x PlayerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_birb_birb_server_proto_enumTypes[1].Descriptor()
}

func (PlayerStatus) Type() protoreflect.EnumType {
	return &file_birb_birb_server_proto_enumTypes[1]
}

func (x PlayerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerStatus.Descriptor instead.
func (PlayerStatus) EnumDescriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{1}
}

type SessionStatus int32

const (
	SessionStatus_Lobby                  SessionStatus = 0
	SessionStatus_WaitingOnPlayer        SessionStatus = 1
	SessionStatus_WaitingOnServerToStart SessionStatus = 2
	SessionStatus_Playing                SessionStatus = 10
	SessionStatus_PlayingWaitingOnPlayer SessionStatus = 11
	SessionStatus_PlayingWaitingOnServer SessionStatus = 12
	SessionStatus_Finished               SessionStatus = 50
)

// Enum value maps for SessionStatus.
var (
	SessionStatus_name = map[int32]string{
		0:  "Lobby",
		1:  "WaitingOnPlayer",
		2:  "WaitingOnServerToStart",
		10: "Playing",
		11: "PlayingWaitingOnPlayer",
		12: "PlayingWaitingOnServer",
		50: "Finished",
	}
	SessionStatus_value = map[string]int32{
		"Lobby":                  0,
		"WaitingOnPlayer":        1,
		"WaitingOnServerToStart": 2,
		"Playing":                10,
		"PlayingWaitingOnPlayer": 11,
		"PlayingWaitingOnServer": 12,
		"Finished":               50,
	}
)

func (x SessionStatus) Enum() *SessionStatus {
	p := new(SessionStatus)
	*p = x
	return p
}

func (x SessionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_birb_birb_server_proto_enumTypes[2].Descriptor()
}

func (SessionStatus) Type() protoreflect.EnumType {
	return &file_birb_birb_server_proto_enumTypes[2]
}

func (x SessionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionStatus.Descriptor instead.
func (SessionStatus) EnumDescriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{2}
}

type BirbEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	PlayerId  int64                  `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	EventType EventType              `protobuf:"varint,3,opt,name=eventType,proto3,enum=EventType" json:"eventType,omitempty"`
	EventJson string                 `protobuf:"bytes,4,opt,name=eventJson,proto3" json:"eventJson,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3,oneof" json:"createdAt,omitempty"`
}

func (x *BirbEvent) Reset() {
	*x = BirbEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BirbEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BirbEvent) ProtoMessage() {}

func (x *BirbEvent) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BirbEvent.ProtoReflect.Descriptor instead.
func (*BirbEvent) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{0}
}

func (x *BirbEvent) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BirbEvent) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *BirbEvent) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_StartGame
}

func (x *BirbEvent) GetEventJson() string {
	if x != nil {
		return x.EventJson
	}
	return ""
}

func (x *BirbEvent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type BirbPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId int64  `protobuf:"varint,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ai      bool   `protobuf:"varint,4,opt,name=ai,proto3" json:"ai,omitempty"`
	State   int64  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *BirbPlayer) Reset() {
	*x = BirbPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BirbPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BirbPlayer) ProtoMessage() {}

func (x *BirbPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BirbPlayer.ProtoReflect.Descriptor instead.
func (*BirbPlayer) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{1}
}

func (x *BirbPlayer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BirbPlayer) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *BirbPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BirbPlayer) GetAi() bool {
	if x != nil {
		return x.Ai
	}
	return false
}

func (x *BirbPlayer) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId      int64          `protobuf:"varint,1,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Seed         int64          `protobuf:"varint,2,opt,name=seed,proto3" json:"seed,omitempty"`
	Status       SessionStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=SessionStatus" json:"status,omitempty"`
	PlayerStates []*PlayerState `protobuf:"bytes,4,rep,name=playerStates,proto3" json:"playerStates,omitempty"`
	WaitingOnIds []int64        `protobuf:"varint,5,rep,packed,name=WaitingOnIds,proto3" json:"WaitingOnIds,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{2}
}

func (x *GameState) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *GameState) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *GameState) GetStatus() SessionStatus {
	if x != nil {
		return x.Status
	}
	return SessionStatus_Lobby
}

func (x *GameState) GetPlayerStates() []*PlayerState {
	if x != nil {
		return x.PlayerStates
	}
	return nil
}

func (x *GameState) GetWaitingOnIds() []int64 {
	if x != nil {
		return x.WaitingOnIds
	}
	return nil
}

type GameStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed    int64         `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Players []*BirbPlayer `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Events  []*BirbEvent  `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GameStateRequest) Reset() {
	*x = GameStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStateRequest) ProtoMessage() {}

func (x *GameStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStateRequest.ProtoReflect.Descriptor instead.
func (*GameStateRequest) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{3}
}

func (x *GameStateRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *GameStateRequest) GetPlayers() []*BirbPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameStateRequest) GetEvents() []*BirbEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type PlayerState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId         int64        `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Status           PlayerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=PlayerStatus" json:"status,omitempty"`
	StateJson        string       `protobuf:"bytes,3,opt,name=stateJson,proto3" json:"stateJson,omitempty"`
	AvailableOptions []*BirbEvent `protobuf:"bytes,4,rep,name=availableOptions,proto3" json:"availableOptions,omitempty"`
}

func (x *PlayerState) Reset() {
	*x = PlayerState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerState) ProtoMessage() {}

func (x *PlayerState) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerState.ProtoReflect.Descriptor instead.
func (*PlayerState) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerState) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *PlayerState) GetStatus() PlayerStatus {
	if x != nil {
		return x.Status
	}
	return PlayerStatus_Idle
}

func (x *PlayerState) GetStateJson() string {
	if x != nil {
		return x.StateJson
	}
	return ""
}

func (x *PlayerState) GetAvailableOptions() []*BirbEvent {
	if x != nil {
		return x.AvailableOptions
	}
	return nil
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed    int64         `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Event   *BirbEvent    `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Players []*BirbPlayer `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	Events  []*BirbEvent  `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *ValidateRequest) GetEvent() *BirbEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ValidateRequest) GetPlayers() []*BirbPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ValidateRequest) GetEvents() []*BirbEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type ValidationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameState *GameState `protobuf:"bytes,1,opt,name=gameState,proto3,oneof" json:"gameState,omitempty"`
	Error     *string    `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *ValidationResponse) Reset() {
	*x = ValidationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birb_birb_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResponse) ProtoMessage() {}

func (x *ValidationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_birb_birb_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResponse.ProtoReflect.Descriptor instead.
func (*ValidationResponse) Descriptor() ([]byte, []int) {
	return file_birb_birb_server_proto_rawDescGZIP(), []int{6}
}

func (x *ValidationResponse) GetGameState() *GameState {
	if x != nil {
		return x.GameState
	}
	return nil
}

func (x *ValidationResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_birb_birb_server_proto protoreflect.FileDescriptor

var file_birb_birb_server_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x69, 0x72, 0x62, 0x2f, 0x62, 0x69, 0x72, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x09, 0x42, 0x69,
	0x72, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e,
	0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x70, 0x0a, 0x0a, 0x42, 0x69, 0x72, 0x62, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x61, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x61, 0x69,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65,
	0x65, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0c, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x49, 0x64, 0x73,
	0x22, 0x71, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x69, 0x72, 0x62,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12,
	0x22, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x42, 0x69, 0x72, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4a,
	0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x42, 0x69, 0x72, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x69, 0x72, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x69, 0x72, 0x62, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x42, 0x69, 0x72, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x76, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x1a, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x10, 0x00, 0x2a, 0x31, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x64, 0x6c, 0x65, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x0a, 0x2a, 0x9e, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x4f, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x57, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x0b, 0x12,
	0x1a, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x4f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x32, 0x32, 0x68, 0x0a, 0x04, 0x42, 0x69, 0x72,
	0x62, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x31, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x69, 0x72, 0x62, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x69, 0x72, 0x62, 0xaa, 0x02, 0x0e, 0x42, 0x69, 0x72, 0x62, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_birb_birb_server_proto_rawDescOnce sync.Once
	file_birb_birb_server_proto_rawDescData = file_birb_birb_server_proto_rawDesc
)

func file_birb_birb_server_proto_rawDescGZIP() []byte {
	file_birb_birb_server_proto_rawDescOnce.Do(func() {
		file_birb_birb_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_birb_birb_server_proto_rawDescData)
	})
	return file_birb_birb_server_proto_rawDescData
}

var file_birb_birb_server_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_birb_birb_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_birb_birb_server_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: EventType
	(PlayerStatus)(0),             // 1: PlayerStatus
	(SessionStatus)(0),            // 2: SessionStatus
	(*BirbEvent)(nil),             // 3: BirbEvent
	(*BirbPlayer)(nil),            // 4: BirbPlayer
	(*GameState)(nil),             // 5: GameState
	(*GameStateRequest)(nil),      // 6: GameStateRequest
	(*PlayerState)(nil),           // 7: PlayerState
	(*ValidateRequest)(nil),       // 8: ValidateRequest
	(*ValidationResponse)(nil),    // 9: ValidationResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_birb_birb_server_proto_depIdxs = []int32{
	0,  // 0: BirbEvent.eventType:type_name -> EventType
	10, // 1: BirbEvent.createdAt:type_name -> google.protobuf.Timestamp
	2,  // 2: GameState.status:type_name -> SessionStatus
	7,  // 3: GameState.playerStates:type_name -> PlayerState
	4,  // 4: GameStateRequest.players:type_name -> BirbPlayer
	3,  // 5: GameStateRequest.events:type_name -> BirbEvent
	1,  // 6: PlayerState.status:type_name -> PlayerStatus
	3,  // 7: PlayerState.availableOptions:type_name -> BirbEvent
	3,  // 8: ValidateRequest.event:type_name -> BirbEvent
	4,  // 9: ValidateRequest.players:type_name -> BirbPlayer
	3,  // 10: ValidateRequest.events:type_name -> BirbEvent
	5,  // 11: ValidationResponse.gameState:type_name -> GameState
	6,  // 12: Birb.GetGameState:input_type -> GameStateRequest
	8,  // 13: Birb.Validate:input_type -> ValidateRequest
	5,  // 14: Birb.GetGameState:output_type -> GameState
	9,  // 15: Birb.Validate:output_type -> ValidationResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_birb_birb_server_proto_init() }
func file_birb_birb_server_proto_init() {
	if File_birb_birb_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_birb_birb_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BirbEvent); i {
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
		file_birb_birb_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BirbPlayer); i {
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
		file_birb_birb_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameState); i {
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
		file_birb_birb_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStateRequest); i {
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
		file_birb_birb_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerState); i {
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
		file_birb_birb_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_birb_birb_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResponse); i {
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
	file_birb_birb_server_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_birb_birb_server_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_birb_birb_server_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_birb_birb_server_proto_goTypes,
		DependencyIndexes: file_birb_birb_server_proto_depIdxs,
		EnumInfos:         file_birb_birb_server_proto_enumTypes,
		MessageInfos:      file_birb_birb_server_proto_msgTypes,
	}.Build()
	File_birb_birb_server_proto = out.File
	file_birb_birb_server_proto_rawDesc = nil
	file_birb_birb_server_proto_goTypes = nil
	file_birb_birb_server_proto_depIdxs = nil
}

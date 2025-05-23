// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: games/tictactoe.proto

package gen

import (
	user "github.com/bozoteam/roshan/adapter/grpc/gen/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Room struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId     string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Users         []*user.User           `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_games_tictactoe_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetUsers() []*user.User {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_games_tictactoe_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Room          *Room                  `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_games_tictactoe_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_games_tictactoe_proto protoreflect.FileDescriptor

const file_games_tictactoe_proto_rawDesc = "" +
	"\n" +
	"\x15games/tictactoe.proto\x12\froshan.games\x1a\x1cgoogle/api/annotations.proto\x1a\x0fuser/user.proto\"r\n" +
	"\x04Room\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x02 \x01(\tR\tcreatorId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12'\n" +
	"\x05users\x18\x04 \x03(\v2\x11.roshan.user.UserR\x05users\"'\n" +
	"\x11CreateRoomRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"<\n" +
	"\x12CreateRoomResponse\x12&\n" +
	"\x04room\x18\x01 \x01(\v2\x12.roshan.games.RoomR\x04room2}\n" +
	"\vChatService\x12n\n" +
	"\n" +
	"CreateRoom\x12\x1f.roshan.games.CreateRoomRequest\x1a .roshan.games.CreateRoomResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/api/v1/chat/roomsBAZ?github.com/bozoteam/roshan/adapter/grpc/gen/games/tictactoe;genb\x06proto3"

var (
	file_games_tictactoe_proto_rawDescOnce sync.Once
	file_games_tictactoe_proto_rawDescData []byte
)

func file_games_tictactoe_proto_rawDescGZIP() []byte {
	file_games_tictactoe_proto_rawDescOnce.Do(func() {
		file_games_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_games_tictactoe_proto_rawDesc), len(file_games_tictactoe_proto_rawDesc)))
	})
	return file_games_tictactoe_proto_rawDescData
}

var file_games_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_games_tictactoe_proto_goTypes = []any{
	(*Room)(nil),               // 0: roshan.games.Room
	(*CreateRoomRequest)(nil),  // 1: roshan.games.CreateRoomRequest
	(*CreateRoomResponse)(nil), // 2: roshan.games.CreateRoomResponse
	(*user.User)(nil),          // 3: roshan.user.User
}
var file_games_tictactoe_proto_depIdxs = []int32{
	3, // 0: roshan.games.Room.users:type_name -> roshan.user.User
	0, // 1: roshan.games.CreateRoomResponse.room:type_name -> roshan.games.Room
	1, // 2: roshan.games.ChatService.CreateRoom:input_type -> roshan.games.CreateRoomRequest
	2, // 3: roshan.games.ChatService.CreateRoom:output_type -> roshan.games.CreateRoomResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_games_tictactoe_proto_init() }
func file_games_tictactoe_proto_init() {
	if File_games_tictactoe_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_games_tictactoe_proto_rawDesc), len(file_games_tictactoe_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_games_tictactoe_proto_goTypes,
		DependencyIndexes: file_games_tictactoe_proto_depIdxs,
		MessageInfos:      file_games_tictactoe_proto_msgTypes,
	}.Build()
	File_games_tictactoe_proto = out.File
	file_games_tictactoe_proto_goTypes = nil
	file_games_tictactoe_proto_depIdxs = nil
}

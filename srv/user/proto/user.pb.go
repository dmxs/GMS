// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserItem struct {
	Img                  string   `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserItem) Reset()         { *m = UserItem{} }
func (m *UserItem) String() string { return proto.CompactTextString(m) }
func (*UserItem) ProtoMessage()    {}
func (*UserItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *UserItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserItem.Unmarshal(m, b)
}
func (m *UserItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserItem.Marshal(b, m, deterministic)
}
func (m *UserItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserItem.Merge(m, src)
}
func (m *UserItem) XXX_Size() int {
	return xxx_messageInfo_UserItem.Size(m)
}
func (m *UserItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UserItem.DiscardUnknown(m)
}

var xxx_messageInfo_UserItem proto.InternalMessageInfo

func (m *UserItem) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *UserItem) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Error struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PostReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReq) Reset()         { *m = PostReq{} }
func (m *PostReq) String() string { return proto.CompactTextString(m) }
func (*PostReq) ProtoMessage()    {}
func (*PostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *PostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReq.Unmarshal(m, b)
}
func (m *PostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReq.Marshal(b, m, deterministic)
}
func (m *PostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReq.Merge(m, src)
}
func (m *PostReq) XXX_Size() int {
	return xxx_messageInfo_PostReq.Size(m)
}
func (m *PostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReq.DiscardUnknown(m)
}

var xxx_messageInfo_PostReq proto.InternalMessageInfo

func (m *PostReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PostResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResp) Reset()         { *m = PostResp{} }
func (m *PostResp) String() string { return proto.CompactTextString(m) }
func (*PostResp) ProtoMessage()    {}
func (*PostResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *PostResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResp.Unmarshal(m, b)
}
func (m *PostResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResp.Marshal(b, m, deterministic)
}
func (m *PostResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResp.Merge(m, src)
}
func (m *PostResp) XXX_Size() int {
	return xxx_messageInfo_PostResp.Size(m)
}
func (m *PostResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResp.DiscardUnknown(m)
}

var xxx_messageInfo_PostResp proto.InternalMessageInfo

func (m *PostResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type CheckPasswordReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPasswordReq) Reset()         { *m = CheckPasswordReq{} }
func (m *CheckPasswordReq) String() string { return proto.CompactTextString(m) }
func (*CheckPasswordReq) ProtoMessage()    {}
func (*CheckPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *CheckPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPasswordReq.Unmarshal(m, b)
}
func (m *CheckPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPasswordReq.Marshal(b, m, deterministic)
}
func (m *CheckPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPasswordReq.Merge(m, src)
}
func (m *CheckPasswordReq) XXX_Size() int {
	return xxx_messageInfo_CheckPasswordReq.Size(m)
}
func (m *CheckPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPasswordReq proto.InternalMessageInfo

func (m *CheckPasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CheckPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CheckPasswordResp struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Result               bool      `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	User                 *UserItem `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CheckPasswordResp) Reset()         { *m = CheckPasswordResp{} }
func (m *CheckPasswordResp) String() string { return proto.CompactTextString(m) }
func (*CheckPasswordResp) ProtoMessage()    {}
func (*CheckPasswordResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *CheckPasswordResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPasswordResp.Unmarshal(m, b)
}
func (m *CheckPasswordResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPasswordResp.Marshal(b, m, deterministic)
}
func (m *CheckPasswordResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPasswordResp.Merge(m, src)
}
func (m *CheckPasswordResp) XXX_Size() int {
	return xxx_messageInfo_CheckPasswordResp.Size(m)
}
func (m *CheckPasswordResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPasswordResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPasswordResp proto.InternalMessageInfo

func (m *CheckPasswordResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CheckPasswordResp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CheckPasswordResp) GetUser() *UserItem {
	if m != nil {
		return m.User
	}
	return nil
}

type GetListReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListReq) Reset()         { *m = GetListReq{} }
func (m *GetListReq) String() string { return proto.CompactTextString(m) }
func (*GetListReq) ProtoMessage()    {}
func (*GetListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *GetListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListReq.Unmarshal(m, b)
}
func (m *GetListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListReq.Marshal(b, m, deterministic)
}
func (m *GetListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListReq.Merge(m, src)
}
func (m *GetListReq) XXX_Size() int {
	return xxx_messageInfo_GetListReq.Size(m)
}
func (m *GetListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetListReq proto.InternalMessageInfo

func (m *GetListReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetListReq) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetListResp struct {
	Error                *Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	List                 []*UserItem `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetListResp) Reset()         { *m = GetListResp{} }
func (m *GetListResp) String() string { return proto.CompactTextString(m) }
func (*GetListResp) ProtoMessage()    {}
func (*GetListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *GetListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResp.Unmarshal(m, b)
}
func (m *GetListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResp.Marshal(b, m, deterministic)
}
func (m *GetListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResp.Merge(m, src)
}
func (m *GetListResp) XXX_Size() int {
	return xxx_messageInfo_GetListResp.Size(m)
}
func (m *GetListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResp proto.InternalMessageInfo

func (m *GetListResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetListResp) GetList() []*UserItem {
	if m != nil {
		return m.List
	}
	return nil
}

type GetReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{8}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResp struct {
	Error                *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Result               bool      `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	User                 *UserItem `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{9}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetResp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *GetResp) GetUser() *UserItem {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserItem)(nil), "nezha.chat.user.srv.service.UserItem")
	proto.RegisterType((*Error)(nil), "nezha.chat.user.srv.service.Error")
	proto.RegisterType((*PostReq)(nil), "nezha.chat.user.srv.service.PostReq")
	proto.RegisterType((*PostResp)(nil), "nezha.chat.user.srv.service.PostResp")
	proto.RegisterType((*CheckPasswordReq)(nil), "nezha.chat.user.srv.service.CheckPasswordReq")
	proto.RegisterType((*CheckPasswordResp)(nil), "nezha.chat.user.srv.service.CheckPasswordResp")
	proto.RegisterType((*GetListReq)(nil), "nezha.chat.user.srv.service.GetListReq")
	proto.RegisterType((*GetListResp)(nil), "nezha.chat.user.srv.service.GetListResp")
	proto.RegisterType((*GetReq)(nil), "nezha.chat.user.srv.service.GetReq")
	proto.RegisterType((*GetResp)(nil), "nezha.chat.user.srv.service.GetResp")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x3d, 0x8f, 0xda, 0x40,
	0x10, 0x8d, 0x3f, 0x30, 0xce, 0xa0, 0x44, 0x64, 0xa4, 0x44, 0x96, 0xd3, 0xa0, 0x0d, 0x28, 0x34,
	0x38, 0x92, 0xd3, 0x84, 0x32, 0x5f, 0x22, 0x89, 0x52, 0x20, 0x4b, 0x34, 0x29, 0x22, 0x39, 0xf6,
	0x0a, 0xac, 0x60, 0xec, 0xec, 0x2e, 0x44, 0xba, 0xf2, 0xee, 0x3f, 0x5c, 0x7d, 0x3f, 0xf5, 0xb4,
	0x63, 0xc3, 0x1d, 0x14, 0x86, 0x13, 0xcd, 0x35, 0xd6, 0xcc, 0x78, 0xde, 0x9b, 0xb7, 0xb3, 0xcf,
	0x86, 0x97, 0xa5, 0x28, 0x54, 0xf1, 0x6e, 0x2d, 0xb9, 0xa0, 0x47, 0x40, 0x39, 0xbe, 0x5e, 0xf1,
	0x8b, 0x45, 0x1c, 0x24, 0x8b, 0x58, 0x05, 0x54, 0x96, 0x62, 0x13, 0x48, 0x2e, 0x36, 0x59, 0xc2,
	0xd9, 0x37, 0x70, 0x67, 0x92, 0x8b, 0xef, 0x8a, 0xe7, 0xd8, 0x05, 0x2b, 0xcb, 0xe7, 0x9e, 0xd1,
	0x33, 0x86, 0x4f, 0x23, 0x1d, 0xa2, 0x0f, 0xae, 0x46, 0xac, 0xe2, 0x9c, 0x7b, 0x26, 0x95, 0x77,
	0x39, 0x3e, 0x07, 0x33, 0x4b, 0x3d, 0x8b, 0xaa, 0x66, 0x96, 0xb2, 0x11, 0xb4, 0xbe, 0x0a, 0x51,
	0x08, 0x44, 0xb0, 0x93, 0x22, 0xe5, 0xc4, 0x63, 0x45, 0x14, 0x6b, 0xea, 0x5c, 0xce, 0x6b, 0x0e,
	0x1d, 0xb2, 0x8f, 0xd0, 0x9e, 0x16, 0x52, 0x45, 0xfc, 0xdf, 0xde, 0x14, 0xe3, 0x60, 0x8a, 0x0f,
	0x6e, 0x19, 0x4b, 0xf9, 0xbf, 0x10, 0xe9, 0x56, 0xc1, 0x36, 0x67, 0x5f, 0xc0, 0xad, 0x28, 0x64,
	0x89, 0x1f, 0xa0, 0xc5, 0xf5, 0x74, 0x22, 0xe8, 0x84, 0x2c, 0x68, 0x38, 0x74, 0x40, 0x3a, 0xa3,
	0x0a, 0xc0, 0x7e, 0x40, 0xf7, 0xf3, 0x82, 0x27, 0x7f, 0xa7, 0x35, 0xed, 0x39, 0x8a, 0x6e, 0x0c,
	0x78, 0x71, 0x40, 0x76, 0x8e, 0x36, 0x7c, 0x05, 0x8e, 0xe0, 0x72, 0xbd, 0x54, 0x34, 0xc9, 0x8d,
	0xea, 0x0c, 0xc7, 0x60, 0x6b, 0x20, 0x6d, 0xbf, 0x13, 0x0e, 0x1a, 0x09, 0xb7, 0xd7, 0x1b, 0x11,
	0x84, 0x85, 0x00, 0x13, 0xae, 0x7e, 0x66, 0xd5, 0xea, 0x11, 0xec, 0x7b, 0x87, 0xa4, 0x98, 0x6c,
	0x90, 0x4a, 0xcf, 0xec, 0x59, 0x64, 0x83, 0x54, 0xb2, 0x4b, 0x03, 0x3a, 0x3b, 0xd0, 0x59, 0x07,
	0x1a, 0x83, 0xbd, 0xcc, 0xa4, 0x22, 0xf2, 0xd3, 0x85, 0x6b, 0x08, 0xf3, 0xc0, 0x99, 0x70, 0x12,
	0x5d, 0x39, 0xcf, 0xd8, 0x39, 0xef, 0xda, 0x80, 0x36, 0xbd, 0x7a, 0x64, 0xbb, 0x0e, 0xaf, 0x2c,
	0xb0, 0x75, 0x09, 0x67, 0x60, 0x6b, 0xa7, 0x62, 0xbf, 0x11, 0x5d, 0x7f, 0x0f, 0xfe, 0xe0, 0x84,
	0x2e, 0x59, 0xb2, 0x27, 0x58, 0xc2, 0xb3, 0x3d, 0xb7, 0xe1, 0xa8, 0x11, 0x79, 0x68, 0x73, 0x3f,
	0x78, 0x48, 0x3b, 0x4d, 0xfc, 0x4d, 0x9b, 0xd6, 0x46, 0xc0, 0xb7, 0x8d, 0xe0, 0x3b, 0x8f, 0xf9,
	0xc3, 0xd3, 0x1a, 0x89, 0x3f, 0x02, 0x6b, 0xc2, 0x15, 0xbe, 0x39, 0x06, 0xd1, 0xbc, 0xfd, 0xe3,
	0x4d, 0x9a, 0xf3, 0x93, 0xf3, 0x8b, 0x6e, 0xe3, 0x8f, 0x43, 0xbf, 0xc3, 0xf7, 0xb7, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0x09, 0xa1, 0xf0, 0x27, 0x05, 0x00, 0x00,
}

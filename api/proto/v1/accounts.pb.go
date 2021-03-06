// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

//InsertAccount service
type InsertAccountRequest struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertAccountRequest) Reset()         { *m = InsertAccountRequest{} }
func (m *InsertAccountRequest) String() string { return proto.CompactTextString(m) }
func (*InsertAccountRequest) ProtoMessage()    {}
func (*InsertAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *InsertAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertAccountRequest.Unmarshal(m, b)
}
func (m *InsertAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertAccountRequest.Marshal(b, m, deterministic)
}
func (m *InsertAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertAccountRequest.Merge(m, src)
}
func (m *InsertAccountRequest) XXX_Size() int {
	return xxx_messageInfo_InsertAccountRequest.Size(m)
}
func (m *InsertAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertAccountRequest proto.InternalMessageInfo

func (m *InsertAccountRequest) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type InsertAccountResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertAccountResponse) Reset()         { *m = InsertAccountResponse{} }
func (m *InsertAccountResponse) String() string { return proto.CompactTextString(m) }
func (*InsertAccountResponse) ProtoMessage()    {}
func (*InsertAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *InsertAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertAccountResponse.Unmarshal(m, b)
}
func (m *InsertAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertAccountResponse.Marshal(b, m, deterministic)
}
func (m *InsertAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertAccountResponse.Merge(m, src)
}
func (m *InsertAccountResponse) XXX_Size() int {
	return xxx_messageInfo_InsertAccountResponse.Size(m)
}
func (m *InsertAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsertAccountResponse proto.InternalMessageInfo

func (m *InsertAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

//UpdateAccount Service
type UpdateAccountRequest struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccountRequest) Reset()         { *m = UpdateAccountRequest{} }
func (m *UpdateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountRequest) ProtoMessage()    {}
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *UpdateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountRequest.Unmarshal(m, b)
}
func (m *UpdateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountRequest.Merge(m, src)
}
func (m *UpdateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountRequest.Size(m)
}
func (m *UpdateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountRequest proto.InternalMessageInfo

func (m *UpdateAccountRequest) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateAccountResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccountResponse) Reset()         { *m = UpdateAccountResponse{} }
func (m *UpdateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountResponse) ProtoMessage()    {}
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{4}
}

func (m *UpdateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountResponse.Unmarshal(m, b)
}
func (m *UpdateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountResponse.Marshal(b, m, deterministic)
}
func (m *UpdateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountResponse.Merge(m, src)
}
func (m *UpdateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountResponse.Size(m)
}
func (m *UpdateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountResponse proto.InternalMessageInfo

func (m *UpdateAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

//GetAccount Service
type GetAccountRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{5}
}

func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(m, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAccountResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountResponse) Reset()         { *m = GetAccountResponse{} }
func (m *GetAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountResponse) ProtoMessage()    {}
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{6}
}

func (m *GetAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountResponse.Unmarshal(m, b)
}
func (m *GetAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountResponse.Merge(m, src)
}
func (m *GetAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountResponse.Size(m)
}
func (m *GetAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountResponse proto.InternalMessageInfo

func (m *GetAccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

//DeleteAccount Service
type DeleteAccountRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountRequest) Reset()         { *m = DeleteAccountRequest{} }
func (m *DeleteAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountRequest) ProtoMessage()    {}
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{7}
}

func (m *DeleteAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountRequest.Unmarshal(m, b)
}
func (m *DeleteAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountRequest.Merge(m, src)
}
func (m *DeleteAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountRequest.Size(m)
}
func (m *DeleteAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountRequest proto.InternalMessageInfo

func (m *DeleteAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteAccountResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountResponse) Reset()         { *m = DeleteAccountResponse{} }
func (m *DeleteAccountResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountResponse) ProtoMessage()    {}
func (*DeleteAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{8}
}

func (m *DeleteAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountResponse.Unmarshal(m, b)
}
func (m *DeleteAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountResponse.Merge(m, src)
}
func (m *DeleteAccountResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountResponse.Size(m)
}
func (m *DeleteAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountResponse proto.InternalMessageInfo

func (m *DeleteAccountResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//StreamAccounts Service
type StreamAccountsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccountsRequest) Reset()         { *m = StreamAccountsRequest{} }
func (m *StreamAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamAccountsRequest) ProtoMessage()    {}
func (*StreamAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{9}
}

func (m *StreamAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccountsRequest.Unmarshal(m, b)
}
func (m *StreamAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccountsRequest.Marshal(b, m, deterministic)
}
func (m *StreamAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccountsRequest.Merge(m, src)
}
func (m *StreamAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamAccountsRequest.Size(m)
}
func (m *StreamAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccountsRequest proto.InternalMessageInfo

type StreamAccountsResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccountsResponse) Reset()         { *m = StreamAccountsResponse{} }
func (m *StreamAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamAccountsResponse) ProtoMessage()    {}
func (*StreamAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{10}
}

func (m *StreamAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccountsResponse.Unmarshal(m, b)
}
func (m *StreamAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccountsResponse.Marshal(b, m, deterministic)
}
func (m *StreamAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccountsResponse.Merge(m, src)
}
func (m *StreamAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamAccountsResponse.Size(m)
}
func (m *StreamAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccountsResponse proto.InternalMessageInfo

func (m *StreamAccountsResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "account.Account")
	proto.RegisterType((*InsertAccountRequest)(nil), "account.InsertAccountRequest")
	proto.RegisterType((*InsertAccountResponse)(nil), "account.InsertAccountResponse")
	proto.RegisterType((*UpdateAccountRequest)(nil), "account.UpdateAccountRequest")
	proto.RegisterType((*UpdateAccountResponse)(nil), "account.UpdateAccountResponse")
	proto.RegisterType((*GetAccountRequest)(nil), "account.GetAccountRequest")
	proto.RegisterType((*GetAccountResponse)(nil), "account.GetAccountResponse")
	proto.RegisterType((*DeleteAccountRequest)(nil), "account.DeleteAccountRequest")
	proto.RegisterType((*DeleteAccountResponse)(nil), "account.DeleteAccountResponse")
	proto.RegisterType((*StreamAccountsRequest)(nil), "account.StreamAccountsRequest")
	proto.RegisterType((*StreamAccountsResponse)(nil), "account.StreamAccountsResponse")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xe5, 0x43, 0x04, 0xc7, 0x80, 0xba, 0x29, 0xda, 0xd4, 0xf8, 0x91, 0x35, 0x31, 0xc6, 0x03,
	0x51, 0xfc, 0x03, 0x82, 0x24, 0xc4, 0x1b, 0x81, 0x70, 0xf1, 0xb6, 0xb6, 0x13, 0xd3, 0x04, 0xda,
	0xda, 0x69, 0xf5, 0x27, 0xf9, 0x37, 0x8d, 0xed, 0xb6, 0xb4, 0xcb, 0x72, 0x00, 0x6f, 0x7d, 0x7d,
	0x33, 0x6f, 0x5e, 0x67, 0x5e, 0x0a, 0x1d, 0x61, 0xdb, 0x7e, 0xec, 0x45, 0xd4, 0x0b, 0x42, 0x3f,
	0xf2, 0x59, 0x53, 0x62, 0xfe, 0x01, 0xcd, 0x41, 0xfa, 0xc8, 0x3a, 0x50, 0x73, 0x1d, 0xb3, 0x7a,
	0x5d, 0xbd, 0x3b, 0x98, 0xd6, 0x5c, 0x87, 0x59, 0xd0, 0x8a, 0x09, 0x43, 0x4f, 0x2c, 0xd1, 0xac,
	0x25, 0x6f, 0x73, 0xfc, 0xc7, 0x05, 0x82, 0xe8, 0xdb, 0x0f, 0x1d, 0xb3, 0x9e, 0x72, 0x19, 0x66,
	0x06, 0x34, 0x70, 0x29, 0xdc, 0x85, 0xb9, 0x97, 0x10, 0x29, 0xe0, 0x43, 0x30, 0x5e, 0x3d, 0xc2,
	0x30, 0x92, 0xe3, 0xa6, 0xf8, 0x19, 0x23, 0x45, 0xec, 0x1e, 0x32, 0x2f, 0xc9, 0xe8, 0xc3, 0xfe,
	0x71, 0x4f, 0xe2, 0x5e, 0x56, 0x99, 0x9b, 0x7d, 0x81, 0xae, 0xa2, 0x41, 0x81, 0xef, 0x11, 0x6e,
	0x25, 0x32, 0x04, 0x63, 0x1e, 0x38, 0x22, 0xc2, 0xff, 0x19, 0x51, 0x34, 0x76, 0x30, 0x72, 0x03,
	0x27, 0x63, 0x54, 0xd7, 0xa1, 0x1c, 0x81, 0x3f, 0x03, 0x2b, 0x16, 0xed, 0x30, 0xe6, 0x16, 0x8c,
	0x11, 0x2e, 0x70, 0xed, 0x7b, 0xd5, 0x49, 0x8f, 0xd0, 0x55, 0xea, 0xe4, 0x30, 0x13, 0x9a, 0x14,
	0xdb, 0x36, 0x12, 0x25, 0xd5, 0xad, 0x69, 0x06, 0xf9, 0x19, 0x74, 0x67, 0x51, 0x88, 0x62, 0x29,
	0x5b, 0x48, 0x6a, 0xf3, 0x11, 0x9c, 0xaa, 0xc4, 0xf6, 0xce, 0xfb, 0x3f, 0x75, 0x38, 0x1a, 0xd8,
	0x29, 0x98, 0x61, 0xf8, 0xe5, 0xda, 0xc8, 0x26, 0xd0, 0x2e, 0x45, 0x80, 0x5d, 0xe4, 0xfd, 0xba,
	0x78, 0x59, 0x97, 0x9b, 0xe8, 0xd4, 0x0f, 0xaf, 0xb0, 0x31, 0xc0, 0x6a, 0xc3, 0xcc, 0xca, 0xeb,
	0xd7, 0x6e, 0x63, 0x9d, 0x6b, 0xb9, 0x5c, 0x68, 0x02, 0xed, 0xd2, 0x02, 0x0b, 0xd6, 0x74, 0x07,
	0x28, 0x58, 0xd3, 0xee, 0x3d, 0x55, 0x2c, 0xc5, 0xac, 0xa0, 0xa8, 0x8b, 0x70, 0x41, 0x51, 0x9b,
	0x4e, 0x5e, 0x61, 0x73, 0xe8, 0x94, 0x0f, 0xc3, 0x56, 0x3d, 0xda, 0x53, 0x5a, 0x57, 0x1b, 0xf9,
	0x4c, 0xf4, 0xa1, 0x3a, 0x6c, 0xbc, 0xd5, 0x45, 0xe0, 0xbe, 0xef, 0x27, 0x3f, 0x97, 0xa7, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x16, 0xd7, 0x0b, 0x6e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AcccountServiceClient is the client API for AcccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AcccountServiceClient interface {
	InsertAccount(ctx context.Context, in *InsertAccountRequest, opts ...grpc.CallOption) (*InsertAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	StreamAccounts(ctx context.Context, in *StreamAccountsRequest, opts ...grpc.CallOption) (AcccountService_StreamAccountsClient, error)
}

type acccountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAcccountServiceClient(cc *grpc.ClientConn) AcccountServiceClient {
	return &acccountServiceClient{cc}
}

func (c *acccountServiceClient) InsertAccount(ctx context.Context, in *InsertAccountRequest, opts ...grpc.CallOption) (*InsertAccountResponse, error) {
	out := new(InsertAccountResponse)
	err := c.cc.Invoke(ctx, "/account.AcccountService/InsertAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acccountServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/account.AcccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acccountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/account.AcccountService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acccountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.AcccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acccountServiceClient) StreamAccounts(ctx context.Context, in *StreamAccountsRequest, opts ...grpc.CallOption) (AcccountService_StreamAccountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AcccountService_serviceDesc.Streams[0], "/account.AcccountService/StreamAccounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &acccountServiceStreamAccountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AcccountService_StreamAccountsClient interface {
	Recv() (*StreamAccountsResponse, error)
	grpc.ClientStream
}

type acccountServiceStreamAccountsClient struct {
	grpc.ClientStream
}

func (x *acccountServiceStreamAccountsClient) Recv() (*StreamAccountsResponse, error) {
	m := new(StreamAccountsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AcccountServiceServer is the server API for AcccountService service.
type AcccountServiceServer interface {
	InsertAccount(context.Context, *InsertAccountRequest) (*InsertAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	StreamAccounts(*StreamAccountsRequest, AcccountService_StreamAccountsServer) error
}

// UnimplementedAcccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAcccountServiceServer struct {
}

func (*UnimplementedAcccountServiceServer) InsertAccount(ctx context.Context, req *InsertAccountRequest) (*InsertAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAccount not implemented")
}
func (*UnimplementedAcccountServiceServer) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedAcccountServiceServer) DeleteAccount(ctx context.Context, req *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (*UnimplementedAcccountServiceServer) UpdateAccount(ctx context.Context, req *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (*UnimplementedAcccountServiceServer) StreamAccounts(req *StreamAccountsRequest, srv AcccountService_StreamAccountsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAccounts not implemented")
}

func RegisterAcccountServiceServer(s *grpc.Server, srv AcccountServiceServer) {
	s.RegisterService(&_AcccountService_serviceDesc, srv)
}

func _AcccountService_InsertAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcccountServiceServer).InsertAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AcccountService/InsertAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcccountServiceServer).InsertAccount(ctx, req.(*InsertAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AcccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcccountServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AcccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AcccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcccountService_StreamAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamAccountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AcccountServiceServer).StreamAccounts(m, &acccountServiceStreamAccountsServer{stream})
}

type AcccountService_StreamAccountsServer interface {
	Send(*StreamAccountsResponse) error
	grpc.ServerStream
}

type acccountServiceStreamAccountsServer struct {
	grpc.ServerStream
}

func (x *acccountServiceStreamAccountsServer) Send(m *StreamAccountsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AcccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AcccountService",
	HandlerType: (*AcccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertAccount",
			Handler:    _AcccountService_InsertAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AcccountService_GetAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AcccountService_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AcccountService_UpdateAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccounts",
			Handler:       _AcccountService_StreamAccounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "accounts.proto",
}

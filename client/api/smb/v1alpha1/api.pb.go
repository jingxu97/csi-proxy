// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/api/smb/v1alpha1/api.proto

package v1alpha1

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

type MountSmbShareRequest struct {
	// A remote SMB share to mount
	// All unicode characters allowed in SMB server name specifications are
	// permitted except for restrictions below
	//
	// Restrictions:
	// SMB remote path specified in the format: \\server-name\sharename, \\server.fqdn\sharename or \\a.b.c.d\sharename
	// If not an IP address, share name has to be a valid DNS name.
	// UNC specifications to local paths or prefix: \\?\ is not allowed.
	// Characters: + [ ] " / : ; | < > , ? * = $ are not allowed.
	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	// Local path in the host to stage the SMB share.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// Needs to be an absolute path under kubelet-csi-plugins-path.
	// Needs to exist already in host
	// Path needs to be specified with drive letter prefix: "X:\".
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// Maximum path length will be capped to 260 characters (MAX_PATH).
	LocalPath string `protobuf:"bytes,2,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	// Mount the share read-only
	Readonly bool `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	// Username credential associated with the share
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	// Password credential associated with the share
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountSmbShareRequest) Reset()         { *m = MountSmbShareRequest{} }
func (m *MountSmbShareRequest) String() string { return proto.CompactTextString(m) }
func (*MountSmbShareRequest) ProtoMessage()    {}
func (*MountSmbShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33320d52df402875, []int{0}
}

func (m *MountSmbShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountSmbShareRequest.Unmarshal(m, b)
}
func (m *MountSmbShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountSmbShareRequest.Marshal(b, m, deterministic)
}
func (m *MountSmbShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountSmbShareRequest.Merge(m, src)
}
func (m *MountSmbShareRequest) XXX_Size() int {
	return xxx_messageInfo_MountSmbShareRequest.Size(m)
}
func (m *MountSmbShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MountSmbShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MountSmbShareRequest proto.InternalMessageInfo

func (m *MountSmbShareRequest) GetRemotePath() string {
	if m != nil {
		return m.RemotePath
	}
	return ""
}

func (m *MountSmbShareRequest) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

func (m *MountSmbShareRequest) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *MountSmbShareRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MountSmbShareRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MountSmbShareResponse struct {
	// Windows error code
	// Success is represented as 0
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountSmbShareResponse) Reset()         { *m = MountSmbShareResponse{} }
func (m *MountSmbShareResponse) String() string { return proto.CompactTextString(m) }
func (*MountSmbShareResponse) ProtoMessage()    {}
func (*MountSmbShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33320d52df402875, []int{1}
}

func (m *MountSmbShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountSmbShareResponse.Unmarshal(m, b)
}
func (m *MountSmbShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountSmbShareResponse.Marshal(b, m, deterministic)
}
func (m *MountSmbShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountSmbShareResponse.Merge(m, src)
}
func (m *MountSmbShareResponse) XXX_Size() int {
	return xxx_messageInfo_MountSmbShareResponse.Size(m)
}
func (m *MountSmbShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MountSmbShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MountSmbShareResponse proto.InternalMessageInfo

func (m *MountSmbShareResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UnmountSmbShareRequest struct {
	// A remote SMB share mapping to remove
	// All unicode characters allowed in SMB server name specifications are
	// permitted except for restrictions below
	//
	// Restrictions:
	// SMB share specified in the format: \\server-name\sharename, \\server.fqdn\sharename or \\a.b.c.d\sharename
	// If not an IP address, share name has to be a valid DNS name.
	// UNC specifications to local paths or prefix: \\?\ is not allowed.
	// Characters: + [ ] " / : ; | < > , ? * = $ are not allowed.
	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	// Local path in the host to stage the SMB share.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// Needs to be an absolute path under kubelet-csi-plugins-path.
	// Needs to exist already in host
	// Path needs to be specified with drive letter prefix: "X:\".
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// Maximum path length will be capped to 260 characters (MAX_PATH).
	LocalPath            string   `protobuf:"bytes,2,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnmountSmbShareRequest) Reset()         { *m = UnmountSmbShareRequest{} }
func (m *UnmountSmbShareRequest) String() string { return proto.CompactTextString(m) }
func (*UnmountSmbShareRequest) ProtoMessage()    {}
func (*UnmountSmbShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33320d52df402875, []int{2}
}

func (m *UnmountSmbShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnmountSmbShareRequest.Unmarshal(m, b)
}
func (m *UnmountSmbShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnmountSmbShareRequest.Marshal(b, m, deterministic)
}
func (m *UnmountSmbShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnmountSmbShareRequest.Merge(m, src)
}
func (m *UnmountSmbShareRequest) XXX_Size() int {
	return xxx_messageInfo_UnmountSmbShareRequest.Size(m)
}
func (m *UnmountSmbShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnmountSmbShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnmountSmbShareRequest proto.InternalMessageInfo

func (m *UnmountSmbShareRequest) GetRemotePath() string {
	if m != nil {
		return m.RemotePath
	}
	return ""
}

func (m *UnmountSmbShareRequest) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

type UnmountSmbShareResponse struct {
	// Windows error code
	// Success is represented as 0
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnmountSmbShareResponse) Reset()         { *m = UnmountSmbShareResponse{} }
func (m *UnmountSmbShareResponse) String() string { return proto.CompactTextString(m) }
func (*UnmountSmbShareResponse) ProtoMessage()    {}
func (*UnmountSmbShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33320d52df402875, []int{3}
}

func (m *UnmountSmbShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnmountSmbShareResponse.Unmarshal(m, b)
}
func (m *UnmountSmbShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnmountSmbShareResponse.Marshal(b, m, deterministic)
}
func (m *UnmountSmbShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnmountSmbShareResponse.Merge(m, src)
}
func (m *UnmountSmbShareResponse) XXX_Size() int {
	return xxx_messageInfo_UnmountSmbShareResponse.Size(m)
}
func (m *UnmountSmbShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnmountSmbShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnmountSmbShareResponse proto.InternalMessageInfo

func (m *UnmountSmbShareResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*MountSmbShareRequest)(nil), "v1alpha1.MountSmbShareRequest")
	proto.RegisterType((*MountSmbShareResponse)(nil), "v1alpha1.MountSmbShareResponse")
	proto.RegisterType((*UnmountSmbShareRequest)(nil), "v1alpha1.UnmountSmbShareRequest")
	proto.RegisterType((*UnmountSmbShareResponse)(nil), "v1alpha1.UnmountSmbShareResponse")
}

func init() { proto.RegisterFile("client/api/smb/v1alpha1/api.proto", fileDescriptor_33320d52df402875) }

var fileDescriptor_33320d52df402875 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x8d, 0xb5, 0x92, 0x8e, 0x88, 0xb0, 0x54, 0x0d, 0x01, 0x6d, 0x9a, 0x53, 0x2f, 0x26,
	0x54, 0x9f, 0x43, 0x90, 0x04, 0xa1, 0x37, 0x99, 0xb4, 0x03, 0x29, 0x64, 0xff, 0xb8, 0xbb, 0x51,
	0x7c, 0x25, 0xaf, 0xbe, 0xa0, 0x64, 0xd7, 0x54, 0xac, 0xb5, 0x37, 0x8f, 0xdf, 0xf7, 0xdb, 0x99,
	0xfd, 0x66, 0x67, 0x61, 0xba, 0x6c, 0xd6, 0x24, 0x6c, 0x8e, 0x6a, 0x9d, 0x1b, 0x5e, 0xe5, 0x2f,
	0x73, 0x6c, 0x54, 0x8d, 0xf3, 0xce, 0xc8, 0x94, 0x96, 0x56, 0xb2, 0xb0, 0xf7, 0xd2, 0xf7, 0x00,
	0xc6, 0xf7, 0xb2, 0x15, 0xb6, 0xe4, 0x55, 0x59, 0xa3, 0xa6, 0x82, 0x9e, 0x5b, 0x32, 0x96, 0x4d,
	0xe0, 0x44, 0x13, 0x97, 0x96, 0x9e, 0x14, 0xda, 0x3a, 0x0a, 0x92, 0x60, 0x36, 0x2a, 0xc0, 0x5b,
	0x0f, 0x68, 0x6b, 0x76, 0x05, 0xd0, 0xc8, 0x25, 0x36, 0x9e, 0x1f, 0x3a, 0x3e, 0x72, 0x8e, 0xc3,
	0x31, 0x84, 0x9a, 0x70, 0x25, 0x45, 0xf3, 0x16, 0x0d, 0x92, 0x60, 0x16, 0x16, 0x1b, 0xdd, 0xb1,
	0xd6, 0x90, 0x16, 0xc8, 0x29, 0x3a, 0x72, 0x85, 0x1b, 0xdd, 0x31, 0x85, 0xc6, 0xbc, 0x4a, 0xbd,
	0x8a, 0x86, 0x9e, 0xf5, 0x3a, 0xbd, 0x81, 0xf3, 0xad, 0xac, 0x46, 0x49, 0x61, 0x88, 0x8d, 0x61,
	0x48, 0x5a, 0x4b, 0xfd, 0x15, 0xd3, 0x8b, 0x74, 0x01, 0x17, 0x8f, 0x82, 0xff, 0xc3, 0x70, 0x69,
	0x0e, 0x97, 0xbf, 0x3a, 0xef, 0x8b, 0x72, 0xfb, 0x11, 0xc0, 0xa0, 0xe4, 0x15, 0x2b, 0xe0, 0xf4,
	0xc7, 0x04, 0xec, 0x3a, 0xeb, 0x57, 0x91, 0xed, 0x5a, 0x43, 0x3c, 0xf9, 0x93, 0xfb, 0xfb, 0xd2,
	0x03, 0xb6, 0x80, 0xb3, 0xad, 0x30, 0x2c, 0xf9, 0xae, 0xda, 0xfd, 0x02, 0xf1, 0x74, 0xcf, 0x89,
	0xbe, 0x73, 0x75, 0xec, 0x7e, 0xcb, 0xdd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x1d, 0xdc,
	0xc0, 0x52, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmbClient is the client API for Smb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmbClient interface {
	// MountSmbShare mounts a remote share over SMB on the host at the requested global staging path.
	MountSmbShare(ctx context.Context, in *MountSmbShareRequest, opts ...grpc.CallOption) (*MountSmbShareResponse, error)
	// UnmountSmbShare unmounts a remote share over SMB on the host at the requested global staging path.
	UnmountSmbShare(ctx context.Context, in *UnmountSmbShareRequest, opts ...grpc.CallOption) (*UnmountSmbShareResponse, error)
}

type smbClient struct {
	cc *grpc.ClientConn
}

func NewSmbClient(cc *grpc.ClientConn) SmbClient {
	return &smbClient{cc}
}

func (c *smbClient) MountSmbShare(ctx context.Context, in *MountSmbShareRequest, opts ...grpc.CallOption) (*MountSmbShareResponse, error) {
	out := new(MountSmbShareResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Smb/MountSmbShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smbClient) UnmountSmbShare(ctx context.Context, in *UnmountSmbShareRequest, opts ...grpc.CallOption) (*UnmountSmbShareResponse, error) {
	out := new(UnmountSmbShareResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Smb/UnmountSmbShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmbServer is the server API for Smb service.
type SmbServer interface {
	// MountSmbShare mounts a remote share over SMB on the host at the requested global staging path.
	MountSmbShare(context.Context, *MountSmbShareRequest) (*MountSmbShareResponse, error)
	// UnmountSmbShare unmounts a remote share over SMB on the host at the requested global staging path.
	UnmountSmbShare(context.Context, *UnmountSmbShareRequest) (*UnmountSmbShareResponse, error)
}

// UnimplementedSmbServer can be embedded to have forward compatible implementations.
type UnimplementedSmbServer struct {
}

func (*UnimplementedSmbServer) MountSmbShare(ctx context.Context, req *MountSmbShareRequest) (*MountSmbShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountSmbShare not implemented")
}
func (*UnimplementedSmbServer) UnmountSmbShare(ctx context.Context, req *UnmountSmbShareRequest) (*UnmountSmbShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnmountSmbShare not implemented")
}

func RegisterSmbServer(s *grpc.Server, srv SmbServer) {
	s.RegisterService(&_Smb_serviceDesc, srv)
}

func _Smb_MountSmbShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountSmbShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmbServer).MountSmbShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Smb/MountSmbShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmbServer).MountSmbShare(ctx, req.(*MountSmbShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Smb_UnmountSmbShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmountSmbShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmbServer).UnmountSmbShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Smb/UnmountSmbShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmbServer).UnmountSmbShare(ctx, req.(*UnmountSmbShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Smb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Smb",
	HandlerType: (*SmbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MountSmbShare",
			Handler:    _Smb_MountSmbShare_Handler,
		},
		{
			MethodName: "UnmountSmbShare",
			Handler:    _Smb_UnmountSmbShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/api/smb/v1alpha1/api.proto",
}

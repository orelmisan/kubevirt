// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/handler-launcher-com/cmd/v1/cmd.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	pkg/handler-launcher-com/cmd/v1/cmd.proto

It has these top-level messages:
	VMI
	SMBios
	VirtualMachineOptions
	VMIRequest
	MigrationRequest
	EmptyRequest
	Response
	DomainResponse
	DomainStatsResponse
	GuestInfoResponse
	GuestUserListResponse
	GuestFilesystemsResponse
*/
package v1

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VMI struct {
	VmiJson []byte `protobuf:"bytes,1,opt,name=vmiJson,proto3" json:"vmiJson,omitempty"`
}

func (m *VMI) Reset()                    { *m = VMI{} }
func (m *VMI) String() string            { return proto.CompactTextString(m) }
func (*VMI) ProtoMessage()               {}
func (*VMI) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VMI) GetVmiJson() []byte {
	if m != nil {
		return m.VmiJson
	}
	return nil
}

type SMBios struct {
	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer" json:"manufacturer,omitempty"`
	Product      string `protobuf:"bytes,2,opt,name=product" json:"product,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Sku          string `protobuf:"bytes,4,opt,name=sku" json:"sku,omitempty"`
	Family       string `protobuf:"bytes,5,opt,name=family" json:"family,omitempty"`
}

func (m *SMBios) Reset()                    { *m = SMBios{} }
func (m *SMBios) String() string            { return proto.CompactTextString(m) }
func (*SMBios) ProtoMessage()               {}
func (*SMBios) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SMBios) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *SMBios) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *SMBios) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SMBios) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *SMBios) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

type VirtualMachineOptions struct {
	VirtualMachineSMBios  *SMBios `protobuf:"bytes,1,opt,name=VirtualMachineSMBios" json:"VirtualMachineSMBios,omitempty"`
	MemBalloonStatsPeriod uint32  `protobuf:"varint,2,opt,name=MemBalloonStatsPeriod" json:"MemBalloonStatsPeriod,omitempty"`
}

func (m *VirtualMachineOptions) Reset()                    { *m = VirtualMachineOptions{} }
func (m *VirtualMachineOptions) String() string            { return proto.CompactTextString(m) }
func (*VirtualMachineOptions) ProtoMessage()               {}
func (*VirtualMachineOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VirtualMachineOptions) GetVirtualMachineSMBios() *SMBios {
	if m != nil {
		return m.VirtualMachineSMBios
	}
	return nil
}

func (m *VirtualMachineOptions) GetMemBalloonStatsPeriod() uint32 {
	if m != nil {
		return m.MemBalloonStatsPeriod
	}
	return 0
}

type VMIRequest struct {
	Vmi     *VMI                   `protobuf:"bytes,1,opt,name=vmi" json:"vmi,omitempty"`
	Options *VirtualMachineOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *VMIRequest) Reset()                    { *m = VMIRequest{} }
func (m *VMIRequest) String() string            { return proto.CompactTextString(m) }
func (*VMIRequest) ProtoMessage()               {}
func (*VMIRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VMIRequest) GetVmi() *VMI {
	if m != nil {
		return m.Vmi
	}
	return nil
}

func (m *VMIRequest) GetOptions() *VirtualMachineOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type MigrationRequest struct {
	Vmi     *VMI   `protobuf:"bytes,1,opt,name=vmi" json:"vmi,omitempty"`
	Options []byte `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (m *MigrationRequest) Reset()                    { *m = MigrationRequest{} }
func (m *MigrationRequest) String() string            { return proto.CompactTextString(m) }
func (*MigrationRequest) ProtoMessage()               {}
func (*MigrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MigrationRequest) GetVmi() *VMI {
	if m != nil {
		return m.Vmi
	}
	return nil
}

func (m *MigrationRequest) GetOptions() []byte {
	if m != nil {
		return m.Options
	}
	return nil
}

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Response struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DomainResponse struct {
	Response *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Domain   string    `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
}

func (m *DomainResponse) Reset()                    { *m = DomainResponse{} }
func (m *DomainResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainResponse) ProtoMessage()               {}
func (*DomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DomainResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DomainResponse) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type DomainStatsResponse struct {
	Response    *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	DomainStats string    `protobuf:"bytes,2,opt,name=domainStats" json:"domainStats,omitempty"`
}

func (m *DomainStatsResponse) Reset()                    { *m = DomainStatsResponse{} }
func (m *DomainStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainStatsResponse) ProtoMessage()               {}
func (*DomainStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DomainStatsResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DomainStatsResponse) GetDomainStats() string {
	if m != nil {
		return m.DomainStats
	}
	return ""
}

type GuestInfoResponse struct {
	Response          *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	GuestInfoResponse string    `protobuf:"bytes,2,opt,name=guestInfoResponse" json:"guestInfoResponse,omitempty"`
}

func (m *GuestInfoResponse) Reset()                    { *m = GuestInfoResponse{} }
func (m *GuestInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GuestInfoResponse) ProtoMessage()               {}
func (*GuestInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GuestInfoResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GuestInfoResponse) GetGuestInfoResponse() string {
	if m != nil {
		return m.GuestInfoResponse
	}
	return ""
}

type GuestUserListResponse struct {
	Response              *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	GuestUserListResponse string    `protobuf:"bytes,2,opt,name=guestUserListResponse" json:"guestUserListResponse,omitempty"`
}

func (m *GuestUserListResponse) Reset()                    { *m = GuestUserListResponse{} }
func (m *GuestUserListResponse) String() string            { return proto.CompactTextString(m) }
func (*GuestUserListResponse) ProtoMessage()               {}
func (*GuestUserListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GuestUserListResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GuestUserListResponse) GetGuestUserListResponse() string {
	if m != nil {
		return m.GuestUserListResponse
	}
	return ""
}

type GuestFilesystemsResponse struct {
	Response                 *Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	GuestFilesystemsResponse string    `protobuf:"bytes,2,opt,name=guestFilesystemsResponse" json:"guestFilesystemsResponse,omitempty"`
}

func (m *GuestFilesystemsResponse) Reset()                    { *m = GuestFilesystemsResponse{} }
func (m *GuestFilesystemsResponse) String() string            { return proto.CompactTextString(m) }
func (*GuestFilesystemsResponse) ProtoMessage()               {}
func (*GuestFilesystemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GuestFilesystemsResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GuestFilesystemsResponse) GetGuestFilesystemsResponse() string {
	if m != nil {
		return m.GuestFilesystemsResponse
	}
	return ""
}

func init() {
	proto.RegisterType((*VMI)(nil), "kubevirt.cmd.v1.VMI")
	proto.RegisterType((*SMBios)(nil), "kubevirt.cmd.v1.SMBios")
	proto.RegisterType((*VirtualMachineOptions)(nil), "kubevirt.cmd.v1.VirtualMachineOptions")
	proto.RegisterType((*VMIRequest)(nil), "kubevirt.cmd.v1.VMIRequest")
	proto.RegisterType((*MigrationRequest)(nil), "kubevirt.cmd.v1.MigrationRequest")
	proto.RegisterType((*EmptyRequest)(nil), "kubevirt.cmd.v1.EmptyRequest")
	proto.RegisterType((*Response)(nil), "kubevirt.cmd.v1.Response")
	proto.RegisterType((*DomainResponse)(nil), "kubevirt.cmd.v1.DomainResponse")
	proto.RegisterType((*DomainStatsResponse)(nil), "kubevirt.cmd.v1.DomainStatsResponse")
	proto.RegisterType((*GuestInfoResponse)(nil), "kubevirt.cmd.v1.GuestInfoResponse")
	proto.RegisterType((*GuestUserListResponse)(nil), "kubevirt.cmd.v1.GuestUserListResponse")
	proto.RegisterType((*GuestFilesystemsResponse)(nil), "kubevirt.cmd.v1.GuestFilesystemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cmd service

type CmdClient interface {
	SyncVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	PauseVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	UnpauseVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	ShutdownVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	KillVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	MigrateVirtualMachine(ctx context.Context, in *MigrationRequest, opts ...grpc.CallOption) (*Response, error)
	SyncMigrationTarget(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	CancelVirtualMachineMigration(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	SetVirtualMachineGuestTime(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	SignalTargetPodCleanup(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error)
	GetDomain(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DomainResponse, error)
	GetDomainStats(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DomainStatsResponse, error)
	GetGuestInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestInfoResponse, error)
	GetUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestUserListResponse, error)
	GetFilesystems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestFilesystemsResponse, error)
	Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error)
}

type cmdClient struct {
	cc *grpc.ClientConn
}

func NewCmdClient(cc *grpc.ClientConn) CmdClient {
	return &cmdClient{cc}
}

func (c *cmdClient) SyncVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/SyncVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) PauseVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/PauseVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) UnpauseVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/UnpauseVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) ShutdownVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/ShutdownVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) KillVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/KillVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) DeleteVirtualMachine(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/DeleteVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) MigrateVirtualMachine(ctx context.Context, in *MigrationRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/MigrateVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) SyncMigrationTarget(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/SyncMigrationTarget", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) CancelVirtualMachineMigration(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/CancelVirtualMachineMigration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) SetVirtualMachineGuestTime(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/SetVirtualMachineGuestTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) SignalTargetPodCleanup(ctx context.Context, in *VMIRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/SignalTargetPodCleanup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) GetDomain(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DomainResponse, error) {
	out := new(DomainResponse)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/GetDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) GetDomainStats(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DomainStatsResponse, error) {
	out := new(DomainStatsResponse)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/GetDomainStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) GetGuestInfo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestInfoResponse, error) {
	out := new(GuestInfoResponse)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/GetGuestInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) GetUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestUserListResponse, error) {
	out := new(GuestUserListResponse)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) GetFilesystems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GuestFilesystemsResponse, error) {
	out := new(GuestFilesystemsResponse)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/GetFilesystems", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdClient) Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/kubevirt.cmd.v1.Cmd/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cmd service

type CmdServer interface {
	SyncVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	PauseVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	UnpauseVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	ShutdownVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	KillVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	DeleteVirtualMachine(context.Context, *VMIRequest) (*Response, error)
	MigrateVirtualMachine(context.Context, *MigrationRequest) (*Response, error)
	SyncMigrationTarget(context.Context, *VMIRequest) (*Response, error)
	CancelVirtualMachineMigration(context.Context, *VMIRequest) (*Response, error)
	SetVirtualMachineGuestTime(context.Context, *VMIRequest) (*Response, error)
	SignalTargetPodCleanup(context.Context, *VMIRequest) (*Response, error)
	GetDomain(context.Context, *EmptyRequest) (*DomainResponse, error)
	GetDomainStats(context.Context, *EmptyRequest) (*DomainStatsResponse, error)
	GetGuestInfo(context.Context, *EmptyRequest) (*GuestInfoResponse, error)
	GetUsers(context.Context, *EmptyRequest) (*GuestUserListResponse, error)
	GetFilesystems(context.Context, *EmptyRequest) (*GuestFilesystemsResponse, error)
	Ping(context.Context, *EmptyRequest) (*Response, error)
}

func RegisterCmdServer(s *grpc.Server, srv CmdServer) {
	s.RegisterService(&_Cmd_serviceDesc, srv)
}

func _Cmd_SyncVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).SyncVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/SyncVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).SyncVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_PauseVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).PauseVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/PauseVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).PauseVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_UnpauseVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).UnpauseVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/UnpauseVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).UnpauseVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_ShutdownVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).ShutdownVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/ShutdownVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).ShutdownVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_KillVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).KillVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/KillVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).KillVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_DeleteVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).DeleteVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/DeleteVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).DeleteVirtualMachine(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_MigrateVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).MigrateVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/MigrateVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).MigrateVirtualMachine(ctx, req.(*MigrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_SyncMigrationTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).SyncMigrationTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/SyncMigrationTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).SyncMigrationTarget(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_CancelVirtualMachineMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).CancelVirtualMachineMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/CancelVirtualMachineMigration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).CancelVirtualMachineMigration(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_SetVirtualMachineGuestTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).SetVirtualMachineGuestTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/SetVirtualMachineGuestTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).SetVirtualMachineGuestTime(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_SignalTargetPodCleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).SignalTargetPodCleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/SignalTargetPodCleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).SignalTargetPodCleanup(ctx, req.(*VMIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/GetDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).GetDomain(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_GetDomainStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).GetDomainStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/GetDomainStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).GetDomainStats(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_GetGuestInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).GetGuestInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/GetGuestInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).GetGuestInfo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).GetUsers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_GetFilesystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).GetFilesystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/GetFilesystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).GetFilesystems(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmd_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubevirt.cmd.v1.Cmd/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServer).Ping(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cmd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubevirt.cmd.v1.Cmd",
	HandlerType: (*CmdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncVirtualMachine",
			Handler:    _Cmd_SyncVirtualMachine_Handler,
		},
		{
			MethodName: "PauseVirtualMachine",
			Handler:    _Cmd_PauseVirtualMachine_Handler,
		},
		{
			MethodName: "UnpauseVirtualMachine",
			Handler:    _Cmd_UnpauseVirtualMachine_Handler,
		},
		{
			MethodName: "ShutdownVirtualMachine",
			Handler:    _Cmd_ShutdownVirtualMachine_Handler,
		},
		{
			MethodName: "KillVirtualMachine",
			Handler:    _Cmd_KillVirtualMachine_Handler,
		},
		{
			MethodName: "DeleteVirtualMachine",
			Handler:    _Cmd_DeleteVirtualMachine_Handler,
		},
		{
			MethodName: "MigrateVirtualMachine",
			Handler:    _Cmd_MigrateVirtualMachine_Handler,
		},
		{
			MethodName: "SyncMigrationTarget",
			Handler:    _Cmd_SyncMigrationTarget_Handler,
		},
		{
			MethodName: "CancelVirtualMachineMigration",
			Handler:    _Cmd_CancelVirtualMachineMigration_Handler,
		},
		{
			MethodName: "SetVirtualMachineGuestTime",
			Handler:    _Cmd_SetVirtualMachineGuestTime_Handler,
		},
		{
			MethodName: "SignalTargetPodCleanup",
			Handler:    _Cmd_SignalTargetPodCleanup_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _Cmd_GetDomain_Handler,
		},
		{
			MethodName: "GetDomainStats",
			Handler:    _Cmd_GetDomainStats_Handler,
		},
		{
			MethodName: "GetGuestInfo",
			Handler:    _Cmd_GetGuestInfo_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Cmd_GetUsers_Handler,
		},
		{
			MethodName: "GetFilesystems",
			Handler:    _Cmd_GetFilesystems_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Cmd_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/handler-launcher-com/cmd/v1/cmd.proto",
}

func init() { proto.RegisterFile("pkg/handler-launcher-com/cmd/v1/cmd.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0xc7, 0x13, 0x42, 0x21, 0xbc, 0xa4, 0x14, 0x06, 0x42, 0x5d, 0x2a, 0x04, 0x1d, 0x55, 0xa8,
	0x48, 0x25, 0x08, 0x4a, 0x2f, 0x3d, 0x54, 0x55, 0xa0, 0x8d, 0x28, 0x0d, 0xa4, 0x0e, 0x64, 0xb5,
	0xab, 0x95, 0x56, 0x83, 0x3d, 0x38, 0x23, 0xec, 0x19, 0xaf, 0x67, 0x9c, 0xdd, 0xdc, 0xf7, 0xb4,
	0xd2, 0xde, 0xf6, 0xb4, 0x7f, 0xed, 0xca, 0x63, 0x27, 0x90, 0xd8, 0xd9, 0x68, 0xe5, 0x9c, 0xe2,
	0x37, 0xef, 0xcd, 0xe7, 0xfb, 0xde, 0xfc, 0x78, 0x13, 0x38, 0xf0, 0x1f, 0x9c, 0xa3, 0x1e, 0xe1,
	0xb6, 0x4b, 0x83, 0x43, 0x97, 0x84, 0xdc, 0xea, 0xd1, 0xe0, 0xd0, 0x12, 0xde, 0x91, 0xe5, 0xd9,
	0x47, 0xfd, 0xe3, 0xe8, 0xa7, 0xee, 0x07, 0x42, 0x09, 0xf4, 0xdd, 0x43, 0x78, 0x47, 0xfb, 0x2c,
	0x50, 0xf5, 0x68, 0xac, 0x7f, 0x8c, 0x77, 0xa1, 0xd4, 0x6d, 0x5d, 0x20, 0x03, 0x96, 0xfb, 0x1e,
	0xfb, 0x57, 0x0a, 0x6e, 0x14, 0xf7, 0x8a, 0xbf, 0x54, 0xcd, 0xa1, 0x89, 0xdf, 0x17, 0x61, 0xa9,
	0xd3, 0x6a, 0x30, 0x21, 0x11, 0x86, 0xaa, 0x47, 0x78, 0x78, 0x4f, 0x2c, 0x15, 0x06, 0x34, 0xd0,
	0x91, 0x2b, 0xe6, 0xd8, 0x58, 0x04, 0xf2, 0x03, 0x61, 0x87, 0x96, 0x32, 0x16, 0xb4, 0x7b, 0x68,
	0x6a, 0x09, 0x1a, 0x48, 0x26, 0xb8, 0x51, 0x8a, 0x3d, 0x89, 0x89, 0xd6, 0xa0, 0x24, 0x1f, 0x42,
	0x63, 0x51, 0x8f, 0x46, 0x9f, 0x68, 0x0b, 0x96, 0xee, 0x89, 0xc7, 0xdc, 0x81, 0xf1, 0x8d, 0x1e,
	0x4c, 0x2c, 0xfc, 0xa9, 0x08, 0xb5, 0x2e, 0x0b, 0x54, 0x48, 0xdc, 0x16, 0xb1, 0x7a, 0x8c, 0xd3,
	0x6b, 0x5f, 0x31, 0xc1, 0x25, 0xba, 0x84, 0xcd, 0x71, 0x47, 0x9c, 0xb3, 0xce, 0xb1, 0x72, 0xf2,
	0x7d, 0x7d, 0xa2, 0xee, 0x7a, 0xec, 0x36, 0x33, 0x27, 0xa1, 0x53, 0xa8, 0xb5, 0xa8, 0xd7, 0x20,
	0xae, 0x2b, 0x04, 0xef, 0x28, 0xa2, 0x64, 0x9b, 0x06, 0x4c, 0xd8, 0xba, 0xa4, 0x6f, 0xcd, 0x6c,
	0x27, 0xee, 0x03, 0x74, 0x5b, 0x17, 0x26, 0x7d, 0x1d, 0x52, 0xa9, 0xd0, 0x3e, 0x94, 0xfa, 0x1e,
	0x4b, 0xf4, 0x37, 0x53, 0xfa, 0x51, 0x64, 0x14, 0x80, 0xfe, 0x82, 0x65, 0x11, 0xd7, 0xa0, 0xe9,
	0x95, 0x93, 0xfd, 0x74, 0x6c, 0x56, 0xc5, 0xe6, 0x70, 0x1a, 0xbe, 0x81, 0xb5, 0x16, 0x73, 0x02,
	0x12, 0x59, 0x5f, 0xab, 0x6e, 0x8c, 0xab, 0x57, 0x1f, 0xa9, 0xab, 0x50, 0xfd, 0xdb, 0xf3, 0xd5,
	0x20, 0x21, 0xe2, 0x3f, 0xa1, 0x6c, 0x52, 0xe9, 0x0b, 0x2e, 0x69, 0x34, 0x4b, 0x86, 0x96, 0x45,
	0x65, 0xbc, 0xbe, 0x65, 0x73, 0x68, 0x46, 0x1e, 0x8f, 0x4a, 0x49, 0x1c, 0x3a, 0xdc, 0xfe, 0xc4,
	0xc4, 0xaf, 0x60, 0xf5, 0x5c, 0x78, 0x84, 0xf1, 0x11, 0xe5, 0x77, 0x28, 0x07, 0xc9, 0x77, 0x92,
	0xe8, 0x0f, 0xa9, 0x44, 0x87, 0xc1, 0xe6, 0x28, 0x34, 0x3a, 0x1b, 0xb6, 0x06, 0x25, 0x0a, 0x89,
	0x85, 0x39, 0x6c, 0xc4, 0x02, 0x7a, 0x4f, 0xf2, 0xaa, 0xec, 0x41, 0xc5, 0x7e, 0xa4, 0x25, 0x52,
	0x4f, 0x87, 0xf0, 0x5b, 0x58, 0x6f, 0x46, 0x2b, 0x73, 0xc1, 0xef, 0x45, 0x5e, 0xb5, 0x5f, 0x61,
	0xdd, 0x99, 0x64, 0x25, 0x9a, 0x69, 0x07, 0x7e, 0x57, 0x84, 0x9a, 0x96, 0xbe, 0x95, 0x34, 0xf8,
	0x8f, 0x49, 0x95, 0x57, 0xfe, 0x14, 0x6a, 0x4e, 0x16, 0x2f, 0x49, 0x21, 0xdb, 0x89, 0x3f, 0x14,
	0xc1, 0xd0, 0x69, 0xfc, 0xc3, 0x5c, 0x2a, 0x07, 0x52, 0x51, 0x2f, 0xf7, 0xb2, 0xff, 0x01, 0x86,
	0x33, 0x05, 0x99, 0x24, 0x33, 0xd5, 0x7f, 0xf2, 0xb1, 0x02, 0xa5, 0x33, 0xcf, 0x46, 0x57, 0x80,
	0x3a, 0x03, 0x6e, 0x8d, 0xdf, 0x1a, 0xf4, 0x63, 0xe6, 0x25, 0x88, 0x0f, 0xf7, 0xf6, 0xf4, 0xdc,
	0x70, 0x01, 0x5d, 0xc3, 0x46, 0x9b, 0x84, 0x92, 0xce, 0x0d, 0xf8, 0x3f, 0xd4, 0x6e, 0xb9, 0x3f,
	0x57, 0xa4, 0x09, 0x5b, 0x9d, 0x5e, 0xa8, 0x6c, 0xf1, 0x86, 0xcf, 0x8d, 0x79, 0x05, 0xe8, 0x92,
	0xb9, 0xee, 0xdc, 0x78, 0x6d, 0xd8, 0x3c, 0xa7, 0x2e, 0x55, 0xf3, 0xab, 0xfa, 0x19, 0xd4, 0xe2,
	0xce, 0x37, 0x89, 0xfc, 0x29, 0x35, 0x6b, 0xb2, 0x43, 0xce, 0xdc, 0xf2, 0xe8, 0x08, 0x8d, 0x26,
	0xdd, 0x90, 0xc0, 0xa1, 0x2a, 0x47, 0xa6, 0xcf, 0x61, 0xe7, 0x8c, 0x70, 0x8b, 0x4e, 0xac, 0xe6,
	0x48, 0x20, 0x07, 0xba, 0x0b, 0xdb, 0x1d, 0xaa, 0xc6, 0xb9, 0xfa, 0x5a, 0xde, 0x30, 0x2f, 0xef,
	0x91, 0x62, 0x0e, 0x27, 0x6e, 0x5c, 0x7c, 0x5b, 0xd8, 0x67, 0x2e, 0x25, 0x3c, 0xf4, 0x73, 0x30,
	0x5b, 0xb0, 0xd2, 0xa4, 0x2a, 0x6e, 0xd3, 0x68, 0x27, 0x15, 0xf9, 0xf4, 0xc1, 0xd9, 0xde, 0x4d,
	0xb9, 0xc7, 0xdf, 0x0f, 0xbd, 0xff, 0xab, 0x23, 0x9c, 0x6e, 0xca, 0xb3, 0x98, 0x3f, 0x4f, 0x61,
	0x8e, 0x3d, 0x19, 0xb8, 0x80, 0x3a, 0x50, 0x6d, 0x52, 0x35, 0x6a, 0xef, 0xb3, 0xb0, 0x38, 0xe5,
	0x4e, 0xbd, 0x0c, 0x1a, 0x5a, 0x6e, 0x52, 0xdd, 0x46, 0x67, 0xe6, 0xb9, 0x9f, 0x0d, 0x4c, 0xb5,
	0xe0, 0x02, 0x7a, 0xa9, 0x97, 0xe0, 0x49, 0x3b, 0x9c, 0x85, 0x3e, 0xc8, 0x46, 0x67, 0x34, 0x54,
	0x5c, 0x40, 0x0d, 0x58, 0x6c, 0x33, 0xee, 0xcc, 0x62, 0x7e, 0x69, 0xcf, 0x1b, 0x8b, 0x2f, 0x16,
	0xfa, 0xc7, 0x77, 0x4b, 0xfa, 0xff, 0xe7, 0x6f, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x82,
	0xd4, 0x10, 0xac, 0x0a, 0x00, 0x00,
}

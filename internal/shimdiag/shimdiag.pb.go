// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Microsoft/hcsshim/internal/shimdiag/shimdiag.proto

package shimdiag

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExecProcessRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Workdir              string   `protobuf:"bytes,2,opt,name=workdir,proto3" json:"workdir,omitempty"`
	Terminal             bool     `protobuf:"varint,3,opt,name=terminal,proto3" json:"terminal,omitempty"`
	Stdin                string   `protobuf:"bytes,4,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout               string   `protobuf:"bytes,5,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               string   `protobuf:"bytes,6,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecProcessRequest) Reset()      { *m = ExecProcessRequest{} }
func (*ExecProcessRequest) ProtoMessage() {}
func (*ExecProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7933dc6ffbb8784, []int{0}
}
func (m *ExecProcessRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecProcessRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecProcessRequest.Merge(m, src)
}
func (m *ExecProcessRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExecProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecProcessRequest proto.InternalMessageInfo

type ExecProcessResponse struct {
	ExitCode             int32    `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecProcessResponse) Reset()      { *m = ExecProcessResponse{} }
func (*ExecProcessResponse) ProtoMessage() {}
func (*ExecProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7933dc6ffbb8784, []int{1}
}
func (m *ExecProcessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecProcessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecProcessResponse.Merge(m, src)
}
func (m *ExecProcessResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExecProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecProcessResponse proto.InternalMessageInfo

type StacksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StacksRequest) Reset()      { *m = StacksRequest{} }
func (*StacksRequest) ProtoMessage() {}
func (*StacksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7933dc6ffbb8784, []int{2}
}
func (m *StacksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StacksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StacksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StacksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StacksRequest.Merge(m, src)
}
func (m *StacksRequest) XXX_Size() int {
	return m.Size()
}
func (m *StacksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StacksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StacksRequest proto.InternalMessageInfo

type StacksResponse struct {
	Stacks               string   `protobuf:"bytes,1,opt,name=stacks,proto3" json:"stacks,omitempty"`
	GuestStacks          string   `protobuf:"bytes,2,opt,name=guest_stacks,json=guestStacks,proto3" json:"guest_stacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StacksResponse) Reset()      { *m = StacksResponse{} }
func (*StacksResponse) ProtoMessage() {}
func (*StacksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7933dc6ffbb8784, []int{3}
}
func (m *StacksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StacksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StacksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StacksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StacksResponse.Merge(m, src)
}
func (m *StacksResponse) XXX_Size() int {
	return m.Size()
}
func (m *StacksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StacksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StacksResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExecProcessRequest)(nil), "containerd.runhcs.v1.diag.ExecProcessRequest")
	proto.RegisterType((*ExecProcessResponse)(nil), "containerd.runhcs.v1.diag.ExecProcessResponse")
	proto.RegisterType((*StacksRequest)(nil), "containerd.runhcs.v1.diag.StacksRequest")
	proto.RegisterType((*StacksResponse)(nil), "containerd.runhcs.v1.diag.StacksResponse")
}

func init() {
	proto.RegisterFile("github.com/Microsoft/hcsshim/internal/shimdiag/shimdiag.proto", fileDescriptor_c7933dc6ffbb8784)
}

var fileDescriptor_c7933dc6ffbb8784 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x14, 0x5c, 0xd3, 0xee, 0xb2, 0xfb, 0x80, 0x22, 0x99, 0x0a, 0x99, 0x45, 0x8a, 0x96, 0x9c, 0xc2,
	0x01, 0x47, 0x94, 0x03, 0x07, 0xc4, 0x85, 0x0f, 0x09, 0x84, 0x90, 0x20, 0xbd, 0x20, 0x2e, 0x95,
	0xeb, 0x98, 0xc4, 0x6a, 0xe3, 0x57, 0x6c, 0x07, 0x7a, 0xe4, 0xc7, 0xf0, 0x63, 0x7a, 0xe4, 0xc8,
	0x09, 0xd1, 0xfc, 0x12, 0x64, 0x27, 0x59, 0xa9, 0x42, 0xa0, 0xed, 0x29, 0x6f, 0x26, 0x33, 0x99,
	0xb1, 0xf3, 0xe0, 0x69, 0xa5, 0x7d, 0xdd, 0x1e, 0x72, 0x89, 0x4d, 0xfe, 0x56, 0x4b, 0x8b, 0x0e,
	0x3f, 0xf9, 0xbc, 0x96, 0xce, 0xd5, 0xba, 0xc9, 0xb5, 0xf1, 0xca, 0x1a, 0x71, 0x9c, 0x07, 0x54,
	0x6a, 0x51, 0xad, 0x07, 0x7e, 0x62, 0xd1, 0x23, 0xbd, 0x23, 0xd1, 0x78, 0xa1, 0x8d, 0xb2, 0x25,
	0xb7, 0xad, 0xa9, 0xa5, 0xe3, 0x5f, 0x1e, 0xf2, 0x20, 0x58, 0xee, 0x56, 0x58, 0x61, 0x54, 0xe5,
	0x61, 0xea, 0x0d, 0xe9, 0x77, 0x02, 0xf4, 0xe5, 0xa9, 0x92, 0xef, 0x2c, 0x4a, 0xe5, 0x5c, 0xa1,
	0x3e, 0xb7, 0xca, 0x79, 0x4a, 0x61, 0x5b, 0xd8, 0xca, 0x31, 0xb2, 0xda, 0xca, 0x16, 0x45, 0x9c,
	0x29, 0x83, 0xab, 0x5f, 0xd1, 0x1e, 0x95, 0xda, 0xb2, 0x2b, 0x2b, 0x92, 0x2d, 0x8a, 0x11, 0xd2,
	0x25, 0xcc, 0xbd, 0xb2, 0x8d, 0x36, 0xe2, 0x98, 0x6d, 0xad, 0x48, 0x36, 0x2f, 0xd6, 0x98, 0xee,
	0xc2, 0xd4, 0xf9, 0x52, 0x1b, 0xb6, 0x1d, 0x3d, 0x3d, 0xa0, 0xb7, 0x61, 0xe6, 0x7c, 0x89, 0xad,
	0x67, 0xd3, 0x48, 0x0f, 0x68, 0xe0, 0x95, 0xb5, 0x6c, 0xb6, 0xe6, 0x95, 0xb5, 0xe9, 0x1e, 0xdc,
	0xba, 0xd0, 0xd2, 0x9d, 0xa0, 0x71, 0x8a, 0xde, 0x85, 0x85, 0x3a, 0xd5, 0xfe, 0x40, 0x62, 0xa9,
	0x18, 0x59, 0x91, 0x6c, 0x5a, 0xcc, 0x03, 0xf1, 0x1c, 0x4b, 0x95, 0xde, 0x84, 0x1b, 0xfb, 0x5e,
	0xc8, 0xa3, 0xf1, 0x50, 0xe9, 0x1b, 0xd8, 0x19, 0x89, 0xc1, 0x1f, 0xe3, 0x02, 0x13, 0xcd, 0x31,
	0x2e, 0x20, 0x7a, 0x0f, 0xae, 0x57, 0xc1, 0x72, 0x30, 0xbc, 0xed, 0xcf, 0x7b, 0x2d, 0x72, 0xfd,
	0x27, 0xf6, 0x7e, 0x11, 0x98, 0xef, 0xd7, 0xba, 0x79, 0xa1, 0x45, 0x45, 0x11, 0x76, 0xc2, 0x33,
	0x54, 0x7c, 0x6d, 0x5e, 0xa1, 0xf3, 0xf4, 0x01, 0xff, 0xe7, 0x9f, 0xe0, 0x7f, 0xdf, 0xf7, 0x92,
	0x6f, 0x2a, 0x1f, 0x8a, 0x0b, 0x80, 0x10, 0xd8, 0x77, 0xa1, 0xd9, 0x7f, 0xdc, 0x17, 0xae, 0x60,
	0x79, 0x7f, 0x03, 0x65, 0x1f, 0xf1, 0xec, 0xfd, 0xd9, 0x79, 0x32, 0xf9, 0x79, 0x9e, 0x4c, 0xbe,
	0x75, 0x09, 0x39, 0xeb, 0x12, 0xf2, 0xa3, 0x4b, 0xc8, 0xef, 0x2e, 0x21, 0x1f, 0x1f, 0x5f, 0x6e,
	0x47, 0x9f, 0x8c, 0xc3, 0x87, 0xc9, 0xe1, 0x2c, 0x6e, 0xdd, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x49, 0xfc, 0x83, 0xd9, 0xe7, 0x02, 0x00, 0x00,
}

func (m *ExecProcessRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecProcessRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Workdir) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.Workdir)))
		i += copy(dAtA[i:], m.Workdir)
	}
	if m.Terminal {
		dAtA[i] = 0x18
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Stdin) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.Stdin)))
		i += copy(dAtA[i:], m.Stdin)
	}
	if len(m.Stdout) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.Stdout)))
		i += copy(dAtA[i:], m.Stdout)
	}
	if len(m.Stderr) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.Stderr)))
		i += copy(dAtA[i:], m.Stderr)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ExecProcessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecProcessResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ExitCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(m.ExitCode))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StacksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StacksRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StacksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StacksResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Stacks) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.Stacks)))
		i += copy(dAtA[i:], m.Stacks)
	}
	if len(m.GuestStacks) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintShimdiag(dAtA, i, uint64(len(m.GuestStacks)))
		i += copy(dAtA[i:], m.GuestStacks)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintShimdiag(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExecProcessRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovShimdiag(uint64(l))
		}
	}
	l = len(m.Workdir)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExecProcessResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExitCode != 0 {
		n += 1 + sovShimdiag(uint64(m.ExitCode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StacksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StacksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Stacks)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	l = len(m.GuestStacks)
	if l > 0 {
		n += 1 + l + sovShimdiag(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShimdiag(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShimdiag(x uint64) (n int) {
	return sovShimdiag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ExecProcessRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExecProcessRequest{`,
		`Args:` + fmt.Sprintf("%v", this.Args) + `,`,
		`Workdir:` + fmt.Sprintf("%v", this.Workdir) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExecProcessResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExecProcessResponse{`,
		`ExitCode:` + fmt.Sprintf("%v", this.ExitCode) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StacksRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StacksRequest{`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StacksResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StacksResponse{`,
		`Stacks:` + fmt.Sprintf("%v", this.Stacks) + `,`,
		`GuestStacks:` + fmt.Sprintf("%v", this.GuestStacks) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringShimdiag(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type ShimDiagService interface {
	DiagExecInHost(ctx context.Context, req *ExecProcessRequest) (*ExecProcessResponse, error)
	DiagStacks(ctx context.Context, req *StacksRequest) (*StacksResponse, error)
}

func RegisterShimDiagService(srv *github_com_containerd_ttrpc.Server, svc ShimDiagService) {
	srv.Register("containerd.runhcs.v1.diag.ShimDiag", map[string]github_com_containerd_ttrpc.Method{
		"DiagExecInHost": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req ExecProcessRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.DiagExecInHost(ctx, &req)
		},
		"DiagStacks": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req StacksRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.DiagStacks(ctx, &req)
		},
	})
}

type shimDiagClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewShimDiagClient(client *github_com_containerd_ttrpc.Client) ShimDiagService {
	return &shimDiagClient{
		client: client,
	}
}

func (c *shimDiagClient) DiagExecInHost(ctx context.Context, req *ExecProcessRequest) (*ExecProcessResponse, error) {
	var resp ExecProcessResponse
	if err := c.client.Call(ctx, "containerd.runhcs.v1.diag.ShimDiag", "DiagExecInHost", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimDiagClient) DiagStacks(ctx context.Context, req *StacksRequest) (*StacksResponse, error) {
	var resp StacksResponse
	if err := c.client.Call(ctx, "containerd.runhcs.v1.diag.ShimDiag", "DiagStacks", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *ExecProcessRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShimdiag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecProcessRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecProcessRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Workdir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Workdir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShimdiag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecProcessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShimdiag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecProcessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecProcessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitCode", wireType)
			}
			m.ExitCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShimdiag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StacksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShimdiag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StacksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StacksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShimdiag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StacksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShimdiag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StacksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StacksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stacks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stacks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuestStacks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShimdiag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShimdiag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuestStacks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShimdiag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShimdiag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShimdiag(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShimdiag
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShimdiag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthShimdiag
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthShimdiag
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShimdiag
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipShimdiag(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthShimdiag
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthShimdiag = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShimdiag   = fmt.Errorf("proto: integer overflow")
)

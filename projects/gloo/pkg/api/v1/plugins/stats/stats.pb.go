// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/stats/stats.proto

package stats

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// This plugin provides additional configuration options to expose statistics.
type Stats struct {
	// Virtual clusters allow exposing additional statistics for traffic served by a Virtual Host.
	VirtualClusters      []*VirtualCluster `protobuf:"bytes,10,rep,name=virtual_clusters,json=virtualClusters,proto3" json:"virtual_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_027fb1039a85630a, []int{0}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetVirtualClusters() []*VirtualCluster {
	if m != nil {
		return m.VirtualClusters
	}
	return nil
}

// Virtual clusters allow you to expose statistics for virtual host traffic that matches certain criteria.
// This is useful because what the application considers to be an endpoint does often not map directly to
// the routing configuration, so Envoy does not emit per endpoint statistics. Using virtual clusters you can define
// logical endpoints and have Envoy emit dedicated statistics for any matching request. Virtual cluster statistics
// are emitted on the downstream side and thus include network level failures.
//
// Please note that virtual clusters add overhead to the processing of each requests and should not be overused.
type VirtualCluster struct {
	// The name of the virtual cluster. This value will be used together with the virtual host name to
	// compute the name of the statistics emitted by this virtual cluster. Statistics names will be in the form:
	// vhost.<virtual host name>.vcluster.<virtual cluster name>.<stat name>.
	// See [the official Envoy documentation](https://www.envoyproxy.io/docs/envoy/v1.5.0/configuration/http_filters/router_filter#config-http-filters-router-stats)
	// for more information about the statistics emitted when virtual cluster configurations are specified.
	//
	// Note: This string should not contain any dots ("."), as this is a reserved character for Envoy statistics names.
	// Any dot present in the virtual cluster name will be replaced with an underscore ("_") character by Gloo.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The regex pattern used by Envoy to decide whether to expose statistics for a particular request.
	// Please note that **the entire path** of the request must match the regex (e.g. the regex `/rides/d+` matches
	// the path `/rides/0`, but not `/rides/123/456`).
	// The regex grammar used is defined [here](https://en.cppreference.com/w/cpp/regex/ecmascript).
	Pattern string `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// If specified, statistics will be exposed only for requests matching the given HTTP method.
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualCluster) Reset()         { *m = VirtualCluster{} }
func (m *VirtualCluster) String() string { return proto.CompactTextString(m) }
func (*VirtualCluster) ProtoMessage()    {}
func (*VirtualCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_027fb1039a85630a, []int{1}
}
func (m *VirtualCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualCluster.Unmarshal(m, b)
}
func (m *VirtualCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualCluster.Marshal(b, m, deterministic)
}
func (m *VirtualCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualCluster.Merge(m, src)
}
func (m *VirtualCluster) XXX_Size() int {
	return xxx_messageInfo_VirtualCluster.Size(m)
}
func (m *VirtualCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualCluster.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualCluster proto.InternalMessageInfo

func (m *VirtualCluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualCluster) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *VirtualCluster) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func init() {
	proto.RegisterType((*Stats)(nil), "stats.plugins.gloo.solo.io.Stats")
	proto.RegisterType((*VirtualCluster)(nil), "stats.plugins.gloo.solo.io.VirtualCluster")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/stats/stats.proto", fileDescriptor_027fb1039a85630a)
}

var fileDescriptor_027fb1039a85630a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xab, 0x15, 0x47, 0x50, 0x09, 0x22, 0xa1, 0x07, 0x29, 0x3d, 0x15, 0xc1, 0x04,
	0xf5, 0xee, 0x41, 0xc5, 0x07, 0xa8, 0xd8, 0x83, 0x07, 0x25, 0x5d, 0x43, 0x1a, 0xcd, 0xee, 0x84,
	0x64, 0x76, 0x9f, 0xc9, 0xe7, 0xf2, 0x49, 0x24, 0xc9, 0x7a, 0x28, 0x28, 0xf4, 0x12, 0xe6, 0xcb,
	0xfc, 0xdf, 0x1c, 0x7e, 0x78, 0x34, 0x96, 0xd6, 0xdd, 0x4a, 0xd4, 0xd8, 0xc8, 0x88, 0x0e, 0x2f,
	0x2d, 0x4a, 0xe3, 0x10, 0xa5, 0x0f, 0xf8, 0xa1, 0x6b, 0x8a, 0x85, 0x94, 0xb7, 0xb2, 0xbf, 0x92,
	0xde, 0x75, 0xc6, 0xb6, 0x51, 0x46, 0x52, 0x34, 0xbc, 0xc2, 0x07, 0x24, 0x64, 0x93, 0x01, 0x4a,
	0x40, 0x24, 0x49, 0xa4, 0x7b, 0xc2, 0xe2, 0xe4, 0xd4, 0xa0, 0xc1, 0x1c, 0x93, 0x69, 0x2a, 0xc6,
	0xec, 0x15, 0xf6, 0x9e, 0x92, 0xc3, 0x9e, 0xe1, 0xa4, 0xb7, 0x81, 0x3a, 0xe5, 0xde, 0x6a, 0xd7,
	0x45, 0xd2, 0x21, 0x72, 0x98, 0x8e, 0xe6, 0x87, 0xd7, 0x17, 0xe2, 0xff, 0xab, 0x62, 0x59, 0x9c,
	0xfb, 0xa2, 0x2c, 0x8e, 0xfb, 0x0d, 0x8e, 0xb3, 0x25, 0x1c, 0x6d, 0x46, 0x18, 0x83, 0xdd, 0x56,
	0x35, 0x9a, 0x57, 0xd3, 0x6a, 0x7e, 0xb0, 0xc8, 0x33, 0xe3, 0xb0, 0xef, 0x15, 0x91, 0x0e, 0x2d,
	0xdf, 0xc9, 0xdf, 0xbf, 0xc8, 0xce, 0x60, 0xdc, 0x68, 0x5a, 0xe3, 0x3b, 0x1f, 0xe5, 0xc5, 0x40,
	0x77, 0x0f, 0x5f, 0xdf, 0xe7, 0xd5, 0xcb, 0xed, 0x76, 0xbd, 0xf9, 0x4f, 0xf3, 0x67, 0x77, 0xab,
	0x71, 0x2e, 0xe1, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x65, 0xe9, 0xc1, 0x80, 0x01, 0x00,
	0x00,
}

func (this *Stats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Stats)
	if !ok {
		that2, ok := that.(Stats)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.VirtualClusters) != len(that1.VirtualClusters) {
		return false
	}
	for i := range this.VirtualClusters {
		if !this.VirtualClusters[i].Equal(that1.VirtualClusters[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualCluster) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualCluster)
	if !ok {
		that2, ok := that.(VirtualCluster)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Pattern != that1.Pattern {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

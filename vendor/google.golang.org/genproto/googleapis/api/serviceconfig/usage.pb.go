// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration controlling usage of a service.
type Usage struct {
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements" json:"requirements,omitempty"`
	// A list of usage rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules" json:"rules,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// True, if the method allows unregistered calls; false otherwise.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls" json:"allow_unregistered_calls,omitempty"`
}

func (m *UsageRule) Reset()                    { *m = UsageRule{} }
func (m *UsageRule) String() string            { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()               {}
func (*UsageRule) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto", fileDescriptor15)
}

var fileDescriptor15 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x8f, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6b, 0x4b, 0x3b, 0x8a, 0x87, 0x80, 0x12, 0x7a, 0x0a, 0x01, 0x21, 0x20, 0x24,
	0xa0, 0x17, 0xaf, 0xb6, 0x07, 0xe9, 0x2d, 0x2c, 0x14, 0xbc, 0x95, 0x71, 0x1d, 0x97, 0x85, 0xed,
	0x4e, 0xdc, 0x0f, 0xfd, 0x3f, 0xfe, 0x52, 0xc9, 0x46, 0x6a, 0xbd, 0x7a, 0x9c, 0x79, 0x86, 0x67,
	0xde, 0x17, 0x36, 0x8a, 0x59, 0x19, 0x6a, 0x14, 0x1b, 0xb4, 0xaa, 0x61, 0xa7, 0x5a, 0x45, 0xb6,
	0x77, 0x1c, 0xb8, 0x1d, 0x11, 0xf6, 0xda, 0xb7, 0xd8, 0xeb, 0xd6, 0x93, 0xfb, 0xd0, 0x92, 0x24,
	0xdb, 0x37, 0xad, 0xda, 0xe8, 0x51, 0x51, 0x93, 0x0e, 0x73, 0xf8, 0x91, 0x60, 0xaf, 0x57, 0xdb,
	0xff, 0x0a, 0xd1, 0x5a, 0x0e, 0x18, 0x34, 0x5b, 0x3f, 0x6a, 0xab, 0x67, 0x98, 0xed, 0x86, 0x2f,
	0x79, 0x05, 0x17, 0x8e, 0xde, 0xa3, 0x76, 0x74, 0x20, 0x1b, 0x7c, 0x91, 0x95, 0xd3, 0x7a, 0x29,
	0xfe, 0xec, 0xf2, 0x5b, 0x98, 0xb9, 0x68, 0xc8, 0x17, 0xf3, 0x72, 0x5a, 0x9f, 0xdf, 0x5d, 0x35,
	0xbf, 0x99, 0x9a, 0x64, 0x11, 0xd1, 0x90, 0x18, 0x6f, 0x2a, 0x84, 0xe5, 0x71, 0x97, 0xaf, 0x60,
	0xe1, 0xc9, 0x90, 0x0c, 0xec, 0x8a, 0xac, 0xcc, 0xea, 0xa5, 0x38, 0xce, 0xf9, 0x03, 0x14, 0x68,
	0x0c, 0x7f, 0xee, 0xa3, 0x75, 0xa4, 0xb4, 0x0f, 0xe4, 0xe8, 0x75, 0x2f, 0xd1, 0x18, 0x5f, 0x4c,
	0xca, 0xac, 0x5e, 0x88, 0xeb, 0xc4, 0x77, 0x27, 0x78, 0x33, 0xd0, 0xf5, 0x0d, 0x5c, 0x4a, 0x3e,
	0x9c, 0xa4, 0x58, 0x43, 0x7a, 0xd9, 0x0d, 0xd5, 0xba, 0xec, 0x6b, 0x72, 0xf6, 0xf4, 0xd8, 0x6d,
	0x5f, 0xe6, 0xa9, 0xea, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x2d, 0x47, 0x30, 0x88,
	0x01, 0x00, 0x00,
}

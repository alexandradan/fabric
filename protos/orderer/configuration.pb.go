// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absoluteMaxBytes" json:"absoluteMaxBytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferredMaxBytes" json:"preferredMaxBytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// IngressPolicy is the name of the policy which incoming Broadcast messages are filtered against
type IngressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IngressPolicy) Reset()                    { *m = IngressPolicy{} }
func (m *IngressPolicy) String() string            { return proto.CompactTextString(m) }
func (*IngressPolicy) ProtoMessage()               {}
func (*IngressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// EgressPolicy is the name of the policy which incoming Deliver messages are filtered against
type EgressPolicy struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *EgressPolicy) Reset()                    { *m = EgressPolicy{} }
func (m *EgressPolicy) String() string            { return proto.CompactTextString(m) }
func (*EgressPolicy) ProtoMessage()               {}
func (*EgressPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation
	// policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*IngressPolicy)(nil), "orderer.IngressPolicy")
	proto.RegisterType((*EgressPolicy)(nil), "orderer.EgressPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0x71, 0x13, 0x92, 0x46, 0x24, 0xfd, 0xa3, 0x42, 0x31, 0xe9, 0x25, 0xb8, 0x17, 0xd3,
	0x86, 0xf8, 0xd0, 0x17, 0x28, 0x36, 0x7b, 0x58, 0x96, 0xc0, 0xe2, 0xcd, 0x69, 0x6f, 0xb2, 0x3d,
	0xb6, 0x45, 0x62, 0x8d, 0x19, 0xc9, 0x10, 0xef, 0x4b, 0xec, 0x2b, 0x2f, 0x96, 0x95, 0x1c, 0x36,
	0x87, 0x3d, 0xf9, 0xfb, 0x66, 0x7e, 0x58, 0xdf, 0x8c, 0xc4, 0x7e, 0x21, 0x15, 0x40, 0x40, 0x51,
	0x8e, 0xaa, 0x94, 0x55, 0x47, 0xc2, 0x48, 0x54, 0xbb, 0x96, 0xd0, 0x20, 0x9f, 0xbb, 0xe6, 0xfa,
	0x47, 0x8e, 0x4d, 0x83, 0x2a, 0x1a, 0x3f, 0x63, 0x37, 0xf8, 0xcd, 0x56, 0x09, 0x2a, 0x0d, 0x4a,
	0x77, 0xfa, 0xd0, 0xb7, 0xc0, 0x39, 0x9b, 0x9a, 0xbe, 0x05, 0xdf, 0xdb, 0x78, 0xe1, 0x22, 0xb5,
	0x3a, 0x78, 0xf5, 0xd8, 0x22, 0x16, 0x26, 0xaf, 0x9f, 0xe4, 0x0b, 0xf0, 0x90, 0x7d, 0x6d, 0xc4,
	0x79, 0x0f, 0x5a, 0x8b, 0x0a, 0x12, 0xec, 0x94, 0xb1, 0xf0, 0x2a, 0x7d, 0x5f, 0xe6, 0x7f, 0xd8,
	0x37, 0x91, 0x69, 0x3c, 0x75, 0x06, 0xf6, 0xe2, 0x1c, 0xf7, 0x06, 0xb4, 0xff, 0xc9, 0xa2, 0x37,
	0x75, 0xbe, 0x65, 0xdf, 0x5b, 0x82, 0x12, 0x88, 0xa0, 0xb8, 0xc2, 0x13, 0x0b, 0xdf, 0x36, 0x82,
	0x90, 0x2d, 0x6d, 0xa0, 0x83, 0x6c, 0x00, 0x3b, 0xc3, 0x7d, 0x36, 0x37, 0xa3, 0x74, 0xc1, 0x2f,
	0x36, 0xf8, 0xcf, 0xbe, 0x24, 0x04, 0x76, 0x21, 0x8f, 0x78, 0x92, 0x79, 0xcf, 0x7f, 0xb2, 0x59,
	0x6b, 0x95, 0x43, 0x9d, 0x1b, 0xea, 0x85, 0xac, 0x40, 0x1b, 0x9b, 0x71, 0x99, 0x3a, 0x37, 0xac,
	0xe8, 0x5e, 0x55, 0x04, 0x5a, 0xbb, 0x1f, 0x70, 0x36, 0x55, 0xa2, 0xb9, 0xae, 0x68, 0xd0, 0x41,
	0xc0, 0x96, 0x77, 0x1f, 0x31, 0x7f, 0xd9, 0x2a, 0xa9, 0x85, 0x54, 0x36, 0x0f, 0x92, 0xe6, 0x6b,
	0xf6, 0xd9, 0x9e, 0x2d, 0x41, 0xfb, 0xde, 0x66, 0x12, 0x2e, 0xd2, 0xab, 0x1f, 0x26, 0x7c, 0x10,
	0xe5, 0x51, 0xc4, 0x84, 0x47, 0x20, 0x3d, 0x4c, 0x98, 0x8d, 0xd2, 0xa1, 0x17, 0x1b, 0xef, 0x9e,
	0xb7, 0x95, 0x34, 0x75, 0x97, 0xed, 0x72, 0x6c, 0xa2, 0xba, 0x6f, 0x81, 0x4e, 0x50, 0x54, 0x40,
	0x51, 0x29, 0x32, 0x92, 0x79, 0x64, 0x6f, 0x5a, 0x47, 0xee, 0x1d, 0x64, 0x33, 0xeb, 0xff, 0xbd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x13, 0xbb, 0x3c, 0xc7, 0x36, 0x02, 0x00, 0x00,
}

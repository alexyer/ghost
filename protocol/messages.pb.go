// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Command
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommandId int32

const (
	CommandId_PING CommandId = 0
)

var CommandId_name = map[int32]string{
	0: "PING",
}
var CommandId_value = map[string]int32{
	"PING": 0,
}

func (x CommandId) Enum() *CommandId {
	p := new(CommandId)
	*p = x
	return p
}
func (x CommandId) String() string {
	return proto.EnumName(CommandId_name, int32(x))
}
func (x *CommandId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandId_value, data, "CommandId")
	if err != nil {
		return err
	}
	*x = CommandId(value)
	return nil
}
func (CommandId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Command struct {
	CommandId        *CommandId `protobuf:"varint,1,req,name=commandId,enum=protocol.CommandId" json:"commandId,omitempty"`
	Args             []string   `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetCommandId() CommandId {
	if m != nil && m.CommandId != nil {
		return *m.CommandId
	}
	return CommandId_PING
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "protocol.Command")
	proto.RegisterEnum("protocol.CommandId", CommandId_name, CommandId_value)
}

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9,
	0x39, 0x4a, 0xf6, 0x5c, 0xec, 0xce, 0xf9, 0xb9, 0xb9, 0x89, 0x79, 0x29, 0x42, 0x6a, 0x5c, 0x9c,
	0xc9, 0x10, 0xa6, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x9f, 0x91, 0xb0, 0x1e, 0x4c, 0xa1,
	0x9e, 0x33, 0x4c, 0x4a, 0x88, 0x87, 0x8b, 0x25, 0xb1, 0x28, 0xbd, 0x58, 0x82, 0x49, 0x81, 0x59,
	0x83, 0x53, 0x4b, 0x94, 0x8b, 0x13, 0x21, 0xc5, 0xc1, 0xc5, 0x12, 0xe0, 0xe9, 0xe7, 0x2e, 0xc0,
	0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x16, 0x08, 0xed, 0x72, 0x00, 0x00, 0x00,
}
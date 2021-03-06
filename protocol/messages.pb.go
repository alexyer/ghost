// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Command
	Reply
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
	CommandId_SET  CommandId = 1
	CommandId_GET  CommandId = 2
	CommandId_DEL  CommandId = 3
	CommandId_CGET CommandId = 4
	CommandId_CADD CommandId = 5
)

var CommandId_name = map[int32]string{
	0: "PING",
	1: "SET",
	2: "GET",
	3: "DEL",
	4: "CGET",
	5: "CADD",
}
var CommandId_value = map[string]int32{
	"PING": 0,
	"SET":  1,
	"GET":  2,
	"DEL":  3,
	"CGET": 4,
	"CADD": 5,
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

type Reply struct {
	Values           []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Error            *string  `protobuf:"bytes,2,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Reply) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "protocol.Command")
	proto.RegisterType((*Reply)(nil), "protocol.Reply")
	proto.RegisterEnum("protocol.CommandId", CommandId_name, CommandId_value)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9,
	0x39, 0x4a, 0xf6, 0x5c, 0xec, 0xce, 0xf9, 0xb9, 0xb9, 0x89, 0x79, 0x29, 0x42, 0x6a, 0x5c, 0x9c,
	0xc9, 0x10, 0xa6, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x9f, 0x91, 0xb0, 0x1e, 0x4c, 0xa1,
	0x9e, 0x33, 0x4c, 0x4a, 0x88, 0x87, 0x8b, 0x25, 0xb1, 0x28, 0xbd, 0x58, 0x82, 0x49, 0x81, 0x59,
	0x83, 0x53, 0x49, 0x8d, 0x8b, 0x35, 0x28, 0xb5, 0x20, 0xa7, 0x52, 0x88, 0x8f, 0x8b, 0xad, 0x2c,
	0x31, 0xa7, 0x34, 0xb5, 0x18, 0xa8, 0x17, 0x28, 0x21, 0xc4, 0xcb, 0xc5, 0x9a, 0x5a, 0x54, 0x94,
	0x5f, 0x04, 0x54, 0xc7, 0xa4, 0xc1, 0xa9, 0xe5, 0xc2, 0xc5, 0x89, 0x30, 0x82, 0x83, 0x8b, 0x25,
	0xc0, 0xd3, 0xcf, 0x5d, 0x80, 0x41, 0x88, 0x9d, 0x8b, 0x39, 0xd8, 0x35, 0x44, 0x80, 0x11, 0xc4,
	0x70, 0x07, 0x32, 0x98, 0x40, 0x0c, 0x17, 0x57, 0x1f, 0x01, 0x66, 0x90, 0x22, 0x67, 0x90, 0x10,
	0x0b, 0x98, 0xe5, 0xe8, 0xe2, 0x22, 0xc0, 0x0a, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x14, 0x2b,
	0x8a, 0xc9, 0x00, 0x00, 0x00,
}

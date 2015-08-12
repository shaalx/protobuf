package proto

import (
	"code.google.com/p/goprotobuf/proto"
)

type IMessage struct {
}

func (m *IMessage) Reset()         { *m = IMessage{} }
func (m *IMessage) String() string { return proto.CompactTextString(m) }
func (*IMessage) ProtoMessage()    {}

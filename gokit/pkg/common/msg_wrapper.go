package common

import "google.golang.org/protobuf/reflect/protoreflect"

type ProtoMsgWrapper interface {
	Validate() error
	ProtoReflect() protoreflect.Message
}

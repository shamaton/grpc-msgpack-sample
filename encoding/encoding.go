package encoding

import (
	"google.golang.org/grpc/encoding"
)

var registered = false

func init() {
	RegisterCodec()
}

func RegisterCodec() {
	if registered {
		return
	}
	encoding.RegisterCodec(new(jsonCodec))
	encoding.RegisterCodec(new(msgpackCodec))
	registered = true
}

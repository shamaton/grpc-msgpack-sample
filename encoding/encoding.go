package encoding

import (
	"google.golang.org/grpc/encoding"
)

func RegisterCodec() {
	encoding.RegisterCodec(new(jsonCodec))
	encoding.RegisterCodec(new(msgpackCodec))
}

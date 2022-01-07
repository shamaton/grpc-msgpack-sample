package encoding

import (
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterCodec(new(jsonCodec))
	encoding.RegisterCodec(new(msgpackCodec))
}

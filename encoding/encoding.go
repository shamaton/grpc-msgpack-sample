package encoding

import (
	"fmt"

	"google.golang.org/grpc/encoding"
)

const (
	// JsonCodecName is the name registered for the json encoder.
	JsonCodecName = "json"
)

func init() {
	encoding.RegisterCodec(new(jsonCodec))
	fmt.Println("initid")
}

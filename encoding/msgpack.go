package encoding

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/shamaton/msgpack/v2"
)

const MsgpackCodecName = "msgpack"

// jsonCodec implements encoding.Codec to encode messages into JSON.
type msgpackCodec struct {
}

func (c *msgpackCodec) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("not a proto message but %T: %v", v, v)
	}

	b, err := msgpack.Marshal(msg)
	log.Println("byte size: ", len(b))
	return b, err
}

func (c *msgpackCodec) Unmarshal(data []byte, v interface{}) error {
	_, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("not a proto message but %T: %v", v, v)
	}
	return msgpack.Unmarshal(data, v)
}

// Name returns the identifier of the jsonCodec.
func (c *msgpackCodec) Name() string {
	return MsgpackCodecName
}

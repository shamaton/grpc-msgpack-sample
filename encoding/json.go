package encoding

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const JsonCodecName = "json"

// jsonCodec implements encoding.Codec to encode messages into JSON.
type jsonCodec struct {
	m jsonpb.Marshaler
	u jsonpb.Unmarshaler
}

func (c *jsonCodec) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("not a proto message but %T: %v", v, v)
	}

	var w bytes.Buffer
	if err := c.m.Marshal(&w, msg); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (c *jsonCodec) Unmarshal(data []byte, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("not a proto message but %T: %v", v, v)
	}
	return c.u.Unmarshal(bytes.NewReader(data), msg)
}

// Name returns the identifier of the jsonCodec.
func (c *jsonCodec) Name() string {
	return JsonCodecName
}

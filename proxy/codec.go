package proxy

import (
	"fmt"

	"google.golang.org/grpc/encoding"
	encoding_proto "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/proto"
)

// Codec returns a proxying grpc.Codec with the default protobuf codec as parent.
//
// See CodecWithParent.
func Codec() encoding.Codec {
	return CodecWithParent(&protoCodec{})
}

// CodecWithParent returns a proxying encoding.Codec with a user provided codec as parent.
//
// This codec is *crucial* to the functioning of the proxy. It allows the proxy server to be oblivious
// to the schema of the forwarded message. It basically treats a gRPC message frame as raw bytes.
// However, if the server handler, or the client caller are not proxy-internal functions it will fall back
// to trying to decode the message using a fallback codec.
func CodecWithParent(fallback encoding.Codec) encoding.Codec {
	return &rawCodec{fallback}
}

type rawCodec struct {
	parentCodec encoding.Codec
}

type frame struct {
	payload []byte
}

func (c *rawCodec) Marshal(v interface{}) ([]byte, error) {
	out, ok := v.(*frame)
	if !ok {
		// fallback to the parent codec
		bytes, err := c.parentCodec.Marshal(v)
		if err != nil {
			return bytes, fmt.Errorf("failed to marshal(codec:%s): %w", c.Name(), err)
		}

		return bytes, nil
	}

	return out.payload, nil
}

func (c *rawCodec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*frame)
	if !ok {
		// fallback to the parent codec
		err := c.parentCodec.Unmarshal(data, v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal(codec:%s): %w", c.Name(), err)
		}

		return nil
	}

	dst.payload = data

	return nil
}

func (c *rawCodec) Name() string {
	return fmt.Sprintf("proxy>%s", c.parentCodec.Name())
}

// protoCodec is a Codec implementation with protobuf. It's the default rawCodec for gRPC.
type protoCodec struct{}

func (p protoCodec) Marshal(v interface{}) ([]byte, error) {
	ret, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		return ret, fmt.Errorf("failed to marshal(codec:%s): %w", p.Name(), err)
	}

	return ret, nil
}

func (p protoCodec) Unmarshal(data []byte, v interface{}) error {
	err := proto.Unmarshal(data, v.(proto.Message))
	if err != nil {
		return fmt.Errorf("failed to unmarshal(codec:%s): %w", p.Name(), err)
	}

	return nil
}

func (protoCodec) Name() string {
	return encoding_proto.Name
}

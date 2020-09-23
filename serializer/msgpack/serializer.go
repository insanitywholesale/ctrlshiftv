package json

import (
	"ctrlshiftv/paste"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Paste struct{}

func (p *Paste) Decode(input []byte) (*paste.Paste, error) {
	paste := &paste.Paste{}
	err := msgpack.Unmarshal(input, paste)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return paste, nil
}

func (p *Paste) Encode(input *paste.Paste) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}

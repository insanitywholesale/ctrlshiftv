package json

import (
	"ctrlshiftv/paste"
	"encoding/json"
	"github.com/pkg/errors"
)

type Paste struct{}

func (p *Paste) Decode(input []byte) (*paste.Paste, error) {
	paste := &paste.Paste{}
	err := json.Unmarshal(input, paste)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Paste.Decode")
	}
	return paste, nil
}

func (p *Paste) Encode(input *paste.Paste) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Paste.Encode")
	}
	return rawMsg, nil
}

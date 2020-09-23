package xml

import (
	"ctrlshiftv/paste"
	"encoding/xml"
	"github.com/pkg/errors"
)

type Paste struct{}

func (p *Paste) Decode(input []byte) (*paste.Paste, error) {
	paste := &paste.Paste{}
	err := xml.Unmarshal(input, paste)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return paste, nil
}

func (p *Paste) Encode(input *paste.Paste) ([]byte, error) {
	rawMsg, err := xml.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}

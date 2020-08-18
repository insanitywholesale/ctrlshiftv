package plain

import (
	"ctrlshiftv/paste"
	//"github.com/pkg/errors"
)

type Paste struct{}

func (p *Paste) Decode(input []byte) (*paste.Paste, error) {
	paste := &paste.Paste{}
	paste.Content = string(input)
	return paste, nil
}

func (p *Paste) Encode(input *paste.Paste) ([]byte, error) {
	rawContent := []byte(input.Content)
	return rawContent, nil
}

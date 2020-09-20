package plain

import (
	"ctrlshiftv/paste"
	"strconv"
)

type Paste struct{}

func (p *Paste) Decode(input []byte) (*paste.Paste, error) {
	paste := &paste.Paste{}
	paste.Content = string(input)
	return paste, nil
}

func (p *Paste) Encode(input *paste.Paste) ([]byte, error) {
	// the mess below is for better output to terminal when raw data is sent
	// might remove later cause it's kinda useless, the link is what matters
	rawContent := "code:" + input.Code + "\ncontent:" + input.Content + "\ncreatedat:" + strconv.FormatInt(input.CreatedAt, 10) + "\n"
	return []byte(rawContent), nil
}

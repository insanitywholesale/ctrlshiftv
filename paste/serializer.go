package paste

type PasteSerializer interface {
	Decode(input []byte) (*Paste, error)
	Encode(input *Paste) ([]byte, error)
}

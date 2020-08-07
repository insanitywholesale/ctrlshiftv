package paste

type PasteService interface {
	Find(code string) (*Paste, error)
	Store(paste *Paste) error
}

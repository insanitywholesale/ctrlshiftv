package paste

type PasteRepo interface {
	Find(code string) (*Paste, error)
	Store(paste *Paste) error
}

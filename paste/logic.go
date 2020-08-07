package paste

import (
	"errors"
)

var (
	ErrPasteNotFound = errors.New("Paste Not Found")
	ErrPasteInvalid  = errors.New("Paste Invalid")
)

type pasteService struct {
	pasteRepo PasteRepo
}

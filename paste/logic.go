package paste

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	//"strconv"
	"time"
)

var (
	ErrPasteNotFound = errors.New("Paste Not Found")
	ErrPasteInvalid  = errors.New("Paste Invalid")
)

type pasteService struct {
	pasteRepo PasteRepo
}

func NewPasteService(pasteRepo PasteRepo) PasteService {
	return &pasteService{
		pasteRepo,
	}
}

func (p *pasteService) Find(code string) (*Paste, error) {
	return p.pasteRepo.Find(code)
}

func (p *pasteService) Store(paste *Paste) error {
	err := validate.Validate(paste)
	if err != nil {
		return errs.Wrap(ErrPasteInvalid, "service.Paste")
	}
	paste.Code = shortid.MustGenerate()
	paste.CreatedAt = time.Now().UTC().Unix()
	return p.pasteRepo.Store(paste)
}

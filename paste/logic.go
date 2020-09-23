package paste

import (
	"context"
	protos "ctrlshiftv/proto/shorten"
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"google.golang.org/grpc"
	"gopkg.in/dealancer/validate.v2"
	"log"
	"time"
)

var (
	ErrPasteNotFound = errors.New("Paste Not Found")
	ErrPasteInvalid  = errors.New("Paste Invalid")
)

type pasteService struct {
	pasteRepo PasteRepo
}

var grpconn *grpc.ClientConn

func SaveConn(conn *grpc.ClientConn) {
	grpconn = conn
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
	// this works now so better fix functionality
	client := protos.NewShortenRequestClient(grpconn)
	shortLink, err := client.GetShortURL(context.Background(), &protos.LongLink{
		Link: "https://distro.watch",
	})
	log.Println("shortlink", shortLink)
	if err != nil {
		log.Fatalf("Failed to get short link code: %v", err)
	}
	// see previous comment

	err = validate.Validate(paste)
	if err != nil {
		return errs.Wrap(ErrPasteInvalid, "service.Paste")
	}

	paste.Code = shortid.MustGenerate()
	paste.CreatedAt = time.Now().UTC().Unix()
	return p.pasteRepo.Store(paste)
}

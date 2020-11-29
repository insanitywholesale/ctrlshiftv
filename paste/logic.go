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

var grpclient protos.ShortenRequestClient

func SaveClient(client protos.ShortenRequestClient) {
	grpclient = client
}

//TODO: probably remove this since we can access the client now
var grpconn *grpc.ClientConn

func SaveConn(conn *grpc.ClientConn) {
	grpconn = conn
}

func (p *pasteService) MakeShortURL(url string) string { //*protos.ShortLink {
	shortLink, err := grpclient.GetShortURL(context.Background(), &protos.LongLink{
		Link: url,
	})
	if err != nil {
		log.Println("SaveClient error:", err)
	}
	//TODO: remove ignore thing
	log.Println("IGNORE shortlink", shortLink)
	return shortLink.Link
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

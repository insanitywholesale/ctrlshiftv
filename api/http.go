package api

import (
	"ctrlshiftv/paste"
	bs "ctrlshiftv/serializer/bare"
	js "ctrlshiftv/serializer/json"
	plain "ctrlshiftv/serializer/plain"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ErrPasteNotFound = errors.New("Paste Not Found")
	ErrPasteInvalid  = errors.New("Paste Invalid")
)

type PasteHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	pasteService paste.PasteService
}

var pasteboi paste.PasteService

func NewHandler(pasteService paste.PasteService) PasteHandler {
	pasteboi = pasteService
	return &handler{pasteService: pasteService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) paste.PasteSerializer {
	switch contentType {
	// json
	case "application/json":
		return &js.Paste{}
	// bare
	case "application/octet-stream":
		return &bs.Paste{}
	// plain
	case "text/plain":
		return &plain.Paste{}
	case "application/x-www-form-urlencoded":
		return &plain.Paste{}
	case "multipart/form-data":
		return &plain.Paste{}
	default:
		return nil
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	paste, err := h.pasteService.Find(code)
	if err != nil {
		if errors.Cause(err) == ErrPasteNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//http.Redirect(w, r, paste.Code, http.StatusMovedPermanently)
	// TODO: change this to output the full object in json/msgpack too
	setupResponse(w, r.Header.Get("Content-Type"), []byte(paste.Content+"\n"), http.StatusOK)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	paste, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.pasteService.Store(paste)
	if err != nil {
		if errors.Cause(err) == ErrPasteInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// TODO: add PROPER gRPC call to urlshort
	// following is just a demo, need to integrate properly
	// TODO: find way to automatically determine baseurl
	myurl := "http://localhost:8080/" + paste.Code
	result := pasteboi.MakeShortURL(myurl)
	log.Println("see paste:", "http://localhost:8000/"+result)
	responseBody, err := h.serializer(contentType).Encode(paste)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

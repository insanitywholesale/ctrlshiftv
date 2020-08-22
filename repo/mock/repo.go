package mock

import (
	"ctrlshiftv/paste"
	//"fmt"
	"errors"
)

type mockRepo []*paste.Paste

var pasteList mockRepo = []*paste.Paste{}

func mkPasteList() mockRepo {
	var pasteList mockRepo = []*paste.Paste{
		&paste.Paste{
			Code:      "0123",
			Content:   "pasterooni",
			CreatedAt: 3254,
		},
		&paste.Paste{
			Code:      "1234",
			Content:   "excerpt goes here",
			CreatedAt: 7873,
		},
	}
	return pasteList
}

func NewMockRepo() (paste.PasteRepo, error) {
	pasteList = mkPasteList()
	repo := &mockRepo{}
	return repo, nil
}

func (r *mockRepo) Find(code string) (*paste.Paste, error) {
	for _, paste := range pasteList {
		if paste.Code == code {
			return paste, nil
		}
	}
	return nil, errors.New("Paste Not Found") //fmt.Errorf("paste with code %s was not found", code)
}

func (r *mockRepo) Store(pp *paste.Paste) error {
	pasteList = append(pasteList, pp)
	return nil
}

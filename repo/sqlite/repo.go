package sqlite

import (
	"ctrlshiftv/paste"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// TODO: implement Find and Store

var createDB = "create table if not exists pastes (code text, content text, createdat text)"

type sqliteRepo struct {
	client   *sql.DB // the connection to the database
	database string  // the location of the database
}

func newSQLiteClient(location string) (*sql.DB, error) {
	client, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewSQLiteRepo(location string) (paste.PasteRepo, error) {
	repo := &sqliteRepo{
		database: location,
	}
	client, err := newSQLiteClient(location)
	if err != nil {
		return nil, errors.Wrap(err, "repo.NewSQLiteRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *sqliteRepo) Find(code string) (*paste.Paste, error) {
	fmt.Println("hello from Find method")
	return &paste.Paste{}, nil
}

func (r *sqliteRepo) Store(paste *paste.Paste) error {
	fmt.Println("hello from Store method")
	return nil
}

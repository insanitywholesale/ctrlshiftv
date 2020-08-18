package postgres

import (
	"ctrlshiftv/paste"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// TODO: implement Find and Store
// TODO: add "create table if not exists"
// TODO: check about using pgx instead of pq

type postgresRepo struct {
	client   *sql.DB
	database string
}

func newPostgresClient(postgresURL string) (*sql.DB, error) {
	client, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresRepo(postgresURL string) (paste.PasteRepo, error) {
	repo := &postgresRepo{
		database: postgresURL,
	}
	client, err := newPostgresClient(postgresURL)
	if err != nil {
		return nil, errors.Wrap(err, "repo.NewPostgresRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *postgresRepo) Find(code string) (*paste.Paste, error) {
	fmt.Println("hello from Find method")
	return &paste.Paste{}, nil
}

func (r *postgresRepo) Store(paste *paste.Paste) error {
	fmt.Println("hello from Store method")
	return nil
}

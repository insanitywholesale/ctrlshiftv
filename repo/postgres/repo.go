package postgres

import (
	"ctrlshiftv/paste"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type postgresRepo struct {
	client   *sql.DB
	database string
}

var createtable = `CREATE TABLE if not exists pastes (
	code TEXT,
	content TEXT
);`

func newPostgresClient(postgresURL string) (*sql.DB, error) {
	client, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createtable)
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
	query := `SELECT * FROM pastes where code=$1`
	row := r.client.QueryRow(query, code)
	fmt.Println("row", row)
	var paste paste.Paste
	err := row.Scan(&paste.Code, &paste.Content)
	if err != nil {
		return nil, err
	}
	fmt.Println("hello from Find method")
	return &paste, nil
}

func (r *postgresRepo) Store(paste *paste.Paste) error {
	query := `INSERT INTO pastes (code, content) VALUES ($1, $2)`
	_, err := r.client.Exec(query, paste.Code, paste.Content)
	if err != nil {
		return errors.Wrap(err, "repo.Store")
	}
	fmt.Println("hello from Store method")
	return nil
}

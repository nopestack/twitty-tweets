package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nopestack/twitty/core"
	"github.com/pkg/errors"
)

type mysqlRepository struct {
	client      *sql.Conn
	timeout     time.Duration
	databaseURL string
}

func newMySQLClient(connString string, timeout int) (*sql.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		// Error opening the db at the provided URI
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// Somehow, the DB is not available
		return nil, err
	}

	client, connErr := db.Conn(ctx)

	if connErr != nil {
		// Error establishing a long-lived connection
		return nil, connErr
	}

	return client, nil
}

func NewMySQLRepository(mysqlURI string, timeout int) (*mysqlRepository, error) {

	repo := &mysqlRepository{
		// We need the timeout here for use on query cancellations later on
		timeout:     time.Duration(timeout) * time.Second,
		databaseURL: mysqlURI,
	}

	client, err := newMySQLClient(mysqlURI, timeout)

	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMySQLRepository")
	}

	repo.client = client

	return repo, nil
}

func (r *mysqlRepository) Find(id int) (*core.Tweet, error)    {}
func (r *mysqlRepository) Create(tweet *core.Tweet) error      {}
func (r *mysqlRepository) Update(id int, content string) error {}
func (r *mysqlRepository) Delete(id int) error                 {}

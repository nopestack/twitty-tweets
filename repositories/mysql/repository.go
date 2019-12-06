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

func (r *mysqlRepository) All() ([]*core.Tweet, error) {
	queryString := "select id, content, author, created_at, updated_at from tweets"

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	tweets := []*core.Tweet{}

	results, err := r.client.QueryContext(ctx, queryString)
	if err != nil {
		errors.Wrap(err, "repository.Tweets.All")
	}

	for results.Next() {
		singleTweet := &core.Tweet{}

		err = results.Scan(
			&singleTweet.ID,
			&singleTweet.Content,
			&singleTweet.Author,
			&singleTweet.CreatedAt,
			&singleTweet.UpdatedAt,
		)

		if err != nil {
			errors.Wrap(err, "repository.Tweets.All")
		}

		tweets = append(tweets, singleTweet)
	}

	return tweets, nil
}

func (r *mysqlRepository) Find(id int) (*core.Tweet, error) {
	queryString := "select id, name, content, author, created_at, updated_at from tweets where id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	result := r.client.QueryRowContext(ctx, queryString, 1)

	tweet := &core.Tweet{}

	err := result.Scan(
		&tweet.ID,
		&tweet.Content,
		&tweet.Author,
		&tweet.CreatedAt,
		&tweet.UpdatedAt,
	)

	if err != nil {
		errors.Wrap(err, "repository.Tweets.Find")
		return nil, err
	}

	return tweet, nil
}

func (r *mysqlRepository) Create(tweet *core.Tweet) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	queryString := "insert into tweets (content, author, created_at) values (?, ?, ?)"

	_, err := r.client.ExecContext(ctx, queryString,
		tweet.Content,
		tweet.Author,
		time.Now().Unix(),
	)

	if err != nil {
		errors.Wrap(err, "repository.tweets.Create")
		return err
	}

	return nil
}

func (r *mysqlRepository) Update(id int, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	queryString := "update tweets set content = ?, updated_at = ? where id = ?"

	_, err := r.client.ExecContext(ctx, queryString,
		content,
		time.Now().Unix(),
		id,
	)

	if err != nil {
		errors.Wrap(err, "repository.tweets.Update")
		return err
	}

	return nil
}

func (r *mysqlRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	queryString := "delete from tweets where id = ?"

	_, err := r.client.ExecContext(ctx, queryString, 9)

	if err != nil {
		errors.Wrap(err, "repository.tweets.Delete")
		return err
	}

	return nil
}

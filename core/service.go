package core

import (
	errs "github.com/pkg/errors"
)

var (
	ErrTweetNotFound = errs.New("Tweet not found")
	ErrInvalidTweet  = errs.New("Invalid tweet")
)

type TweetService struct {
	repository Repository
}

// NewTweetService returns an instance of the tweeting service
func NewTweetService(repository Repository) *TweetService {
	return &TweetService{
		repository,
	}

}

func (ts *TweetService) FindByID(id int) (*Tweet, error) {
	// Finds a tweet by ID
	/*
		No input validation needed
		- if error on Find, return a TweetNotFound Error
	*/
	return ts.repository.Find(id)
}

func (ts *TweetService) GetAll(author string) ([]*Tweet, error) {
	// Fetches every tweet?
	return nil, nil
}

func (ts *TweetService) AddTweet(tweet *Tweet) error {
	// Validation logic goes here
	/*
		- check the provided input tweet, e.g. empty string set as content, unparseable dates, empty string author

		- if error, then return an invalid input error
		else create tweet
	*/
	return ts.repository.Create(tweet)
}

func (ts *TweetService) UpdateTweet(id int, content string) error {
	// Input validation here
	/*
		- search for tweet by id
		- if tweet exists, then call Update
		- if not, return a tweet not found error
	*/
	_, err := ts.repository.Find(id)

	if err != nil {
		// Return a tweet not found error
		return errs.Wrap(ErrTweetNotFound, "core.service.UpdateTweet")
	}

	return ts.repository.Update(id, content)
}

func (ts *TweetService) DeleteTweet(id int) error {
	// input validation here
	/*
		- search for tweet by id
		- if tweet exists, then call Delete
		- if not, return a tweet not found error
	*/
	_, err := ts.repository.Find(id)

	if err != nil {
		// Return a tweet not found error
		return errs.Wrap(ErrTweetNotFound, "core.service.DeleteTweet")
	}

	return ts.repository.Delete(id)
}

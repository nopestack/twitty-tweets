package core

type tweetService struct {
	repository Repository
}

func NewTweetService(repository Repository) *tweetService {
	return &tweetService{
		repository,
	}

}

func (ts *tweetService) FindByID(id int) (*Tweet, error) {
	// Finds a tweet by ID
	return ts.repository.Find(id)
}

func (ts *tweetService) GetAll(author string) ([]*Tweet, error) {
	// Fetches every tweet by author
	return nil, nil
}

func (ts *tweetService) AddTweet(redirect *Tweet) error {

	// Validation logic goes here

	return ts.repository.Create(redirect)
}

func (ts *tweetService) Update(id int, content string) error {
	return nil

}

func (ts *tweetService) Delete(id int) error {
	return nil
}

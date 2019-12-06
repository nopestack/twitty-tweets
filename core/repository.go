package core

type Repository interface {
	All() ([]*Tweet, error)
	Find(id int) (*Tweet, error)
	Create(tweet *Tweet) error
	Update(id int, content string) error
	Delete(id int) error
}

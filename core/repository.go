package core

type Repository interface {
	Find(id int) (*Tweet, error)
	Create(tweet *Tweet) error
	Update(id int, content string) error
	Delete(id int) (Tweet, error)
}

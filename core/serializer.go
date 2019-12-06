package core

// TweetSerializer represents a serializer who takes in byte and turns them into tweets and viceversa
type TweetSerializer interface {
	Decode(input []byte) (*Tweet, error)
	Encode(input *Tweet) ([]byte, error)
}

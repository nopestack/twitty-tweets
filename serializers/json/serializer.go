package json

import (
	"encoding/json"

	"github.com/nopestack/twitty/core"
	"github.com/pkg/errors"
)

// Implements the TweetSerializer interface
type Serializer struct{}

func (js *Serializer) Encode(input *core.Tweet) ([]byte, error) {
	encoded, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializers.json.Serializer.Encode")
	}
	return encoded, nil
}

func (js *Serializer) Decode(input []byte) (*core.Tweet, error) {
	tweet := &core.Tweet{}
	if err := json.Unmarshal(input, tweet); err != nil {
		return nil, errors.Wrap(err, "serializers.json.Serializer.Decode")
	}

	return tweet, nil
}

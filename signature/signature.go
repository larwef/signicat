package signature

import (
	"errors"
	"github.com/larwef/signicat"
)

var ErrNotImplemented = errors.New("not implemented")

type Signature struct {
	client *signicat.Client
}

func New(client *signicat.Client) *Signature {
	return &Signature{
		client: client,
	}
}

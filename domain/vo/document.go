package vo

import (
	"errors"
	"fmt"
	"github.com/klassmann/cpfcnpj"
)

type Document struct {
	number string
}

func NewDocument(n string) (*Document, error) {
	validated := cpfcnpj.ValidateCPF(n)

	if validated {
		return &Document{n}, nil
	}

	return nil, errors.New(fmt.Sprintf("Invalid value for Document(%s)", n))
}

func (d Document) Number() string {
	return d.number
}

package vo

import (
	"errors"
	"fmt"
)

type Id struct {
	value uint64
}

func NewId(value uint64) (*Id, error) {
	if value < 1 {
		return nil, errors.New(fmt.Sprintf("Invalid value for Id(%d)", value))
	}

	return &Id{value: value}, nil
}

func (i Id) Value() uint64 {
	return i.value
}

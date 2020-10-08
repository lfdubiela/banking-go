package vo

import (
	"errors"
	"fmt"
	"math"
)

//should be using decimal here
//https://github.com/shopspring/decimal
type Amount struct {
	value float64
}

func NewAmount(n float64) (*Amount, error) {
	if n < 1 {
		return nil, errors.New(fmt.Sprintf("Invalid value for Amount(%g)", n))
	}

	return &Amount{math.Trunc(n*100) / 100}, nil
}

func (a Amount) Value() float64 {
	return a.value
}

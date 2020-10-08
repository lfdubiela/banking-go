package vo

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		param       uint64
		result      *Id
		err         error
	}{
		{
			description: "Invalid value for Id(%d)",
			param:       0,
			result:      nil,
			err:         errors.New("Invalid value for Id(0)"),
		},
		{
			description: "Valid value for Id(%d)",
			param:       1,
			result:      &Id{1},
			err:         nil,
		},
		{
			description: "Valid value for Id(%d)",
			param:       1000000,
			result:      &Id{1000000},
			err:         nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.param), func(t *testing.T) {
			id, err := NewId(test.param)

			if !equalId(test.result, id) {
				t.Errorf("Expected %v, got %v", test.result, id)
			}

			if !equalError(test.err, err) {
				t.Errorf("Expected: %v, got: %v", test.err, err)
			}
		})
	}
}

func equalId(a, b *Id) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Value() == b.Value()
}

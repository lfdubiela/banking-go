package vo

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		param       float64
		result      *Amount
		err         error
	}{
		{
			description: "Invalid value for Amount(%v)",
			param:       -1,
			result:      nil,
			err:         errors.New("Invalid value for Amount(-1)"),
		},
		{
			description: "Invalid value for Amount(%v)",
			param:       -1000,
			result:      nil,
			err:         errors.New("Invalid value for Amount(-1000)"),
		},
		{
			description: "Valid value for Amount(%v)",
			param:       10.01,
			result:      &Amount{10.01},
			err:         nil,
		},
		{
			description: "Valid value for Amount(%v)",
			param:       10.011,
			result:      &Amount{10.01},
			err:         nil,
		},
		{
			description: "Valid value for Amount(%v)",
			param:       10.019,
			result:      &Amount{10.01},
			err:         nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.param), func(t *testing.T) {
			amount, err := NewAmount(test.param)

			if !equalAmount(test.result, amount) {
				t.Errorf("Expected %v, got %v", test.result, amount)
			}

			if !equalError(test.err, err) {
				t.Errorf("Expected: %d, got: %d", test.err, err)
			}
		})
	}
}

func equalError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}

func equalAmount(a, b *Amount) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Value() == b.Value()
}

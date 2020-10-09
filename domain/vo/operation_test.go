package vo

import (
	"errors"
	"fmt"
	"testing"
)

func TestOperationParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		param       uint8
		result      *Operation
		err         error
	}{
		{
			description: "Parse to a valid Operation(%d)",
			param:       1,
			result:      &Operations.CompraAVista,
			err:         nil,
		},
		{
			description: "Parse to a valid Operation(%d)",
			param:       2,
			result:      &Operations.CompraParcelada,
			err:         nil,
		},
		{
			description: "Parse to a valid Operation(%d)",
			param:       3,
			result:      &Operations.Saque,
			err:         nil,
		},
		{
			description: "Try parse to a invalid Operation(%d)",
			param:       5,
			result:      nil,
			err:         errors.New("Invalid value for Operation(5)"),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.param), func(t *testing.T) {
			operation, err := Operations.Parse(test.param)

			if !equalOperation(test.result, operation) {
				t.Errorf("Expected %v, got %v", test.result, operation)
			}

			if !equalError(test.err, err) {
				t.Errorf("Expected: %v, got: %v", test.err, err)
			}
		})
	}
}

func TestOperationList(t *testing.T) {
	t.Parallel()

	t.Run("List all operations", func(t *testing.T) {
		operations := Operations.List()
		length := len(operations)

		if length != 4 {
			t.Error(fmt.Sprintf("Expected 4 operations, got %d", length))
		}
	})
}

func equalOperation(a, b *Operation) bool {
	return a == nil &&
		b == nil || a != nil &&
		b != nil &&
		a.Id() == b.Id() &&
		a.Description() == b.Description() &&
		a.Mode() == b.mode
}

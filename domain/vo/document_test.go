package vo

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewDocument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		param       string
		result      *Document
		err         error
	}{
		{
			description: "Invalid value for Document(%s)",
			param:       "12345678901",
			result:      nil,
			err:         errors.New("Invalid value for Document(12345678901)"),
		},
		{
			description: "Invalid value for Document(%s)",
			param:       "123456789",
			result:      nil,
			err:         errors.New("Invalid value for Document(123456789)"),
		},
		{
			description: "Invalid value for Document(%s)",
			param:       "123456789",
			result:      nil,
			err:         errors.New("Invalid value for Document(123456789)"),
		},
		{
			description: "Invalid value for Document(%s)",
			param:       "896.688.868-28",
			result:      nil,
			err:         errors.New("Invalid value for Document(896.688.868-28)"),
		},
		{
			description: "Valid value for Document(%s)",
			param:       "01714126137",
			result:      &Document{"01714126137"},
			err:         nil,
		},
		{
			description: "Valid value for Document(%s)",
			param:       "08411504948",
			result:      &Document{"08411504948"},
			err:         nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.param), func(t *testing.T) {
			document, err := NewDocument(test.param)

			if !equalDocument(test.result, document) {
				t.Errorf("Expected %v, got %v", test.result, document)
			}

			if !equalError(test.err, err) {
				t.Errorf("Expected: %v, got: %v", test.err, err)
			}
		})
	}
}

func equalDocument(a, b *Document) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Number() == b.Number()
}

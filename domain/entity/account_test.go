package entity

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/lfdubiela/banking-go/domain/vo"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description    string
		documentNumber string
		expected       *Account
		err            error
	}{
		{
			description:    "Valid values for Account(document:%)",
			documentNumber: "56295269443",
			expected: &Account{
				nil,
				newDocument("56295269443"),
			},
			err: nil,
		},
		{
			description:    "Invalid values for Account(document:%)",
			documentNumber: "12345678901",
			expected:       nil,
			err:            errors.New("Invalid value for Document(12345678901)"),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.documentNumber), func(t *testing.T) {
			account, err := NewAccount(test.documentNumber)

			if !reflect.DeepEqual(test.err, err) {
				t.Errorf("Error expected: %s, received: %s", test.err, err)
				return
			}

			if err != nil {
				return
			}

			if !reflect.DeepEqual(test.expected.Document(), account.Document()) {
				t.Errorf("Document() expected %v, received %v", test.expected, account.Id())
			}
		})
	}
}

func TestAccountWithId(t *testing.T) {
	t.Parallel()

	type params struct {
		id       uint64
		document string
	}

	tests := []struct {
		description string
		params      params
		expected    *Account
	}{
		{
			description: "Valid values for Account(id:%)",
			params:      params{1, "56295269443"},
			expected: &Account{
				newId(1),
				newDocument("56295269443"),
			},
		},
		{
			description: "Valid values for Account(id:%)",
			params:      params{1, "56295269443"},
			expected: &Account{
				newId(1),
				newDocument("56295269443"),
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.params.id), func(t *testing.T) {
			account, _ := NewAccount(test.params.document)
			account = account.WithId(newId(test.params.id))

			if !reflect.DeepEqual(test.expected.Id(), account.Id()) {
				t.Errorf("Id() expected %v, received %v", test.expected, account.Id())
			}

			if !reflect.DeepEqual(test.expected.Document(), account.Document()) {
				t.Errorf("Document() expected %v, received %v", test.expected, account.Id())
			}
		})
	}
}

func TestFindAccountSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Test FindAccount returning success", func(t *testing.T) {
		testAccount := &Account{
			newId(1),
			newDocument("56295269443"),
		}

		finder := &accountFinderStub{testAccount}

		id, _ := vo.NewId(1)

		account, _ := FindAccount(finder, id)

		if !reflect.DeepEqual(testAccount, account) {
			t.Errorf("Expected %v, got %v", testAccount, account)
		}
	})
}

func TestFindAccountFailed(t *testing.T) {
	t.Parallel()

	t.Run("Test FindAccount returning failed", func(t *testing.T) {
		expectedError := errors.New("Account not found!")
		finder := &accountFinderStub{expectedError}

		_, err := FindAccount(finder, newId(1))

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})
}

func TestSaveAccountSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning success", func(t *testing.T) {
		expected := &Account{
			newId(1),
			newDocument("56295269443"),
		}

		saver := &accountSaverStub{
			newId(1),
		}

		account, _ := expected.Save(saver)

		if !reflect.DeepEqual(expected, account) {
			t.Errorf("Expected %v, got %v", expected, account)
		}
	})
}

func TestSaveAccountFailed(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning failed", func(t *testing.T) {
		account := &Account{
			newId(1),
			newDocument("56295269443"),
		}

		expected := errors.New("Account already exists with document(56295269443)!")

		saver := &accountSaverStub{expected}

		_, err := account.Save(saver)

		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Expected %v, got %v", expected, err)
		}
	})
}

type accountFinderStub struct {
	result interface{}
}

func (a *accountFinderStub) Find(id *vo.Id) (*Account, error) {
	if err, ok := a.result.(error); ok {
		return nil, err
	}
	account, _ := a.result.(*Account)
	return account, nil
}

type accountSaverStub struct {
	result interface{}
}

func (s *accountSaverStub) Save(a *Account) (*vo.Id, error) {
	if err, ok := s.result.(error); ok {
		return nil, err
	}

	id, _ := s.result.(*vo.Id)

	return id, nil
}

func newId(i uint64) *vo.Id {
	id, _ := vo.NewId(i)
	return id
}

func newDocument(n string) vo.Document {
	doc, _ := vo.NewDocument(n)
	return *doc
}

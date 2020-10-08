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
			expected:       createAccount("56295269443"),
			err:            nil,
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

			if !reflect.DeepEqual(test.expected, account) {
				t.Errorf("Expected %v, got %v", test.expected, account)
			}

			if !equalError(test.err, err) {
				t.Errorf("Expected: %d, got: %d", test.err, err)
			}
		})
	}
}

func TestWithId(t *testing.T) {
	t.Parallel()

	type params struct {
		id       *vo.Id
		document string
	}

	tests := []struct {
		description string
		params      params
		account     *Account
		err         error
	}{
		{
			description: "Valid values for Account(id:%)",
			params:      params{createId(1), "56295269443"},
			account:     createAccountAllParams("56295269443", 1),
			err:         nil,
		},
		{
			description: "Valid values for Account(id:%)",
			params:      params{createId(1000), "56295269443"},
			account:     createAccountAllParams("56295269443", 1000),
			err:         nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.params.id), func(t *testing.T) {
			account, _ := NewAccount(test.params.document)
			account = account.WithId(test.params.id)

			if !reflect.DeepEqual(account, test.account) {
				t.Errorf("Expected %v, got %v", test.account, account)
			}
		})
	}
}

func TestFindAccountSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Test FindAccount returning success", func(t *testing.T) {
		testAccount := createAccountAllParams("56295269443", 1)
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

		id, _ := vo.NewId(1)

		_, err := FindAccount(finder, id)

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})
}

func TestSaveAccountSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning success", func(t *testing.T) {
		expected := createAccountAllParams("56295269443", 1)

		id, _ := vo.NewId(1)
		saver := &accountSaverStub{id}

		account, _ := expected.Save(saver)

		if !reflect.DeepEqual(expected, account) {
			t.Errorf("Expected %v, got %v", expected, account)
		}
	})
}

func TestSaveAccountFailed(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning failed", func(t *testing.T) {
		account := createAccountAllParams("56295269443", 1)
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

func createAccount(document string) *Account {
	d, _ := vo.NewDocument(document)
	return &Account{nil, *d}
}

func createAccountAllParams(document string, id uint64) *Account {
	d, _ := vo.NewDocument(document)
	i, _ := vo.NewId(id)
	return &Account{i, *d}
}

func createId(id uint64) *vo.Id {
	i, _ := vo.NewId(id)
	return i
}

func equalError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}

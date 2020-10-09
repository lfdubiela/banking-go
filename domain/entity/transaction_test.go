package entity

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/lfdubiela/banking-go/domain/vo"
)

func TestNewTransaction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		params      transactionParams
		expected    *Transaction
		err         error
	}{
		{
			description: "Valid values for Transaction",
			params: transactionParams{
				accountId:   1,
				operationId: 1,
				amountValue: 100.11,
				repository: &accountFinderStub{
					&Account{
						newId(1),
						newDocument("30784553513"),
					},
				},
			},
			expected: &Transaction{
				id: nil,
				account: Account{
					newId(1),
					newDocument("30784553513"),
				},
				operation: newOperation(1),
				amount:    newAmount(100.11),
				eventDate: time.Now().UTC(),
			},
			err: nil,
		},
		{
			description: "Invalid accountId value for Id",
			params: transactionParams{
				accountId:   0,
				operationId: 1,
				amountValue: 100.00,
				repository: &accountFinderStub{
					&Account{
						newId(1),
						newDocument("30784553513"),
					},
				},
			},
			expected: nil,
			err:      errors.New("Invalid value for Id(0)"),
		},
		{
			description: "Invalid operation value for Transaction",
			params: transactionParams{
				accountId:   1,
				operationId: 5,
				amountValue: 100.11,
				repository: &accountFinderStub{
					&Account{
						newId(1),
						newDocument("30784553513"),
					},
				},
			},
			expected: nil,
			err:      errors.New("Invalid value for Operation(5)"),
		},
		{
			description: "Invalid amount value for Transaction",
			params: transactionParams{
				accountId:   1,
				operationId: 1,
				amountValue: 0,
				repository: &accountFinderStub{
					&Account{
						newId(1),
						newDocument("30784553513"),
					},
				},
			},
			expected: nil,
			err:      errors.New("Invalid value for Amount(0)"),
		},
		{
			description: "Transaction could not find Account",
			params: transactionParams{
				accountId:   1,
				operationId: 1,
				amountValue: 100.11,
				repository: &accountFinderStub{
					errors.New("Account could not be found!"),
				},
			},
			expected: nil,
			err:      errors.New("Account could not be found!"),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.description, test.params), func(t *testing.T) {
			transaction, err := NewTransction(test.params.accountId,
				test.params.operationId,
				test.params.amountValue,
				test.params.repository)

			if !reflect.DeepEqual(test.err, err) {
				t.Errorf("Error expected: %d, received: %d", test.err, err)
				return
			}

			if err != nil {
				return
			}

			if !reflect.DeepEqual(test.expected.Id(), transaction.Id()) {
				t.Errorf("Id() expected %v, received %v", test.expected, transaction)
			}

			if !reflect.DeepEqual(test.expected.Account(), transaction.Account()) {
				t.Errorf("Account() expected: %v, received: %v", test.expected, transaction)
			}

			if !reflect.DeepEqual(test.expected.Amount(), transaction.Amount()) {
				t.Errorf("Amount() expected: %v, received: %v", test.expected, transaction)
			}

			if !reflect.DeepEqual(test.expected.Operation(), transaction.Operation()) {
				t.Errorf("Operation() expected: %v, received: %v", test.expected, transaction)
			}

			if !reflect.DeepEqual(test.expected.EventDate(), transaction.EventDate()) {
				t.Errorf("EventDate() expected: %v, received: %v", test.expected.EventDate(), transaction.EventDate())
			}
		})
	}
}

func TestTransactionWithId(t *testing.T) {
	t.Parallel()

	t.Run(fmt.Sprintf("Transaction test WithId"), func(t *testing.T) {
		expected := newId(1)

		transaction := &Transaction{
			account: Account{
				newId(1),
				newDocument("30784553513"),
			},
			operation: newOperation(1),
			amount:    newAmount(100.11),
			eventDate: time.Now().UTC(),
		}

		transaction = transaction.WithId(expected)

		if !reflect.DeepEqual(expected, transaction.Id()) {
			t.Errorf("Id() expected %v, received: %v", expected, transaction.Id())
		}
	})
}

func TestSaveTransactionSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning success", func(t *testing.T) {
		expected := &Transaction{
			id: newId(1),
			account: Account{
				newId(1),
				newDocument("30784553513"),
			},
			operation: newOperation(1),
			amount:    newAmount(100.11),
			eventDate: time.Now().UTC(),
		}

		transaction := &Transaction{
			account: Account{
				newId(1),
				newDocument("30784553513"),
			},
			operation: newOperation(1),
			amount:    newAmount(100.11),
			eventDate: time.Now().UTC(),
		}

		result, _ := transaction.Save(&transactionSaverStub{newId(1)})

		if !reflect.DeepEqual(expected.Id(), result.Id()) {
			t.Errorf("Id() expected %v, received %v", expected.Id(), result.Id())
		}
	})
}

func TestSaveTransactionFailed(t *testing.T) {
	t.Parallel()

	t.Run("Test Save returning failed", func(t *testing.T) {
		expected := errors.New("Failed when inserting transaction")

		transaction := &Transaction{
			account: Account{
				newId(1),
				newDocument("30784553513"),
			},
			operation: newOperation(1),
			amount:    newAmount(100.11),
			eventDate: time.Now().UTC(),
		}

		_, err := transaction.Save(&transactionSaverStub{expected})

		if !reflect.DeepEqual(expected, err) {
			t.Errorf("Error expected %v, received %v", expected, err)
		}
	})
}

type transactionParams struct {
	accountId   uint64
	operationId uint8
	amountValue float64
	repository  AccountFinder
}

func newAmount(v float64) vo.Amount {
	amount, _ := vo.NewAmount(v)
	return *amount
}

func newOperation(o uint8) vo.Operation {
	operation, _ := vo.Operations.Parse(o)
	return *operation
}

type transactionSaverStub struct {
	result interface{}
}

func (s *transactionSaverStub) Save(a *Transaction) (*vo.Id, error) {
	if err, ok := s.result.(error); ok {
		return nil, err
	}

	id, _ := s.result.(*vo.Id)

	return id, nil
}

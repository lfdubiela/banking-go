package entity

import (
	"time"

	"github.com/lfdubiela/banking-go/domain/vo"
)

type Transaction struct {
	id        *vo.Id
	account   Account
	operation vo.Operation
	amount    vo.Amount
	eventDate time.Time
}

func (t Transaction) Id() *vo.Id {
	return t.id
}

func (t Transaction) Account() Account {
	return t.account
}

func (t Transaction) Operation() vo.Operation {
	return t.operation
}

func (t Transaction) Amount() vo.Amount {
	return t.amount
}

func (t Transaction) EventDate() string {
	return t.eventDate.Format("2006-01-02 15:04:05")
}

func (t Transaction) WithId(id *vo.Id) *Transaction {
	return &Transaction{
		id:        id,
		account:   t.account,
		operation: t.operation,
		amount:    t.amount,
		eventDate: t.eventDate,
	}
}

type TransactionSaver interface {
	Save(t *Transaction) (*vo.Id, error)
}

func NewTransction(
	accountId uint64,
	operationId uint8,
	amountValue float64,
	repository AccountFinder) (*Transaction, error) {

	id, err := vo.NewId(accountId)

	if err != nil {
		return nil, err
	}

	account, err := FindAccount(repository, id)

	if err != nil {
		return nil, err
	}

	operation, err := vo.Operations.Parse(operationId)

	if err != nil {
		return nil, err
	}

	amount, err := vo.NewAmount(amountValue)

	if err != nil {
		return nil, err
	}

	return &Transaction{
		account:   *account,
		operation: *operation,
		amount:    *amount,
		eventDate: time.Now().UTC(),
	}, nil
}

func (t *Transaction) Save(r TransactionSaver) (*Transaction, error) {
	id, err := r.Save(t)

	if err != nil {
		return nil, err
	}

	return t.WithId(id), nil
}

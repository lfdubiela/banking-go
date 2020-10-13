package entity

import (
	"github.com/lfdubiela/banking-go/domain/vo"
)

type Account struct {
	id             *vo.Id
	document       vo.Document
	creditLimit    *vo.Amount
	availableLimit *vo.Amount
}

func (a Account) Document() vo.Document {
	return a.document
}

func (a Account) Id() *vo.Id {
	return a.id
}

func (a Account) CreditLimit() *vo.Amount {
	return a.creditLimit
}

func (a Account) AvailableLimit() *vo.Amount {
	return a.availableLimit
}

func (a Account) HasLimit(amount vo.Amount) bool {
	return a.availableLimit.Value() >= amount.Value()
}

func NewAccount(number string, limit float64, available float64) (*Account, error) {
	document, err := vo.NewDocument(number)

	if err != nil {
		return nil, err
	}

	creditLimit, err := vo.NewAmount(limit)

	if err != nil {
		return nil, err
	}

	availableLimit, err := vo.NewAmount(available)

	if err != nil {
		return nil, err
	}

	return &Account{
		document:       *document,
		creditLimit:    creditLimit,
		availableLimit: availableLimit,
	}, nil
}

func (a Account) WithId(id *vo.Id) *Account {
	return &Account{
		id:             id,
		document:       a.document,
		creditLimit:    a.creditLimit,
		availableLimit: a.availableLimit,
	}
}

func (a Account) WithAvailableLimit(limit *vo.Amount) *Account {
	return &Account{
		id:             a.id,
		document:       a.document,
		creditLimit:    a.creditLimit,
		availableLimit: limit,
	}
}

func (a Account) UpdateAvailableLimit(operation vo.Operation, amount vo.Amount) (*Account, error) {
	if operation.Mode() == vo.OperationModes.Debito {
		limit, err := a.availableLimit.Decrement(amount)

		if err != nil {
			return nil, &AccountHasNoSufficientAvailableLimit{err.Error()}
		}

		return a.WithAvailableLimit(limit), nil
	}

	limit, err := a.availableLimit.Increment(amount)

	if err != nil {
		return nil, err
	}

	return a.WithAvailableLimit(limit), nil
}

type AccountFinder interface {
	Find(id *vo.Id) (*Account, error)
}

type AccountSaver interface {
	Save(a *Account) (*vo.Id, error)
}

type AccountHasNoSufficientAvailableLimit struct {
	error string
}

func (e AccountHasNoSufficientAvailableLimit) Error() string {
	return e.error
}

func FindAccount(r AccountFinder, id *vo.Id) (*Account, error) {
	account, err := r.Find(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Save(r AccountSaver) (*Account, error) {
	id, err := r.Save(a)

	if err != nil {
		return nil, err
	}

	return a.WithId(id), nil
}

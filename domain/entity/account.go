package entity

import (
	"github.com/lfdubiela/banking-go/domain/vo"
)

type Account struct {
	id       *vo.Id
	document vo.Document
}

func (a Account) Document() vo.Document {
	return a.document
}

func (a Account) Id() *vo.Id {
	return a.id
}

func NewAccount(number string) (*Account, error) {
	document, err := vo.NewDocument(number)

	if err != nil {
		return nil, err
	}

	return &Account{document: *document}, nil
}

func (a Account) WithId(id *vo.Id) *Account {
	return &Account{
		id:       id,
		document: a.document,
	}
}

type AccountFinder interface {
	Find(id *vo.Id) (*Account, error)
}

type AccountSaver interface {
	Save(a *Account) (*vo.Id, error)
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

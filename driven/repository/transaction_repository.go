package repository

import (
	"database/sql"

	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/domain/vo"
)

type TransactionRepository struct {
	db *sql.DB
}

func (r TransactionRepository) DB() *sql.DB {
	return r.db
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (r TransactionRepository) Save(t *entity.Transaction) (*vo.Id, error) {
	query := `
		UPDATE account SET available_limit = ?
		WHERE account.id = ?
`
	stmt, err := r.DB().Prepare(query)

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(
		t.Account().AvailableLimit().Value(),
		t.Account().Id().Value())

	if err != nil {
		return nil, err
	}

	query = `
		INSERT INTO transaction (account_id, operation_id, amount, event_date) 
		VALUES (?, ?, ?, ?)
`

	stmt, err = r.DB().Prepare(query)

	if err != nil {
		return nil, err
	}

	result, err = stmt.Exec(
		t.Account().Id().Value(),
		t.Operation().Id(),
		t.Amount().Value(),
		t.EventDate())

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	voId, _ := vo.NewId(uint64(id))

	return voId, nil
}

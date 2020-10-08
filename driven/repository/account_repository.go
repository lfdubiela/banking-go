package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/domain/vo"
)

type AccountRepository struct {
	db *sql.DB
}

type AccountAlreadyExists struct {
	error string
}

func (e *AccountAlreadyExists) Error() string {
	return e.error
}

type AccountNotFound struct {
	error string
}

func (e *AccountNotFound) Error() string {
	return e.error
}

func (r AccountRepository) DB() *sql.DB {
	return r.db
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return AccountRepository{db}
}

func (r AccountRepository) Save(a *entity.Account) (*vo.Id, error) {
	stmt, err := r.DB().Prepare("INSERT INTO account (document_number) VALUES (?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(a.Document().Number())

	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return nil, &AccountAlreadyExists{
				fmt.Sprintf("account with document_number(%s) already exists", a.Document().Number()),
			}
		}

		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	voId, err := vo.NewId(uint64(id))

	if err != nil {
		return nil, err
	}

	return voId, nil
}

func (r AccountRepository) Find(id *vo.Id) (*entity.Account, error) {
	q := `SELECT document_number FROM account WHERE id = ?`

	var document string

	err := r.DB().QueryRow(q, id.Value()).Scan(&document)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &AccountNotFound{"Account could not be found!"}
		}
		return nil, err
	}

	account, err := entity.NewAccount(document)

	log.Println(err, account)

	if err != nil {
		return nil, err
	}

	return account.WithId(id), nil
}

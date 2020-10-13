package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lfdubiela/banking-go/driven/repository"

	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/driving/request"
	"github.com/lfdubiela/banking-go/driving/response"
)

type CreateTransaction struct {
	accountFinder entity.AccountFinder
	saver         entity.TransactionSaver
}

func NewCreateTransaction(
	accountFinder entity.AccountFinder,
	saver entity.TransactionSaver) CreateTransaction {
	return CreateTransaction{accountFinder, saver}
}

func (t CreateTransaction) Handler(w http.ResponseWriter, r *http.Request) {
	emitter := response.NewResponseEmitter(w)
	body, _ := ioutil.ReadAll(r.Body)

	payload, err := request.NewCreateTransaction(body)

	if err != nil {
		responseError := response.NewErrorResponse(map[string]string{"request.body": "invalid payload"})
		emitter.BadRequest(responseError)
		return
	}

	defer r.Body.Close()

	if errs := payload.Validate(); errs != nil {
		responseError := response.NewErrorResponse(errs)
		emitter.BadRequest(responseError)
		return
	}

	transaction, err := entity.NewTransction(
		payload.AccountId,
		payload.OperationId,
		payload.Amount,
		t.accountFinder)

	if err != nil {
		if _, ok := err.(*entity.AccountHasNoSufficientAvailableLimit); ok {
			emitter.UnprocessableEntity(
				response.NewErrorResponse(map[string]string{"request.body.account_id": "Account has no sufficient!"}))
		}

		if notFound, ok := err.(*repository.AccountNotFound); ok {
			emitter.UnprocessableEntity(
				response.NewErrorResponse(map[string]string{"request.body.account_id": notFound.Error()}))
			return
		}
		emitter.InternalServerError(nil)
		return
	}

	transaction, err = transaction.Save(t.saver)

	if err != nil {
		log.Print(err)
		emitter.InternalServerError(nil)
		return
	}

	emitter.Created(response.NewTransactionResponse(transaction))
}

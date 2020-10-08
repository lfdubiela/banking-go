package handlers

import (
	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/driving/request"
	"github.com/lfdubiela/banking-go/driving/response"
	"io/ioutil"
	"log"
	"net/http"
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
	payload, _ := ioutil.ReadAll(r.Body)

	request, err := request.NewCreateTransaction(payload)

	if err != nil {
		responseError := response.NewErrorResponse(map[string]string{"request.body": "invalid payload"})
		emitter.BadRequest(responseError)
		return
	}

	defer r.Body.Close()

	if errs := request.Validate(); errs != nil {
		responseError := response.NewErrorResponse(errs)
		emitter.BadRequest(responseError)
		return
	}

	//todo corrigir erro pode retonar erro de 421 aqui
	transaction, _ := entity.NewTransction(
		request.AccountId,
		request.OperationId,
		request.Amount,
		t.accountFinder)

	transaction, err = transaction.Save(t.saver)

	log.Print(transaction, err)

	if err != nil {
		emitter.InternalServerError(nil)
		return
	}

	response := response.NewTransactionResponse(transaction)
	emitter.Created(response)
}

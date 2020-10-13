package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/driven/repository"
	"github.com/lfdubiela/banking-go/driving/request"
	"github.com/lfdubiela/banking-go/driving/response"
)

type CreateAccount struct {
	saver entity.AccountSaver
}

func NewCreateAccount(saver entity.AccountSaver) CreateAccount {
	return CreateAccount{saver}
}

func (c CreateAccount) Handler(w http.ResponseWriter, r *http.Request) {
	emitter := response.NewResponseEmitter(w)
	body, _ := ioutil.ReadAll(r.Body)

	payload, err := request.NewCreateAccount(body)

	if err != nil {
		emitter.BadRequest(response.NewErrorResponse(map[string]string{"request.body": "invalid payload"}))
		return
	}

	defer r.Body.Close()

	if errs := payload.Validate(); errs != nil {
		responseError := response.NewErrorResponse(errs)
		emitter.BadRequest(responseError)
		return
	}

	account, _ := entity.NewAccount(payload.Document, payload.CreditLimit, payload.CreditLimit)
	account, err = account.Save(c.saver)

	if err != nil {
		if errExists, ok := err.(*repository.AccountAlreadyExists); ok {
			emitter.Conflict(response.NewErrorResponse(map[string]string{"request.body": errExists.Error()}))
			return
		}

		log.Println(err)

		emitter.InternalServerError(nil)
		return
	}

	emitter.Created(response.NewAccountResponse(account))
}

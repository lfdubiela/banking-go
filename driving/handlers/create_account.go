package handlers

import (
	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/driven/repository"
	"github.com/lfdubiela/banking-go/driving/request"
	"github.com/lfdubiela/banking-go/driving/response"
	"io/ioutil"
	"net/http"
)

type CreateAccount struct {
	saver entity.AccountSaver
}

func NewCreateAccount(saver entity.AccountSaver) CreateAccount {
	return CreateAccount{saver}
}

func (c CreateAccount) Handler(w http.ResponseWriter, r *http.Request) {
	emitter := response.NewResponseEmitter(w)
	payload, _ := ioutil.ReadAll(r.Body)

	request, err := request.NewCreateAccount(payload)

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

	//todo
	account, _ := entity.NewAccount(request.Document)
	account, err = account.Save(c.saver)

	if err != nil {
		errExists, ok := err.(*repository.AccountAlreadyExists)

		if ok {
			responseError := response.NewErrorResponse(map[string]string{"request.body": errExists.Error()})
			emitter.Conflict(responseError)
			return
		}

		emitter.InternalServerError(nil)
		return
	}

	response := response.NewAccountResponse(account)
	emitter.Created(response)
}

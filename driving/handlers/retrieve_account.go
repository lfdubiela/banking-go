package handlers

import (
	"net/http"
	"strconv"

	"github.com/lfdubiela/banking-go/driven/repository"

	"github.com/gorilla/mux"
	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/domain/vo"
	"github.com/lfdubiela/banking-go/driving/response"
)

type RetrieveAccount struct {
	finder entity.AccountFinder
}

func NewRetrieveAccount(finder entity.AccountFinder) RetrieveAccount {
	return RetrieveAccount{finder}
}

func (c RetrieveAccount) Handler(w http.ResponseWriter, r *http.Request) {
	emitter := response.NewResponseEmitter(w)

	i, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		emitter.BadRequest(response.NewErrorResponse(map[string]string{"request.param.id": "invalid parameter"}))
		return
	}

	id, err := vo.NewId(i)

	if err != nil {
		emitter.BadRequest(response.NewErrorResponse(map[string]string{"request.param.id": "invalid parameter"}))
		return
	}

	account, err := c.finder.Find(id)

	if err != nil {
		if notFound, ok := err.(*repository.AccountNotFound); ok {
			emitter.NotFound(response.NewErrorResponse(map[string]string{"request.param.id": notFound.Error()}))
			return
		}
		emitter.InternalServerError(nil)
		return
	}

	emitter.OK(response.NewAccountResponse(account))
}

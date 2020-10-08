package handlers

import (
	"github.com/gorilla/mux"
	"github.com/lfdubiela/banking-go/domain/entity"
	"github.com/lfdubiela/banking-go/domain/vo"
	"github.com/lfdubiela/banking-go/driving/response"
	"net/http"
	"strconv"
)

type RetrieveAccount struct {
	finder entity.AccountFinder
}

func NewRetrieveAccount(finder entity.AccountFinder) RetrieveAccount {
	return RetrieveAccount{finder}
}

func (c RetrieveAccount) Handler(w http.ResponseWriter, r *http.Request) {
	emitter := response.NewResponseEmitter(w)

	i, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	id, _ := vo.NewId(i)

	account, err := c.finder.Find(*id)
	if err != nil {
		emitter.InternalServerError(nil)
		return
	}

	emitter.OK(response.NewAccountResponse(account))
}

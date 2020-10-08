package response

import "github.com/lfdubiela/banking-go/domain/entity"

type AccountResponse struct {
	Id             uint64 `json:"id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccountResponse(a *entity.Account) *AccountResponse {
	return &AccountResponse{
		Id:             a.Id().Value(),
		DocumentNumber: a.Document().Number(),
	}
}

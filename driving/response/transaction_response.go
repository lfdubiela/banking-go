package response

import (
	"github.com/lfdubiela/banking-go/domain/entity"
)

type TransactionResponse struct {
	Id        uint64        `json:"id"`
	AccountId uint64        `json:"document_number"`
	Operation OperationType `json:"operation"`
	Amount    float64       `json:"amount"`
	EventDate string        `json:"event_date"`
}

type OperationType struct {
	Id   uint8  `json:"id"`
	Mode string `json:"mode"`
}

func NewTransactionResponse(t *entity.Transaction) *TransactionResponse {
	return &TransactionResponse{
		Id:        t.Id().Value(),
		AccountId: t.Account().Id().Value(),
		Operation: OperationType{
			Id:   t.Operation().Id(),
			Mode: t.Operation().Mode(),
		},
		Amount:    t.Amount().Value(),
		EventDate: t.EventDateFormated(),
	}
}

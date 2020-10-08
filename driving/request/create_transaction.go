package request

import (
	"encoding/json"
)

type CreateTransaction struct {
	AccountId   uint64  `json:"account_id" validate:"required,gt=0"`
	OperationId uint8   `json:"operation_type_id" validate:"required,gte=1,lte=4"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
}

func NewCreateTransaction(payload []byte) (*CreateTransaction, error) {
	createTransaction := &CreateTransaction{}

	if err := json.Unmarshal(payload, createTransaction); err != nil {
		return nil, err
	}

	return createTransaction, nil
}

func (c CreateTransaction) Validate() map[string]string {
	return validate.invoke(c)
}

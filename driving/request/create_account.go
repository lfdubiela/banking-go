package request

import (
	"encoding/json"
)

type CreateAccount struct {
	Document string `json:"document_number" validate:"required,document_number"`
}

func NewCreateAccount(payload []byte) (*CreateAccount, error) {
	createAccount := &CreateAccount{}

	if err := json.Unmarshal(payload, createAccount); err != nil {
		return nil, err
	}

	return createAccount, nil
}

func (c CreateAccount) Validate() map[string]string {
	return validate.invoke(c)
}

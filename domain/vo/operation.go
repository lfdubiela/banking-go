package vo

import (
	"errors"
	"fmt"
)

var OperationModes = newOperationModeRegistry()

var Operations = newOperationRegistry()

type OperationMode = string

type Operation struct {
	id          uint8
	description string
	mode        OperationMode
}

func (o Operation) Id() uint8 {
	return o.id
}

func (o Operation) Description() string {
	return o.description
}

func (o Operation) Mode() OperationMode {
	return o.mode
}

type operationRegistry struct {
	CompraAVista    Operation
	CompraParcelada Operation
	Saque           Operation
	Pagamento       Operation

	operations []Operation
}

type operationModeRegistry struct {
	Debito  OperationMode
	Credito OperationMode
}

func newOperationModeRegistry() operationModeRegistry {
	return operationModeRegistry{
		Debito:  "D",
		Credito: "C",
	}
}

func newOperationRegistry() operationRegistry {
	compraAVista := Operation{1, "COMPRA A VISTA", OperationModes.Debito}
	compraParcelada := Operation{2, "COMPRA PARCELADA", OperationModes.Debito}
	saque := Operation{3, "SAQUE", OperationModes.Debito}
	pagamento := Operation{4, "PAGAMENTO", OperationModes.Credito}

	return operationRegistry{
		CompraAVista:    compraAVista,
		CompraParcelada: compraParcelada,
		Saque:           saque,
		Pagamento:       pagamento,

		operations: []Operation{compraAVista, compraParcelada, saque, pagamento},
	}
}

func (o operationRegistry) List() []Operation {
	return o.operations
}

func (o operationRegistry) Parse(id uint8) (*Operation, error) {
	for _, operation := range o.List() {
		if operation.Id() == id {
			return &operation, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Invalid value for Operation(%d)", id))
}

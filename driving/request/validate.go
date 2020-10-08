package request

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"reflect"
	"strings"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate = newValidator()

type RequestValidater interface {
	Validate() map[string]string
}

type payloadValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func newValidator() *payloadValidator {
	en := en.New()
	uni := ut.New(en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validator := &payloadValidator{validate: validate, translator: trans}
	validator.registerTranslator()
	validator.registerTagName()
	validator.registerValidationDocument()

	return validator
}

func (p payloadValidator) invoke(s interface{}) map[string]string {
	err := p.validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		return p.formatErrors(errs)
	}

	return nil
}

func (p payloadValidator) formatErrors(e validator.ValidationErrors) map[string]string {
	var errs = make(map[string]string, 0)

	for _, err := range e {
		field := err.Field()
		err := err.Translate(p.translator)

		errs["body."+field] = err
	}

	return errs
}

func (p *payloadValidator) registerTranslator() {
	p.validate.RegisterTranslation("required", p.translator, func(ut ut.Translator) error {
		return ut.Add("required", "Cannot be empty!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	p.validate.RegisterTranslation("document_number", p.translator, func(ut ut.Translator) error {
		return ut.Add("document_number", "Must be a valid CPF!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("document_number", fe.Field())
		return t
	})
}

func (p *payloadValidator) registerTagName() {
	p.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func (p *payloadValidator) registerValidationDocument() {
	p.validate.RegisterValidation(`document_number`, func(fl validator.FieldLevel) bool {
		value := fl.Field()
		fmt.Println(value.String())
		return cpfcnpj.ValidateCPF(value.String())
	})
}

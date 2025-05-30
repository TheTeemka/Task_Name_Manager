package validator

import (
	"encoding/json"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: map[string]string{},
	}
}

func (v *Validator) AddError(key, message string) {
	_, ok := v.Errors[key]
	if !ok {
		v.Errors[key] = message
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(constraint bool, key, message string) {
	if !constraint {
		v.AddError(key, message)
	}
}

type ErrorWrapper struct {
	Errors map[string]string `json:"errors"`
}

func (v *Validator) String() string {
	w := ErrorWrapper{
		Errors: v.Errors,
	}
	b, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

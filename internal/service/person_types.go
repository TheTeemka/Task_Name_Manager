package service

import "github.com/TheTeemka/TaskNameManager/pkg/validator"

type CreatePersonReq struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (c *CreatePersonReq) Validate(v *validator.Validator) {
	v.CheckWithRules("Name", c.Name, validator.IsValidLength(1, 20))
	v.CheckWithRules("Surname", c.Surname, validator.IsValidLength(1, 20))
}

type UpdatePersonReq struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func (c *UpdatePersonReq) Validate(v *validator.Validator) {
	v.CheckWithRules("Name", c.Name, validator.IsValidLength(0, 20))
	v.CheckWithRules("Surname", c.Surname, validator.IsValidLength(0, 20))
	v.CheckWithRules("Gender", c.Gender, validator.IsValidLength(0, 10))
	v.CheckWithRules("Nationality", c.Nationality, validator.IsValidLength(0, 10))
}

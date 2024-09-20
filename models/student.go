package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

func (s *Student) Validate() error {
	if err := validator.Validate(s); err != nil {
		return err
	}

	return nil
}

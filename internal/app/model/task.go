package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

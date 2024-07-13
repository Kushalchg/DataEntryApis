package global

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var (
	Logger   *log.Logger
	Validate *validator.Validate
)

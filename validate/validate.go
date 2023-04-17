package validate

import (
	"log"

	"github.com/go-playground/validator"
)

func Validate(json interface{}) error {
	validate := validator.New()
	err := validate.Struct(json)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

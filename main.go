package main

import (
	"github.com/go-playground/validator"
	"log"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}

func validateStruct(payload interface{}) error  {
	validation := validator.New()
	err := validation.Struct(payload).(validator.ValidationErrors)
	return err
}

func main() {
	var user = User{}
	err := validateStruct(user)
	if err != nil {
		log.Println(err)
	}

}
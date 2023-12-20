package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/odanaraujo/golang/send-emails/configurations/validation"
	"github.com/odanaraujo/golang/send-emails/internal/domain/campaign"
)

func main() {
	cont := []campaign.Contact{{Email: "test@gmail.com.br"}, {""}}
	camp := campaign.Campaign{
		Name:    "",
		Contact: cont,
	}
	validate := validator.New()

	err := validate.Struct(camp)

	fmt.Println(validation.ValidateError(err).Error())
}

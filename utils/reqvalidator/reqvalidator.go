package reqvalidator

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

func ReadBytes(value []byte, request interface{}) error {
	if err := json.Unmarshal(value, request); err != nil {
		return err
	}

	return validate.Struct(request)
}

func MakeIncorrectParametersCause(err error) string {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return fmt.Sprintf("Incorrect parameters: %s", err.Error())
	}

	msg := "Incorrect parameters: "

	for i, validationError := range validationErrors {
		if i != 0 {
			msg += "; "
		}
		msg += fmt.Sprintf(
			"Field validation for '%s' field on the '%s' tag with value '%s'",
			validationError.Field(),
			validationError.Tag(),
			validationError.Value(),
		)
	}

	return msg
}

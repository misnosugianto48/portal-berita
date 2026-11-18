package handlers

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestBindErrorMessage(t *testing.T) {
	validate := validator.New()

	var req struct {
		Email string `validate:"required,email"`
	}
	err := validate.Struct(req)
	if got := bindErrorMessage(err); got != "Email: failed on 'required'" {
		t.Fatalf("validation error message: got %q", got)
	}

	if got := bindErrorMessage(fmt.Errorf("bad json")); got != "bad json" {
		t.Fatalf("non-validation error message: got %q", got)
	}
}

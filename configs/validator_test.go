package configs

import (
	"testing"
)

type sampleLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func TestValidatorUserFriendlyMessages(t *testing.T) {
	v := NewStructValidator()

	req := sampleLogin{
		Email:    "invalid",
		Password: "123",
	}

	err := v.Validate(req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	expected := "email must be a valid email, password must be at least 8 characters"
	if err.Error() != expected {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}
}

func TestValidatorRequired(t *testing.T) {
	v := NewStructValidator()

	req := sampleLogin{
		Email:    "",
		Password: "",
	}

	err := v.Validate(req)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	expected := "email is required, password is required"
	if err.Error() != expected {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}
}

func TestValidatorSuccess(t *testing.T) {
	v := NewStructValidator()

	req := sampleLogin{
		Email:    "user@example.com",
		Password: "password123",
	}

	if err := v.Validate(req); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

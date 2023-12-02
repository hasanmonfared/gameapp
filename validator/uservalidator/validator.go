package uservalidator

import (
	"fmt"
	"gameapp/dto"
	"gameapp/pkg/phonenumber"
	"gameapp/pkg/richerror"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return
}
func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) error {
	const op = "uservalidator.ValidateRegisterRequest"
	if !phonenumber.IsValid(req.PhoneNumber) {
		return richerror.New(op).
			WithMessage("phone number is not valid").
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
		if !isUnique {
			return dto.RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}
	if len(req.Name) < 3 {
		return dto.RegisterResponse{}, fmt.Errorf("name lengh should be greater")
	}

	if len(req.Password) < 8 {
		return dto.RegisterResponse{}, fmt.Errorf("password lengh should be greater than 8")
	}
}

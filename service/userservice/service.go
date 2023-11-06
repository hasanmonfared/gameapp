package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
	"gameapp/repository/mysql"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}
type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}
type RegisterResponse struct {
	User entity.User
}

func New(repo *mysql.MySQLDB) Service {
	return Service{repo: repo}
}
func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid. ")
	}
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name lengh should be greater")
	}

	u := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	createdUser, err := s.repo.Register(u)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return RegisterResponse{User: createdUser}, nil
}
func IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	panic("s")
}

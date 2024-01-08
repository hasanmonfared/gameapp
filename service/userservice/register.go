package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/param"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	u := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMD5Hash(req.Password),
		Role:        entity.UserRole,
	}

	createdUser, err := s.repo.Register(u)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return param.RegisterResponse{param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}

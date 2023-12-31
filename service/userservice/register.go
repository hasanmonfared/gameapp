package userservice

import (
	"fmt"
	"gameapp/dto"
	"gameapp/entity"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	u := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(u)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return dto.RegisterResponse{dto.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}

package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
	"gameapp/repository/mysql"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}
type Service struct {
	signKey string
	repo    Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User entity.User
}

func New(repo *mysql.MySQLDB, signKey string) Service {
	return Service{
		repo:    repo,
		signKey: signKey,
	}
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

	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password lengh should be greater than 8")
	}
	u := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(u)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)

	}
	return RegisterResponse{User: createdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("user or password isn't correct.")

	}
	if user.Password != GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("password isn't correct")
	}
	token, err := createToken(user.ID, s.signKey)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return LoginResponse{
		AccessToken: token,
	}, nil
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
func createToken(userID uint, signKey string) (string, error) {
	// create a signer for rsa 256
	// set our claims
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(signKey))
	if err != nil {
		return "", nil
	}

	// Creat token string
	return tokenString, nil
}

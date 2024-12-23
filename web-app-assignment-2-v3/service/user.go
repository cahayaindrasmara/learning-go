package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	//mengambil data user berdsarkan email
	dbLogin, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	//validasi password
	if dbLogin.Password != user.Password {
		return nil, errors.New("wrong email password")
	}

	// Mengisi data user yang berhasil login ke pointer user
	user.ID = dbLogin.ID
	user.Fullname = dbLogin.Fullname

	//membuat token JWT jika login berhasil
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", dbLogin.ID),
		ExpiresAt: expirationTime.Unix(),
	}

	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenJWT.SignedString([]byte("secret-key"))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil // TODO: replace this
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	userTaskCategories, err := s.userRepo.GetUserTaskCategory()
	return userTaskCategories, err // TODO: replace this
}

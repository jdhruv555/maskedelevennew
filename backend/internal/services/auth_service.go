package services

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
)

type AuthService struct {
	UserRepo interfaces.UserRepository
}

func NewAuthService(repo interfaces.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) RegisterUser(user *models.User) error {
	if errs := ValidateUserInput(user); errs != nil {
		var sb strings.Builder 
		for field, msg := range errs {
			sb.WriteString(field + ": " + msg + ";")
		}
		return errors.New(strings.TrimSuffix(sb.String(), ": "))
	}

	existingUser, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return  errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	
	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	user.Password = ""
	return user, nil
}

func (s *AuthService) GetUserByID(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.UserRepo.GetUserByID(objID)
}

func (s *AuthService) UpdateUser(id string, updated *models.User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updated.ID = primitive.NilObjectID
	updateData := bson.M{
		"name":     updated.Name,
		"email":    updated.Email,
		"phone":    updated.Phone,
		"address":  updated.Address,
		"updatedAt": time.Now(),
	}
	return s.UserRepo.UpdateUser(objID, updateData)
	
}


func (s *AuthService) DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.UserRepo.DeleteUser(objID)
}

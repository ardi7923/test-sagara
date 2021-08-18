package service

import (
	"errors"
	"log"

	"github.com/ardi7923/test-sagara/entity"
	"github.com/ardi7923/test-sagara/repository"
	"github.com/ardi7923/test-sagara/request"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user request.UserCreateRequest) entity.User
	FindByUsername(username string) entity.User
	IsDuplicateUsername(username string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)

	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) CreateUser(user request.UserCreateRequest) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Fatal("Failed map", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateUsername(username)
	return !errors.Is(res.Error, gorm.ErrRecordNotFound)
}

func (service *authService) FindByUsername(username string) entity.User {
	return service.userRepository.FindByUsername(username)
}

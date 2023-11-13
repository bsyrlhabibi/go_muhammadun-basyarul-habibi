package service

import (
	models "belajar-go-echo/module/entities"
	"belajar-go-echo/module/feature/users/repository"
	"errors"

	"github.com/golang-jwt/jwt"
)

var ErrAuthenticationFailed = errors.New("authentication failed")

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) error
	Authenticate(email, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	secretKey      []byte
}

func NewUserService(userRepository repository.UserRepository, secretKey string) UserService {
	return &userService{
		userRepository: userRepository,
		secretKey:      []byte(secretKey),
	}
}

func (s *userService) CreateUser(user models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) Authenticate(email, password string) (string, error) {
	user, _ := s.userRepository.GetUserByEmail(email)

	if user.Password == password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		tokenString, err := token.SignedString(s.secretKey)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}

	return "", ErrAuthenticationFailed
}

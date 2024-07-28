package service

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	salt = "fasf32d"
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *UserService) CreateUser(user structure.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	fmt.Println(user)
	return s.repository.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

func (s *UserService) CreateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) GetUser(user structure.User) (int, error) {
	password := generatePasswordHash(user.Password)
	checkPassword, userID, err := s.repository.GetUserByEmail(user.Email)
	fmt.Println(checkPassword, userID)
	if err != nil {
		//TODO: error sql request
	}

	if checkPassword == password && userID != 0 {
		return userID, nil
	}

	//TODO: error invalid user
	return 0, nil
}

package service

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	salt = "fasf32d"
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

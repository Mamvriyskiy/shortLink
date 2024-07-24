package service 

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"crypto/rand"
	"math/big"
	"fmt"
)

const (
	letter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	sizeShortLink = 6
)

type LinkService struct {
	repository repository.LinkRepository
}

func NewLinkService(repository repository.LinkRepository) *LinkService {
	return &LinkService{repository: repository}
}

func shortingLink() string {
	str := make([]byte, 6)

	maxInt := big.NewInt(int64(len(letter)))

	for i := 0; i < sizeShortLink; i++ {
		randomNumber, err := rand.Int(rand.Reader, maxInt)
		if err != nil {
			//TODO: log err
			return "", err
		}

		str[i] = letter[int(randomNumber.Int64())]
	}

	return string(str)
}

func (s *LinkService) CreateShortLink(link structure.Link) (string, error) {
	shortLink, err := shortingLink()
	if err != nil {
		//TODO: log err
		return "", err
	}
	//result, err = checkLink(shortLink)

	return "", nil
}

func (s * LinkService) GetLongLink(link structure.Link) (string, error) {
	return "", err
}

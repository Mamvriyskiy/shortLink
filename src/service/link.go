package service 

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"crypto/rand"
	"math/big"
	"net/url"
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

func isValidLink(link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	u, err := url.Parse(link)
	if err != nil || u.Host == "" {
		return false
	}

	return true
}

func shortingLink() (string, error) {
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

	return string(str), nil
}

func (s *LinkService) CheckValidLink(link string) bool {
	if isValidLink(link) {
		return true
	}

	return false
}

func (s *LinkService) CreateShortLink(link structure.Link) (string, error) {
	shortLink, err := shortingLink()
	if err != nil {
		//TODO: log err
		return "", err
	}

	return shortLink, nil
}

func (s * LinkService) CheckDuplicateShortLink(link string) (bool, error) {
	return s.repository.CheckDuplicateShortLink(link)
}

func (s * LinkService) AddLink(link structure.Link, userID int) (int, error) {
	linkID, err := s.repository.AddLink(link, userID)
	return linkID, err
}	

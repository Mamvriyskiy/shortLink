package service

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/repository"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
)

type LinkServices interface {
	CreateShortLink(link structure.Link) (string, error)
	CheckValidLink(link string) bool
	CheckDuplicateShortLink(link string) (bool, error)
	AddLink(link structure.Link, userID int)(int, error)
}

type Services struct {
	LinkServices
}

func NewService(repo *repository.Repository) *Services {
	return &Services{
		LinkServices: NewLinkService(repo),
	}
}


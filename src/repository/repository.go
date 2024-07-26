package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
)

type LinkRepository interface {
	AddLink(link structure.Link, userID int) (int, error)
	CheckDuplicateShortLink(link string) (bool, error)
	GetLongLink(shortLink string) (string, error)
}

type UserRepository interface {
	CreateUser(user structure.User) (int, error)
}

type Repository struct {
	LinkRepository
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository {
		LinkRepository: NewLinkPostgres(db),
		UserRepository: NewUserPostgres(db),
	}
}

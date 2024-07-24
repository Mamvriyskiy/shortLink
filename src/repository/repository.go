package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
)

type LinkRepository interface {
	AddLink(link structure.Link) (int, error)
}

type Repository struct {
	LinkRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository {
		LinkRepository: NewLinkPostgres(db),
	}
}

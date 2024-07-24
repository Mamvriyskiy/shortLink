package repository

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/jmoiron/sqlx"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (l *LinkPostgres) AddLink(link structure.Link) (int, error) {
	_ = link
	return 1, nil
}

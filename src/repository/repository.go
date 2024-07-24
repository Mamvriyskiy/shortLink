package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	//interface
}

func NewRepository(db *sqlx.DB) *repository {
	return &Repository{
		
	}
}
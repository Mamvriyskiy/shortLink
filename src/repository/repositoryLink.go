package repository

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (l *LinkPostgres) GetLongLink(shortLink string) (string, error) {
	var longLink string
	query := fmt.Sprintf("select longlink from link where shortlink = $1")
	err := l.db.Get(&longLink, query, shortLink)
	if err != nil {
		fmt.Println(err)
		return "", err
		//TODO: logger
	}

	return longLink, err
}

func (l *LinkPostgres) AddLink(link structure.Link, userID int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (shortlink, longlink) values ($1, $2) RETURNING linkID", "link")
	row := l.db.QueryRow(query, link.ShortLink, link.LongLink)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		//logger.Log("Error", "Scan", "Error insert into link:", err, userID, link)
		return 0, err
	}

	return 1, nil
}

func (l *LinkPostgres) CheckDuplicateShortLink(link string) (bool, error) {
	var flag int
	query := fmt.Sprintf("select count(*) from link where shortlink = $1")

	err := l.db.Get(&flag, query, link)
	if err != nil {
		logger.Log("Error", "l.db.Get(&flag, query, link)", "Error get duplicatelink:", err)
		return false, err
	}

	if flag == 1 {
		return true, nil
	}

	return false, nil
}

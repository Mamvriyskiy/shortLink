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
		logger.Log("Error", " GetLongLink(shortLink string)", fmt.Sprintf("No search long link for %s", shortLink), err)
		return "", err
	}

	return longLink, err
}

func (l *LinkPostgres) AddLink(link structure.Link, clientID int) (int, error) {
	var linkID int
	query := fmt.Sprintf("INSERT INTO %s (shortlink, longlink) values ($1, $2) RETURNING linkID", "link")
	row := l.db.QueryRow(query, link.ShortLink, link.LongLink)
	if err := row.Scan(&linkID); err != nil {
		logger.Log("Error", "Scan", fmt.Sprintf("Error insert into link: %s, clinetID %s", link, clientID), err)
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO clientlink (linkID, clientID) values ($1, $2)")
	_, err := l.db.Exec(query, linkID, clientID)
	if err != nil {
		logger.Log("Error", "Exec", fmt.Sprintf("Error insert into historydevice: %s, clinetID %s", link, clientID), err)
		return 0, err
	}


	return linkID, nil
}

func (l *LinkPostgres) CheckDuplicateShortLink(link string) (bool, error) {
	var flag int
	query := fmt.Sprintf("select count(*) from link where shortlink = $1")

	err := l.db.Get(&flag, query, link)
	if err != nil {
		logger.Log("Error", "Get", fmt.Sprintf("rror get duplicatelink %s", link), err)
		return false, err
	}

	if flag == 1 {
		return true, nil
	}

	return false, nil
}

package repository

import (
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user structure.User) (int, error) {
	fmt.Println(user)
	var userID int
	query := fmt.Sprintf("INSERT into client (name, password, email) values ($1, $2, $3) RETURNING clientID")
	row := r.db.QueryRow(query, user.Login, user.Password, user.Email)
	if err := row.Scan(&userID); err != nil {
		fmt.Println(err)
		//logger.Log("Error", "Scan", "Error insert into client:", err, ownerID, h)
		return 0, err
	}
	//insert into client (name, password, email) values ('123', '123', '123');

	return userID, nil
}

type GetUser struct {
	ClientID int `db:"clientid"`
	Password string `db:"password"`
}

func (r *UserPostgres) GetUserByEmail(email string) (string, int, error) {
	var user GetUser
	fmt.Println(email)
	query := fmt.Sprintf("select clientid, password from client where email = $1")
	err := r.db.Get(&user, query, email)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
		//TODO: logger
	}

	return user.Password, user.ClientID, nil
}

package repos

import (
	"database/sql"
	"fmt"
	"tulahackTest/models"
)

type AuthRepo struct {
	db        *sql.DB
	tablename string
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		tablename: "auth",
		db: db,
	}
}

func (ar *AuthRepo) Get(username string) (interface{}, error) {
	var user models.User

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE username='%s'",
		ar.tablename,
		username,
	)
	err := ar.db.QueryRow(query).Scan(&user.ID, &user.UserID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ar *AuthRepo) Insert(data interface{}) error {
	user := data.(models.User)
	query := fmt.Sprintf(
		`INSERT INTO %s (authid)
              VALUES ('%s')`,
			  ar.tablename,
			  user.UserID,
	)

	_, err := ar.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepo) Delete(userid string) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE userid='%s'`,
		ar.tablename,
		userid,
	)

	_, err := ar.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepo) Modify(username string, modification interface{}) error {
	user := modification.(models.User)
	query := fmt.Sprintf(
		`UPDATE %s
				SET %s='%s'
				WHERE username='%s'`,
				ar.tablename,
				user.UserID, user.UserID,
				username,
	)

	_, err := ar.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
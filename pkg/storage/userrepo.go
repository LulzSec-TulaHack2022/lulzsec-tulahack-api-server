package storage

import (
	"fmt"
	"tulahackTest/models"
)

func (s *Storage) AddUser(user models.User) error {
	query := fmt.Sprintf(
		`INSERT INTO users (userid)
				VALUES ('%s')`,
				user.UserID,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

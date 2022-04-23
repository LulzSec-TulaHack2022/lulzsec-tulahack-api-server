package storage

import (
	"fmt"
	"github.com/rs/xid"
	"tulahackTest/models"
)

func (s *Storage) AddFlower(flower models.UserFlower) error {
	query := fmt.Sprintf(
		`INSERT INTO user_flowers (catalog_id, flower_id, owner_id, name)
				VALUES ('%v', '%s', '%s', '%s')`,
				flower.CatalogID,
				xid.New().String(),
				flower.OwnerID,
				flower.Name,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteFlower() {

}


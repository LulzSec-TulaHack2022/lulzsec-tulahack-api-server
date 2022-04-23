package storage

import (
	"fmt"
	"github.com/rs/xid"
	"tulahackTest/models"
)

func (s *Storage) AddFlower(flower models.UserFlower) error {
	query := fmt.Sprintf(
		`INSERT INTO user_flowers (catalog_id, flower_id, owner_userid, name)
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

func (s *Storage) DeleteFlower(flowerid string) error {
	query := fmt.Sprintf(
		`DELETE FROM user_flowers
				WHERE flower_id='%s'`,
				flowerid,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) GetAllUserFlowers(userid string) ([]models.UserFlower, error) {
	var flowers []models.UserFlower

	query := fmt.Sprintf(
		`SELECT * FROM user_flowers
                WHERE owner_userid='%s'`,
				userid,
	)

	rows, err := s.db.Query(query)
	if err != nil {
		return []models.UserFlower{}, err
	}

	for rows.Next() {
		var flower models.UserFlower

		err := rows.Scan(
			&flower.ID,
			&flower.CatalogID,
			&flower.FlowerID,
			&flower.OwnerID,
			&flower.Name,
			&flower.Alive,
			&flower.NeedWater,
		)
		if err != nil {
			return []models.UserFlower{}, err
		}

		flowers = append(flowers, flower)
	}

	return flowers, nil
}

func (s *Storage) GetUserFlower(flowerid string) (models.UserFlower, error) {
	var flower models.UserFlower

	query := fmt.Sprintf(
		`SELECT * FROM user_flowers
				WHERE flower_id='%s'`,
				flowerid,
	)

	err := s.db.QueryRow(query).Scan(
		&flower.ID,
		&flower.CatalogID,
		&flower.FlowerID,
		&flower.OwnerID,
		&flower.Name,
		&flower.Alive,
	)

	if err != nil {
		return models.UserFlower{}, err
	}

	return flower, nil
}

func (s *Storage) Dead(flowerid string) error {
	query := fmt.Sprintf(
		`UPDATE user_flowers
				SET alive='false'
				WHERE flower_id='%s'`,
				flowerid,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
package storage

import (
	"fmt"
	"tulahackTest/models"
)

func (s *Storage) GetFlowerInfo(name string) (models.Flower, error) {
	var flower models.Flower

	query := fmt.Sprintf(
		"SELECT * FROM flowers WHERE name='%s'",
		name,
	)
	err := s.db.QueryRow(query).Scan(&flower.ID,
									  &flower.Name,
									  &flower.Description,
									  &flower.Temperature,
									  &flower.Humidity,
									  &flower.Illumination,
	)
	if err != nil {
		return models.Flower{}, err
	}

	return flower, nil
}

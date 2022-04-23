package storage

import (
	"tulahackTest/models"
)

func (s *Storage) Catalog() ([]models.Flower, error) {
	var flowers []models.Flower

	query := "SELECT * FROM catalog"

	rows, err := s.db.Query(query)

	for rows.Next() {
		var flower models.Flower

		err := rows.Scan(
			&flower.ID,
			&flower.Name,
			&flower.Description,
			&flower.Temperature,
			&flower.Humidity,
			&flower.Illumination,
		)

		if err != nil {
			return []models.Flower{}, nil
		}

		flowers = append(flowers, flower)
	}

	if err != nil {
		return []models.Flower{}, err
	}

	return flowers, nil
}

package storage

// * Структура базы данных, методы создания и закрытия подключения

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db       *sql.DB
	config   *Config
}

func NewStorage() *Storage {
	s := &Storage{
		config: NewConfig(),
	}

	err := s.Open()
	if err != nil {
		panic(err)
		return nil
	}

	return s
}

func (s *Storage) Open() error {
	url := fmt.Sprintf(
		"port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		s.config.port,
		s.config.host,
		s.config.user,
		s.config.password,
		s.config.dbname,
		s.config.sslmode,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}


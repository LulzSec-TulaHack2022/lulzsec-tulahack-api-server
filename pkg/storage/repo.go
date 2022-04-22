package storage

// * Интерфейс репозитория базы данных

type Repo interface {
	Get(string)                 (interface{}, error)
	Insert(interface{})         error
	Delete(string)              error
	Modify(string, interface{}) error
}

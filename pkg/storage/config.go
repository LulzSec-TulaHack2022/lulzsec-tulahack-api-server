package storage

// * Конфигурационная структура для инициализации объекта базы данных

type Config struct {
	user     string
	password string
	dbname   string
	sslmode  string
	port     string
	host     string
}

func NewConfig() *Config {
	return &Config{
		user: "gtcroocyndeudh",
		password: "1686ad823168ef73205560775752f4e2fad3f39bd6509019a3b10dcb14cc2b2b",
		dbname: "d5moneq96fl7a5",
		sslmode: "require",
		port: "5432",
		host: "ec2-52-18-116-67.eu-west-1.compute.amazonaws.com",
	}
}

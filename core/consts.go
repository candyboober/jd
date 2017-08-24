package core

import "os"

const (
	MaxBodySize int64 = 1048576
)

var Secret = []byte("secret")

type PostgresData struct {
	Host     string
	Port     string
	User     string
	Password string
}

var Postgres = PostgresData{
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("POSTGRES_PORT"),
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
}

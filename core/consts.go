package core

const (
	MaxBodySize int64 = 1048576

	PostgresLogMode bool = true
	PageSize uint16 = 20
)

var Secret = []byte("secret")

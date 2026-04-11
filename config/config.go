package config

const (
	DBname   = "go_shopping"
	Username = "postgres"
	Password = "12345678"
	Host     = "localhost"
	Port     = "5432"
)

const (
	ServerIp   = "127.0.0.1"
	ServerPort = "8080"
)

const (
	JWTSecret = "my_secret"
)

const (
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDB       = 0
)

const (
	CacheTTLSeconds        = 300
	OrderAutoCancelMinutes = 30
	OrderCancelScanSeconds = 10
)

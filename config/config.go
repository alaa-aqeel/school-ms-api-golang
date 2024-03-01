package config

import (
	"time"
)

// Default Database
var DB_DRIVER = "postgres"
var DB_PORT = 5432
var DB_HOST = "localhost"
var DB_NAME = "school-ms"
var DB_USERNAME = "root"
var DB_PASSWORD = "root"
var DB_SSLMODE = "disable"

// JWT
var JWT_SECRET_KEY = Env("JWT_SECRET_KEY", []byte("107684d9d2b8664121807f3d654"))
var JWT_EXPIRE_TOKEN = Env("JWT_EXPIRE_TOKEN", time.Now().Add(time.Hour*24).Unix())

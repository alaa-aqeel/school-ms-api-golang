package config

import "time"

// JWT
var JWT_SECRET_KEY []byte = []byte("107684d9d2b8664121807f3d654")
var JWT_EXPIRE_TOKEN = time.Now().Add(time.Hour * 24).Unix()

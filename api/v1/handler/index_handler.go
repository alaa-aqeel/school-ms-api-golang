package handler

import (
	"net/http"

	"github.com/alaa-aqeel/school-ms-api-golang/internal/util"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	util.Response(w).Success("Hello world", User{
		Name:  "Alaa Aqeel",
		Email: "alaa.21.iraq@gmail.com",
	})
}

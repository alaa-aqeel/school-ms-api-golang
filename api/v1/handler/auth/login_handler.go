package auth

import (
	"net/http"

	"github.com/alaa-aqeel/school-ms-api-golang/internal/util"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	util.Response(w).Json(util.Map{
		"message": "Hello world",
	})
}

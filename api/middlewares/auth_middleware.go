package middlewares

import (
	"net/http"

	"github.com/alaa-aqeel/school-ms-api-golang/internal/auth"
	"github.com/alaa-aqeel/school-ms-api-golang/internal/util"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var token = auth.GetTokenFromHeader(r)
		var _, err = auth.VerifyToken(token)
		if err != nil {
			util.Response(rw).StatusCode(401).Json(util.Map{
				"status":  "error",
				"message": "Unauthorized",
			})
		}
	})
}

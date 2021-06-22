package middlewares

import (
	"../responses"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}


func AuthJwtVerify(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{
			"status": "failed",
			"message": "Missing authorization token"}

		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)


		jwtToken := strings.Split(header, "Bearer")
		if len(jwtToken) == 1 {
			resp["status"] = "failed"
			resp["message"] = "failed to process the token"
			responses.JSON(w, http.StatusBadRequest, resp)

			return
		}
		header = strings.TrimSpace(jwtToken[1])

		if header == "" {
			responses.JSON(w, http.StatusForbidden, resp)
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil{
			fmt.Println(err)
		}
		if err != nil {
			resp["status"] = "failed"
			resp["message"] = "Invalid token, please login"
			responses.JSON(w, http.StatusForbidden, resp)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "userID", claims["userID"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
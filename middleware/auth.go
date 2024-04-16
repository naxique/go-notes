package middleware

import "net/http"

func (j *JWT) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("x-jwt-token")
		_, err := j.ValidateJWT(tokenString)
		if err != nil {
			respondWithJSON(w, http.StatusForbidden, "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

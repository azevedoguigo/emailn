package endpoint

import (
	"context"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		provider, err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOAK_URL"))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), tokenString)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		token, _ := jwt.Parse(tokenString, nil)
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

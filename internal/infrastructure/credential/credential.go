package credential

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(token string, ctx context.Context) (string, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	provider, err := oidc.NewProvider(ctx, os.Getenv("KEYCLOAK"))
	if err != nil {
		return "", errors.New("error to connect to the provider")
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
	_, err = verifier.Verify(ctx, token)
	if err != nil {
		return "", errors.New("invalid token")
	}

	tokenJwt, _ := jwt.Parse(token, nil)
	claims := tokenJwt.Claims.(jwt.MapClaims)

	return claims["email"].(string), nil
}

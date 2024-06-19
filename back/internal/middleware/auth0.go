package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func NewAuth0() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", os.Getenv("AUTH0_DOMAIN")))
	if err != nil {
		panic(fmt.Errorf("failed to parse issuer URL: %w", err))
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		panic(fmt.Errorf("failed to create a new validator: %w", err))
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		slog.Warn("failed to validate JWT", "err", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

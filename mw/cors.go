package mw

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSMiddleware sets CORS headers manually in all of the responses.
func CORSMiddleware(h http.Handler) http.Handler {
	// List of allowed domains.
	allowedOrigins := []string{
		"https://rest-ui-dev01.digiapi.pizzahut.com",
		"https://rest-ui-stg01.digiapi.pizzahut.com",
		"https://rest-ui-prod.digiapi.pizzahut.com",
	}

	// If the env is set to dev allow localhost.
	// if os.Getenv("ENV") == "dev" {
	allowedOrigins = append(allowedOrigins, "http://localhost:*")
	// }

	// Configure the cors settings for requests.
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            false,
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Continue with the middleware chain.
		c.ServeHTTP(w, r, h.ServeHTTP)
	})
}

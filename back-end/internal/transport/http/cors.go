package http

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsSettings(frontHost string) *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodDelete,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			frontHost,
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		ExposedHeaders: []string{
			"Content-Type",
		},
		Debug: true,
	})
	return c
}

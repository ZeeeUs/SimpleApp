package http

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsSettings(frontHost string) *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
		},
		AllowedOrigins: []string{
			frontHost,
		},
		AllowedHeaders:     []string{},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
	})
	return c
}

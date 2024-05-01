package middleware

import (
	"go_final_project/internal/config"
	"go_final_project/internal/handler/sign"
	"net/http"
)

func Sign(handler http.Handler) http.Handler {
	if len(config.Manager.TODOPass) == 0 {
		return handler
	}
	return sign.Auth(handler)
}

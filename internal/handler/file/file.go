package file

import (
	"go_final_project/internal/config"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	handler := http.FileServer(http.Dir(config.Manager.WebPath))
	handler.ServeHTTP(w, r)
}

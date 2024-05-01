package date

import (
	dateS "go_final_project/internal/service/date"
	"go_final_project/internal/util"
	"net/http"
	"time"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	now, err := time.Parse(util.DateFormat, r.URL.Query().Get("now"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	date := r.URL.Query().Get("date")

	if len(date) == 0 {
		http.Error(w, "Invalid date param", http.StatusBadRequest)
		return
	}
	repeat := r.URL.Query().Get("repeat")

	ans, err := dateS.NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(ans))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

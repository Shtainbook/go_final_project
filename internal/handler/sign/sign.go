package sign

import (
	"bytes"
	"encoding/json"
	"errors"
	signS "go_final_project/internal/service/sign"
	"go_final_project/internal/util"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(util.MarshalError(err))
			return
		}
		if cookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(util.MarshalError(errors.New(signS.UnAuthorized)))
			return
		}
		err = signS.Service.Auth(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(util.MarshalError(err))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func PostPass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var pass signS.Password
	buff := bytes.Buffer{}

	_, err := buff.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	err = json.Unmarshal(buff.Bytes(), &pass)

	token, err := signS.Service.SignIn(pass)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	ansBody, err := json.Marshal(
		struct {
			Token string `json:"token"`
		}{Token: token})

	_, _ = w.Write(ansBody)
}

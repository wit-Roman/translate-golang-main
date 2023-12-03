package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"translate/types"
	"translate/validate"
)

func SessionCreate(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodPost {

		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var session types.ISessionReq
		err := decoder.Decode(&session)
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		match, err := validate.VerifyArgon(session.Hash, session.Login)
		if !match || err != nil {
			fmt.Println("err", err)
			http.Error(w, "wrong session", http.StatusBadRequest)
			return
		}

		origin := r.Header.Get("Origin")

		token := origin + "/?v=" + session.Login

		jsonBytes, err := json.Marshal(token)
		if err != nil {
			http.Error(w, "crash json write", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

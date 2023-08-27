package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"translate/model"
	"translate/types"
	"translate/validate"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)
	if r.Method == http.MethodGet {

		query := r.URL.Query()
		param := validate.Query(query)

		success, result := model.SelectOrders(param)
		if success {
			if jsonBytes, err := json.Marshal(result); err != nil {
				http.Error(w, "crash json read", http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(jsonBytes)
			}
		} else {
			http.Error(w, "crash sql read", http.StatusBadRequest)
		}
	}

	if r.Method == http.MethodPost {
		//TODO: headerAuthFull := r.Header.Get("Authorization")
		//success, headerAuth := clearToken(headerAuthFull)
		//if !success {
		//	http.Error(w, "wrong header", http.StatusUnauthorized)
		//	return
		//}
		//headerReferer := r.Header.Get("Referer")

		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		fmt.Println(r.ContentLength)

		isValid, validOrder := validate.Order(&r.Body)
		if !isValid {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success := false
		result := types.Order{}
		if validOrder.Order_id == 0 { // Новый id 0
			success, result = model.CreateOrder(validOrder)
		} else {
			success, result = model.UpdateOrder(validOrder)
		}

		if success {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "crash json write", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)
		} else {
			http.Error(w, "crash sql write", http.StatusBadRequest)
		}
	}
}

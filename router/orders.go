package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"translate/helper"
	"translate/model"
	"translate/validate"
)

func OrdersData(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodGet {

		query := r.URL.Query()
		param := validate.Query(query)

		success, result := model.SelectOrdersGet(param)
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
		//fmt.Println(r.ContentLength)

		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		isValid, params := validate.OrdersDataReq(&r.Body)
		if !isValid {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success, result := model.SelectOrders(params)
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

func OrdersCreate(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)
	if r.Method == http.MethodPost {
		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		isValid, validOrder := validate.Order(&r.Body)
		if !isValid {
			fmt.Println(isValid, validOrder)
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success, result := model.CreateOrder(validOrder)
		if success {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "crash json write", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)

			headerAuthFull := r.Header.Get("Authorization")
			sessionIsSet, session := helper.ReadToken(headerAuthFull)
			if !sessionIsSet {
				return
			}
			sayHandler("create", session, validOrder)
		} else {
			http.Error(w, "crash sql write", http.StatusBadRequest)
		}
	}
}

func OrdersUpdate(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodPost {

		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		isValid, validOrder := validate.OrderUpd(&r.Body)
		if !isValid {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success, result := model.UpdateOrder(validOrder)
		if success {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "crash json write", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)

			headerAuthFull := r.Header.Get("Authorization")
			sessionIsSet, session := helper.ReadToken(headerAuthFull)
			if !sessionIsSet {
				return
			}
			sayHandler("update", session, result)
		} else {
			http.Error(w, "crash sql write", http.StatusBadRequest)
		}
	}
}

func OrdersDelete(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodGet {

		query := r.URL.Query()

		isValid, Order_id := validate.OrderId(query)
		if !isValid {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success := model.RemoveOrder(Order_id)
		if success {
			jsonBytes, err := json.Marshal("Deleted")
			if err != nil {
				http.Error(w, "crash json write", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)

			headerAuthFull := r.Header.Get("Authorization")
			sessionIsSet, session := helper.ReadToken(headerAuthFull)
			if !sessionIsSet {
				return
			}
			sayHandler("delete", session, struct{ Order_id int64 }{Order_id})
		} else {
			http.Error(w, "crash sql write", http.StatusBadRequest)
		}
	}
}

func OrderGetById(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodGet {

		query := r.URL.Query()
		id, err := strconv.ParseInt(query.Get("id"), 10, 64)
		if err != nil {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success, result := model.GetOrderById(id)
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
}

func OrdersCellEdit(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)

	if r.Method == http.MethodPost {

		r.Body = http.MaxBytesReader(w, r.Body, 16384)

		isValid, validCell := validate.Cell(&r.Body)
		if !isValid {
			fmt.Println("not valid", validCell)
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		success, result := model.EditCell(validCell)
		if success {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "crash json write", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)

			headerAuthFull := r.Header.Get("Authorization")
			sessionIsSet, session := helper.ReadToken(headerAuthFull)
			if !sessionIsSet {
				return
			}
			sayHandler("edit", session, result)
		} else {
			http.Error(w, "crash sql write", http.StatusBadRequest)
		}
	}
}

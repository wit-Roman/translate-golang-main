package validate

import (
	"encoding/json"
	"io"
	"net/url"
	"translate/constants"
	"translate/helper"
	"translate/types"
)

func Query(urlQuery url.Values) types.SelectParam {
	param := make(map[string]string, 0)
	for key, val := range urlQuery {
		for _, orderField := range helper.OrderFields {
			if key == orderField {
				param[key] = url.QueryEscape(val[0])
				break
			}
		}
	}

	return param // Если пустой, то выбор без условия
}

func Order(reqBody *io.ReadCloser) (bool, types.Order) {
	decoder := json.NewDecoder(*reqBody)
	decoder.DisallowUnknownFields()
	var order types.Order
	err := decoder.Decode(&order)
	if err != nil {
		return false, order
	}

	// if order { TODO заполнение\проверки

	order.Text_length = int64(len(order.Text))
	order.Text_translated_length = int64(len(order.Text_translated))
	order.Creating = constants.CurrentDateNow()

	return true, order
}

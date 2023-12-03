package validate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"translate/constants"
	"translate/helper"
	"translate/types"
)

func Query(urlQuery url.Values) types.IOrdersDataParams {
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

func Order(reqBody *io.ReadCloser) (bool, types.IOrderReq) {
	decoder := json.NewDecoder(*reqBody)
	decoder.DisallowUnknownFields()
	var order types.IOrderReq
	err := decoder.Decode(&order)
	if err != nil {
		return false, order
	}

	// if order { TODO заполнение\проверки

	if order.Text_length == 0 {
		order.Text_length = int64(len(order.Text))
	}
	if order.Text_translated_length == 0 {
		order.Text_translated_length = int64(len(order.Text_translated))
	}
	if order.Creating == 0 {
		order.Creating = constants.CurrentDateNow()
	}

	return true, order
}

func OrderUpd(reqBody *io.ReadCloser) (bool, types.IOrder) {
	decoder := json.NewDecoder(*reqBody)
	decoder.DisallowUnknownFields()

	var order types.IOrder
	err := decoder.Decode(&order)
	if err != nil {
		return false, order
	}

	return true, order
}

func OrdersDataReq(reqBody *io.ReadCloser) (bool, types.IOrdersDataReq) {
	decoder := json.NewDecoder(*reqBody)
	decoder.DisallowUnknownFields()

	var ordersDataReq types.IOrdersDataReq
	err := decoder.Decode(&ordersDataReq)
	if err != nil {
		return false, ordersDataReq
	}

	// if ordersDataReq. { TODO заполнение\проверки

	return true, ordersDataReq
}

func OrderId(query url.Values) (bool, int64) {
	Order_id, err := strconv.ParseInt(query.Get("id"), 10, 64)
	if err != nil {
		return false, 0
	}

	return true, Order_id
}

func Cell(reqBody *io.ReadCloser) (bool, types.ICellReq) {
	decoder := json.NewDecoder(*reqBody)
	decoder.DisallowUnknownFields()

	var cell types.ICellReq
	err := decoder.Decode(&cell)
	if err != nil {
		fmt.Println("validate", err)
		return false, cell
	}

	if len(cell.Column) > 32 { // TODO проверка строки для вставки в SQL
		return false, cell
	}

	return true, cell
}

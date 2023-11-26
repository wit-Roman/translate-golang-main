package model

import (
	"fmt"
	"translate/db"
	"translate/helper"
	"translate/types"
)

func SelectOrders(param types.SelectParam) (bool, []types.Order) {
	selectQueryWhere, args := helper.CreateSelectQueryWhere(param)

	i := 0
	result := make([]types.Order, 0)

	selectQuery := "SELECT * FROM `orders`" + selectQueryWhere + "; "
	fmt.Println(selectQuery)
	selectedRows, err := db.Connection.Query(selectQuery, args...)
	if err != nil {
		return false, result
	}
	defer selectedRows.Close()

	//columns, err := selectedRows.Columns()
	for selectedRows.Next() {
		var row types.Order
		err := selectedRows.Scan(&i, &row.Order_id, &row.Order_name, &row.Order_status, &row.Order_is_paid, &row.Is_safe_transaction, &row.Status, &row.Performer, &row.Viewer_id, &row.Group_id, &row.First_name, &row.Last_name, &row.Photo_100, &row.Viewer_type, &row.Rights, &row.Date_time, &row.Date_time_taken, &row.Date_time_deadline, &row.Text, &row.Text_length, &row.Text_type, &row.Text_translated, &row.Text_translated_length, &row.Text_translated_readiness, &row.Text_translated_demo, &row.Description, &row.Timezone, &row.Creating)
		if err != nil {
			return false, result
		} else {
			result = append(result, row)
		}
	}

	return true, result
}

func CreateOrder(order types.Order) (bool, types.Order) {
	var result = types.Order{}
	selectQueryColumns, selectQueryValues := helper.GetSelectQueryInsert()
	selectQuery := "INSERT INTO `orders` (" + selectQueryColumns + ") VALUES (" + selectQueryValues + "); "

	lastOrderId := int64(1)
	hasLast, lastOrder := GetLastOrder()
	if hasLast && (lastOrder != types.Order{}) {
		lastOrderId = lastOrder.Order_id + 1
	}

	insert, err := db.Connection.Exec(selectQuery, lastOrderId, &order.Order_name, &order.Order_status, &order.Order_is_paid, &order.Is_safe_transaction, &order.Status, &order.Performer, &order.Viewer_id, &order.Group_id, &order.First_name, &order.Last_name, &order.Photo_100, &order.Viewer_type, &order.Rights, &order.Date_time, &order.Date_time_taken, &order.Date_time_deadline, &order.Text, &order.Text_length, &order.Text_type, &order.Text_translated, &order.Text_translated_length, &order.Text_translated_readiness, &order.Text_translated_demo, &order.Description, &order.Timezone, &order.Creating)
	if err != nil {
		return false, result
	}
	fmt.Println("insert", insert)

	return true, result
}

func UpdateOrder(order types.Order) (bool, types.Order) {
	var result = types.Order{}
	selectQueryUpdate := helper.GetSelectQueryUpdate()
	selectQuery := "UPDATE `orders` SET " + selectQueryUpdate + " WHERE order_id=?; "

	fmt.Println(selectQuery)
	update, err := db.Connection.Exec(selectQuery, &order.Order_name, &order.Order_status, &order.Order_is_paid, &order.Is_safe_transaction, &order.Status, &order.Performer, &order.Viewer_id, &order.Group_id, &order.First_name, &order.Last_name, &order.Photo_100, &order.Viewer_type, &order.Rights, &order.Date_time, &order.Date_time_taken, &order.Date_time_deadline, &order.Text, &order.Text_length, &order.Text_type, &order.Text_translated, &order.Text_translated_length, &order.Text_translated_readiness, &order.Text_translated_demo, &order.Description, &order.Timezone, &order.Creating, &order.Order_id)
	if err != nil {
		fmt.Println(err)
		return false, result
	}
	fmt.Println("update", update)

	return true, order
}

func GetLastOrder() (bool, types.Order) {
	var result = types.Order{}
	selectQuery := "SELECT * FROM `orders` WHERE order_id = ( SELECT MAX(order_id) FROM orders ) ; "

	err := db.Connection.QueryRow(selectQuery).Scan(&result)
	if err != nil {
		return false, result
	}

	return true, result
}

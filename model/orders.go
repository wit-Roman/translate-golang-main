package model

import (
	"fmt"
	"translate/db"
	"translate/helper"
	"translate/types"
)

func SelectOrders(param types.IOrdersDataReq) (bool, types.IOrdersDataResp) {
	result := make([]types.IOrder, 0)

	selectQuery := helper.BuildSql(param, "`orders`", false)
	selectedRows, err := db.Connection.Query(selectQuery)
	if err != nil {
		return false, types.IOrdersDataResp{result, 0}
	}
	defer selectedRows.Close()

	for selectedRows.Next() {
		var row types.IOrder
		err := selectedRows.Scan(&row.Order_id, &row.Order_name, &row.Order_status, &row.Order_is_paid, &row.Is_safe_transaction, &row.Status, &row.Performer, &row.Viewer_id, &row.Group_id, &row.First_name, &row.Last_name, &row.Photo_100, &row.Viewer_type, &row.Rights, &row.Date_time, &row.Date_time_taken, &row.Date_time_deadline, &row.Text, &row.Text_length, &row.Text_type, &row.Text_translated, &row.Text_translated_length, &row.Text_translated_readiness, &row.Text_translated_demo, &row.Description, &row.Timezone, &row.Creating)
		if err != nil {
			return false, types.IOrdersDataResp{result, 0}
		} else {
			result = append(result, row)
		}
	}

	RowCount := GetTableLength(param)

	return true, types.IOrdersDataResp{result, RowCount}
}

// TODO: приспособить для простых запросов с перечислением параметров
func SelectOrdersGet(param types.IOrdersDataParams) (bool, types.IOrdersDataResp) {
	selectQueryWhere, args := helper.CreateSelectQueryWhere(param)

	result := make([]types.IOrder, 0)

	selectQuery := "SELECT * FROM `orders`" + selectQueryWhere + "; "
	//fmt.Println(selectQuery)
	selectedRows, err := db.Connection.Query(selectQuery, args)
	if err != nil {
		return false, types.IOrdersDataResp{result, 0}
	}
	defer selectedRows.Close()

	//columns, err := selectedRows.Columns()
	for selectedRows.Next() {
		var row types.IOrder
		err := selectedRows.Scan(&row.Order_id, &row.Order_name, &row.Order_status, &row.Order_is_paid, &row.Is_safe_transaction, &row.Status, &row.Performer, &row.Viewer_id, &row.Group_id, &row.First_name, &row.Last_name, &row.Photo_100, &row.Viewer_type, &row.Rights, &row.Date_time, &row.Date_time_taken, &row.Date_time_deadline, &row.Text, &row.Text_length, &row.Text_type, &row.Text_translated, &row.Text_translated_length, &row.Text_translated_readiness, &row.Text_translated_demo, &row.Description, &row.Timezone, &row.Creating)
		if err != nil {
			return false, types.IOrdersDataResp{result, 0}
		} else {
			result = append(result, row)
		}
	}

	return true, types.IOrdersDataResp{result, 0}
}

func CreateOrder(order types.IOrderReq) (bool, types.ICreateOrderRes) {
	selectQuery := "INSERT INTO `orders` (`order_name`, `order_status`, `order_is_paid`, `is_safe_transaction`, `status`, `performer`, `viewer_id`, `group_id`, `first_name`, `last_name`, `photo_100`, `viewer_type`, `rights`, `date_time`, `date_time_taken`, `date_time_deadline`, `text`, `text_length`, `text_type`, `text_translated`, `text_translated_length`, `text_translated_readiness`, `text_translated_demo`, `description`, `timezone`, `creating` ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?); "

	insert, err := db.Connection.Exec(selectQuery, &order.Order_name, &order.Order_status, &order.Order_is_paid, &order.Is_safe_transaction, &order.Status, &order.Performer, &order.Viewer_id, &order.Group_id, &order.First_name, &order.Last_name, &order.Photo_100, &order.Viewer_type, &order.Rights, &order.Date_time, &order.Date_time_taken, &order.Date_time_deadline, &order.Text, &order.Text_length, &order.Text_type, &order.Text_translated, &order.Text_translated_length, &order.Text_translated_readiness, &order.Text_translated_demo, &order.Description, &order.Timezone, &order.Creating)
	if err != nil {
		return false, types.ICreateOrderRes{0}
	}

	lastIndex, err := insert.LastInsertId()

	return true, types.ICreateOrderRes{lastIndex}
}

func UpdateOrder(order types.IOrder) (bool, types.IOrder) {
	var result = types.IOrder{}
	selectQueryUpdate := helper.GetSelectQueryUpdate()
	selectQuery := "UPDATE `orders` SET " + selectQueryUpdate + " WHERE i=?; "

	fmt.Println(selectQuery)
	///
	update, err := db.Connection.Exec(selectQuery, &order.Order_name, &order.Order_status, &order.Order_is_paid, &order.Is_safe_transaction, &order.Status, &order.Performer, &order.Viewer_id, &order.Group_id, &order.First_name, &order.Last_name, &order.Photo_100, &order.Viewer_type, &order.Rights, &order.Date_time, &order.Date_time_taken, &order.Date_time_deadline, &order.Text, &order.Text_length, &order.Text_type, &order.Text_translated, &order.Text_translated_length, &order.Text_translated_readiness, &order.Text_translated_demo, &order.Description, &order.Timezone, &order.Creating, &order.Order_id)
	if err != nil {
		fmt.Println(err)
		return false, result
	}
	fmt.Println("update", update)

	return true, order
}

func GetLastOrder() (bool, types.IOrder) {
	var result = types.IOrder{}
	selectQuery := "SELECT * FROM `orders` WHERE i = ( SELECT MAX(i) FROM orders ) ; "

	err := db.Connection.QueryRow(selectQuery).Scan(&result)
	if err != nil {
		return false, result
	}

	return true, result
}

func GetOrderById(id int64) (bool, types.IOrder) {
	var result = types.IOrder{}
	selectQuery := "SELECT * FROM `orders` WHERE i = ?; "

	err := db.Connection.QueryRow(selectQuery, id).Scan(&result)
	if err != nil {
		return false, result
	}

	return true, result
}

func RemoveOrder(order_id int64) bool {
	selectQuery := "DELETE FROM `orders` WHERE i=?; "

	_, err := db.Connection.Exec(selectQuery, order_id)
	if err != nil {
		return false
	}

	return true
}

func GetTableLength(param types.IOrdersDataReq) int64 {
	count := int64(0)
	selectQuery := helper.BuildSql(param, "`orders`", true)

	err := db.Connection.QueryRow(selectQuery).Scan(&count)
	if err != nil {
		return count
	}

	return count
}

func EditCell(cell types.ICellReq) (bool, types.ICellReq) {
	selectQuery := "UPDATE `orders` SET " + cell.Column + " = ? WHERE i=? ; "
	// TODO: проверить с разными типами cell.Value

	_, err := db.Connection.Query(selectQuery, cell.Value, cell.Order_id)
	if err != nil {
		fmt.Println(err)
		return false, types.ICellReq{}
	}

	return true, cell
}

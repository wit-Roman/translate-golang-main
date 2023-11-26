package helper

import (
	"strconv"
	"translate/types"
)

var OrderFields = []string{
	"order_id",
	"order_name",
	"order_status",
	"order_is_paid",
	"is_safe_transaction",
	"status",
	"performer",
	"viewer_id",
	"group_id",
	"first_name",
	"last_name",
	"photo_100",
	"viewer_type",
	"rights",
	"date_time",
	"date_time_taken",
	"date_time_deadline",
	"text",
	"text_length",
	"text_type",
	"text_translated",
	"text_translated_length",
	"text_translated_readiness",
	"text_translated_demo",
	"description",
	"timezone",
	"creating",
}

func CreateSelectQueryWhere(param types.SelectParam) (string, []interface{}) {
	var selectQueryWhere = ""
	if len(param) > 0 {
		selectQueryWhere = " WHERE "
	}
	var args []interface{}
	for key, val := range param {
		_, err := strconv.Atoi(val)
		isString := err != nil

		if len(selectQueryWhere) > 7 {
			if isString {
				selectQueryWhere += " OR "
			} else {
				selectQueryWhere += " AND "
			}
		}
		if isString {
			selectQueryWhere += key + " LIKE ?"
			args = append(args, "%"+val+"%")
		} else {
			selectQueryWhere += key + "=?"
			args = append(args, val)
		}
	}

	return selectQueryWhere, args
}

func GetSelectQueryInsert() (string, string) {
	var selectQueryColumns = ""
	var selectQueryValues = ""
	for _, key := range OrderFields {
		selectQueryColumns += key + ","
		selectQueryValues += "?,"
	}
	selectQueryColumns = selectQueryColumns[:len(selectQueryColumns)-1]
	selectQueryValues = selectQueryValues[:len(selectQueryValues)-1]

	return selectQueryColumns, selectQueryValues
}

func GetSelectQueryUpdate() string {
	var selectQueryUpdate = ""
	for _, key := range OrderFields {
		if key != "order_id" {
			selectQueryUpdate += key + "=?,"
		}
	}
	selectQueryUpdate = selectQueryUpdate[:len(selectQueryUpdate)-1]

	return selectQueryUpdate
}

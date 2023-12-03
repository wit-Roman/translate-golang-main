package helper

import (
	"net/url"
	"strconv"
	"strings"
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

func CreateSelectQueryWhere(params types.IOrdersDataParams) (string, []string) {
	var selectQueryWhere = ""
	if len(params) > 0 {
		selectQueryWhere = " WHERE "
	}
	var args []string
	for key, val := range params {
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

func BuildSql(params types.IOrdersDataReq, tableName string, isCounting bool) string {
	if isCounting {
		return "SELECT COUNT(*) FROM " + tableName + whereSql(params) + "; "
	}

	return "SELECT * FROM " + tableName + whereSql(params) + orderBySql(params) + limitSql(params) + "; "
}

func whereSql(params types.IOrdersDataReq) string {
	var whereParts []string

	//if params.FilterModel != nil {
	for key, item := range params.FilterModel {
		mapper := textFilterMapper

		if item.FilterType == "text" {
			mapper = textFilterMapper
		}
		if item.FilterType == "number" {
			mapper = numberFilterMapper
		}
		if item.FilterType == "date" {
			mapper = dateFilterMapper
		}

		part := createFilterSql(mapper, key, item)
		whereParts = append(whereParts, part)
	}

	if len(whereParts) > 0 {
		return " WHERE " + strings.Join(whereParts, " AND ")
	}

	return ""
}

func createFilterSql(mapper func(key string, item types.IFilterModelEntity) string, key string, item types.IFilterModel) string {
	if item.Operator != "" {
		var conditions []string
		for _, condition := range item.Conditions {
			conditions = append(conditions, mapper(key, condition))
		}

		return "(" + strings.Join(conditions, " "+item.Operator+" ") + ")"
	}

	return mapper(key, types.IFilterModelEntity{item.Type, item.Filter, item.FilterTo, item.FilterType, item.DateFrom, item.DateTo})
}

func textFilterMapper(key string, item types.IFilterModelEntity) string {
	switch item.Type {
	case "equals":
		return key + " = '" + item.Filter + "'"
	case "notEqual":
		return key + " != '" + item.Filter + "'"
	case "contains":
		return key + " LIKE '%" + item.Filter + "%'"
	case "notContains":
		return key + " NOT LIKE '%" + item.Filter + "%'"
	case "startsWith":
		return key + " LIKE '" + item.Filter + "%'"
	case "endsWith":
		return key + " LIKE '%" + item.Filter + "'"
	case "blank":
		return key + " IS NULL or " + key + " = ''"
	case "notBlank":
		return key + " IS NOT NULL and " + key + " != ''"
	default:
		return ""
	}
}

func numberFilterMapper(key string, item types.IFilterModelEntity) string {
	switch item.Type {
	case "equals":
		return key + " = " + item.Filter
	case "notEqual":
		return key + " != " + item.Filter
	case "greaterThan":
		return key + " > " + item.Filter
	case "greaterThanOrEqual":
		return key + " >= " + item.Filter
	case "lessThan":
		return key + " < " + item.Filter
	case "lessThanOrEqual":
		return key + " <= " + item.Filter
	case "inRange":
		return "(" + key + " >= " + item.Filter + " and " + key + " <= " + item.FilterTo + ")"
	case "blank":
		return key + " IS NULL"
	case "notBlank":
		return key + " IS NOT NULL"
	default:
		return ""
	}
}

func dateFilterMapper(key string, item types.IFilterModelEntity) string {
	switch item.Type {
	case "equals":
		return "DATE(" + key + ") = '" + item.DateFrom + "'"
	case "notEqual":
		return "DATE(" + key + ") != '" + item.DateFrom + "'"
	case "greaterThan":
		return "DATE(" + key + ") > '" + item.DateFrom + "'"
	case "greaterThanOrEqual":
		return "DATE(" + key + ") >= '" + item.DateFrom + "'"
	case "lessThan":
		return "DATE(" + key + ") < '" + item.DateFrom + "'"
	case "lessThanOrEqual":
		return "DATE(" + key + ") <= '" + item.DateFrom + "'"
	case "inRange":
		return "(DATE(" + key + ") >= '" + item.DateFrom + "' and DATE(" + key + ") <= '" + item.DateTo + "')"
	case "blank":
		return "DATE(" + key + ") IS NULL"
	case "notBlank":
		return "DATE(" + key + ") IS NOT NULL"
	default:
		return ""
	}
}

func orderBySql(params types.IOrdersDataReq) string {
	if len(params.SortModel) == 0 {
		return ""
	}
	var sorts []string
	for _, sort := range params.SortModel {
		sorts = append(sorts, sort.ColId+" "+strings.ToUpper(sort.Sort))
	}

	return " ORDER BY " + strings.Join(sorts, ", ")
}

func limitSql(params types.IOrdersDataReq) string {
	if params.EndRow == -1 && params.StartRow == -1 {
		return ""
	}

	start := int64(0)
	offset := " OFFSET 0"
	if params.StartRow != 0 && params.StartRow != -1 {
		offset = " OFFSET " + strconv.Itoa(int(params.StartRow))
		start = params.StartRow
	}

	limit := ""
	if params.EndRow != 0 && params.EndRow != -1 {
		blockSize := params.EndRow - start
		limit = " LIMIT " + strconv.Itoa(int(blockSize))
	}

	return limit + offset
}

func ReadToken(token string) (bool, types.ITokenData) {
	myUrl, err := url.Parse(token)
	if err != nil {
		return false, types.ITokenData{}
	}
	query, err := url.ParseQuery(myUrl.RawQuery)
	if err != nil {
		return false, types.ITokenData{}
	}

	viewer := query.Get("v")
	if viewer == "" {
		return false, types.ITokenData{}
	}

	return true, types.ITokenData{viewer}
}

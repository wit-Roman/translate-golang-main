package types

type IOrder struct {
	Order_id                  int64
	Order_name                string
	Order_status              int64
	Order_is_paid             bool
	Is_safe_transaction       bool
	Status                    int64
	Performer                 string
	Viewer_id                 int64
	Group_id                  int64
	First_name                string
	Last_name                 string
	Photo_100                 string
	Viewer_type               string
	Rights                    int64
	Date_time                 string
	Date_time_taken           string
	Date_time_deadline        string
	Text                      string
	Text_length               int64
	Text_type                 int64
	Text_translated           string
	Text_translated_length    int64
	Text_translated_readiness string
	Text_translated_demo      string
	Description               string
	Timezone                  int64
	Creating                  int64
}
type IOrderReq struct {
	Order_name                string
	Order_status              int64
	Order_is_paid             bool
	Is_safe_transaction       bool
	Status                    int64
	Performer                 string
	Viewer_id                 int64
	Group_id                  int64
	First_name                string
	Last_name                 string
	Photo_100                 string
	Viewer_type               string
	Rights                    int64
	Date_time                 string
	Date_time_taken           string
	Date_time_deadline        string
	Text                      string
	Text_length               int64
	Text_type                 int64
	Text_translated           string
	Text_translated_length    int64
	Text_translated_readiness string
	Text_translated_demo      string
	Description               string
	Timezone                  int64
	Creating                  int64
}

type IOrdersDataParams map[string]string

type IOrdersDataReq struct {
	StartRow    int64
	EndRow      int64
	FilterModel map[string]IFilterModel
	SortModel   []ISortModel
}

type IFilterModel struct {
	Type       string
	Filter     string
	FilterTo   string
	FilterType string
	DateFrom   string
	DateTo     string

	Conditions []IFilterModelEntity
	Operator   string
}

type IFilterModelEntity struct {
	Type       string
	Filter     string
	FilterTo   string
	FilterType string
	DateFrom   string
	DateTo     string
}

type ISortModel struct {
	ColId string
	Sort  string
}

type IOrdersDataResp struct {
	Rows     []IOrder
	RowCount int64
}

type ICreateOrderRes struct {
	LastIndex int64
}

type ICellReq struct {
	Order_id int64
	Column   string
	Value    string
}

type ISayHandlerMessage struct {
	MesType string
	Viewer  string
	Data    any
}

type ITokenData struct {
	Viewer string
}

type ISessionReq struct {
	Login string
	Hash  string
}

package types

type Order struct {
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

type SelectParam map[string]string

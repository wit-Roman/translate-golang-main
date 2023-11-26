package constants

import "time"

//127.0.0.1:3306
//94.228.115.58:3306
const (
	Port            = ":3001"
	MysqlRequisites = "rwit:rerevfhbzdb@tcp(94.228.115.58:3306)/translate?multiStatements=true"
	//mysqlOpen  = "rwit:rerevfhbzdb@tcp(127.0.0.1:3306)/vkapp?multiStatements=true"
	secretKey  = ""
	serviceKey = ""
	app_id     = 0
	CIPHER_KEY = ""
)

func CurrentDateNow() int64 {
	return time.Now().Unix()
}

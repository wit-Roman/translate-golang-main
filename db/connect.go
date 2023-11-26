package db

import (
	"database/sql"
	"translate/constants"

	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable to hold db connection
var Connection *sql.DB

func Connect() {
	db, err := sql.Open("mysql", constants.MysqlRequisites)
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
	//https://habr.com/ru/company/ispring/blog/560032/
	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute)

	Connection = db
}

func CreateTables() {
	SQLquery := `CREATE TABLE IF NOT EXISTS orders ( 
		i INT PRIMARY KEY AUTO_INCREMENT,

		order_id INT NOT NULL, 
		order_name VARCHAR(80) NOT NULL,
		order_status MEDIUMINT NOT NULL,
		order_is_paid BOOLEAN,
		is_safe_transaction BOOLEAN,

		status MEDIUMINT NOT NULL,
		performer VARCHAR(80),

		viewer_id INT, group_id INT, first_name VARCHAR(80), last_name VARCHAR(80), photo_100 VARCHAR(200), viewer_type VARCHAR(16), rights SMALLINT,

		date_time DATETIME, date_time_taken DATETIME, date_time_deadline DATETIME,

		text MEDIUMTEXT NOT NULL, text_length MEDIUMINT, text_type MEDIUMINT,
		text_translated MEDIUMTEXT NOT NULL, text_translated_length MEDIUMINT, text_translated_readiness VARCHAR(80), text_translated_demo TEXT,

		description TEXT NOT NULL,
		timezone SMALLINT, creating INT 
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		
	CREATE TABLE IF NOT EXISTS order_statuses (
		i INT PRIMARY KEY AUTO_INCREMENT,
		status_id MEDIUMINT,
		name VARCHAR(32),
		value VARCHAR(32)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := Connection.Exec(SQLquery)
	if err != nil {
		log.Fatal(err)
	}
}

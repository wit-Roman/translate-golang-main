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
 
		order_name VARCHAR(80) NOT NULL,
		order_status MEDIUMINT NOT NULL,
		order_is_paid BOOLEAN DEFAULT FALSE,
		is_safe_transaction BOOLEAN DEFAULT FALSE,

		status MEDIUMINT NOT NULL,
		performer VARCHAR(80) NOT NULL,

		viewer_id INT NOT NULL, group_id INT NOT NULL, first_name VARCHAR(80) NOT NULL, last_name VARCHAR(80) NOT NULL, photo_100 VARCHAR(200) NOT NULL, viewer_type VARCHAR(16) NOT NULL, rights SMALLINT NOT NULL,

		date_time DATETIME NOT NULL, date_time_taken DATETIME NOT NULL, date_time_deadline DATETIME NOT NULL,

		text MEDIUMTEXT NOT NULL, text_length MEDIUMINT NOT NULL, text_type MEDIUMINT NOT NULL,
		text_translated MEDIUMTEXT NOT NULL, text_translated_length MEDIUMINT NOT NULL, text_translated_readiness VARCHAR(80) NOT NULL, text_translated_demo TEXT NOT NULL,

		description TEXT NOT NULL,
		timezone SMALLINT NOT NULL, creating BIGINT NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		
	CREATE TABLE IF NOT EXISTS order_statuses (
		i INT PRIMARY KEY AUTO_INCREMENT,
		status_id MEDIUMINT NOT NULL,
		name VARCHAR(32) NOT NULL,
		value VARCHAR(32) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := Connection.Exec(SQLquery)
	if err != nil {
		log.Fatal(err)
	}
}

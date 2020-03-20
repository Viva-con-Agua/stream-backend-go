package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

var DB = new(sql.DB)
var err = sql.ErrConnDone

func ConnectDatabase() {
	var (
		host     = Config.DB.Host
		port     = Config.DB.Port
		user     = Config.DB.User
		password = Config.DB.Password
		name     = Config.DB.Name
	)
	mysqlInfo := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + name
	DB, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	fmt.Println("Database successfully connected!")
	//	query, err := ioutil.ReadFile("drops-database.sql")
	//	if err != nil {
	//		log.Fatal("read sql file failed", err)
	//	}
	//	if _, err := DB.Exec(string(query)); err != nil {
	//		log.Fatal("database init failed", err)
	//	}

}

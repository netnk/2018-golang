package main

import (
	"database/sql"
	"fmt"
	"log"
	

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	host string
	sqlstr string
	team01 string
	team02 string
	team03 string
	team04 string
	
)

func main(){
	dbcon("192.168.1.2", "select team01,team02,team03,team04 from team")
}



func dbcon(host string, sqlstr string) {
	condb, err := sql.Open("mssql", "server=" + host + ";database=db;User ID=****;Password=****;encrypt=disable;")
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}

	rows, err := condb.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next(){
		err := rows.Scan(&team01, &team02, &team03, &team04)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(team01, team02, team03, team04)	

	}
	defer condb.Close()

	
}





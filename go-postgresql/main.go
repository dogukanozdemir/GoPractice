package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main(){

	connStr := "postgres://juomnbnofveahe:51df9706d98feab9b4f75e825c5deb916a546ce1584146c2035b37718632e99f@ec2-3-248-121-12.eu-west-1.compute.amazonaws.com:5432/d2u256bd7288o7"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully connected!")
	defer db.Close()

}
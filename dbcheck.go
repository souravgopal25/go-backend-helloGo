package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type PERSON struct {
	PersonID  int    "json:id"
	LastName  string "json:LastName"
	FirstName string "json:FirstName"
	Address   string "json:Address"
	City      string "json:City"
}

var personSlice []PERSON

func main() {
	fmt.Println("Go MySQL ")
	db, err := sql.Open("mysql", "Sourav:rootpass11@tcp(127.0.0.1:3306)/testCLI")
	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Successfully Connected to MySQL database")
	//INSERTING DATA INTO TABLE;
	insert, err := db.Query("INSERT INTO Person VALUES ('4','Hardy','Thomas','KORMANGLA','BANGLORE')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	//GET DATA FROM TABLE
	result, err := db.Query("SELECT * FROM Person")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println()
	for result.Next() {
		var obj PERSON

		result.Scan(&obj.PersonID, &obj.LastName, &obj.FirstName, &obj.Address, &obj.City)
		personSlice = append(personSlice, obj)
		fmt.Println(obj)
	}
	fmt.Println(personSlice[0])
	//print(insert)
	defer db.Close()
}

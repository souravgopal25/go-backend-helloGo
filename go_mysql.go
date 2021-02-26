package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	_ "strconv"
)

type Student struct {
	Sno    int    "json.sNo"
	RegNo  string "json:regNo"
	Name   string "json:name"
	Branch string "json:branch"
}

func main() {

	fmt.Println("Connection Success")
	//db.Query("INSERT INTO tStudent (RegNo,Name,Branch) VALUES('18','Sourav','cse')")
	r := mux.NewRouter()

	r.HandleFunc("/showAll", showAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Hello Sourav")

}

func showAll(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", "Sourav:rootpass11@tcp(127.0.0.1:3306)/studentDB")

	if err != nil {
		panic(err.Error())
	}
	//writer.Header().Set("Content-Type","applciation/json")
	result, err := db.Query("SELECT * FROM tStudent")
	if err != nil {
		panic(err.Error())
	}
	var slice []Student
	for result.Next() {
		var studentObject Student
		result.Scan(&studentObject.Sno, &studentObject.RegNo, &studentObject.Name, &studentObject.Branch)
		slice = append(slice, studentObject)
	}
	json.NewEncoder(writer).Encode(&slice)
}

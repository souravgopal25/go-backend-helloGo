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
	//db.Query("INSERT INTO tStudent (RegNo,Name,Branch) VALUES('20','Sourav Sharma','CSE')")
	r := mux.NewRouter()

	r.HandleFunc("/showAll", showAll).Methods("GET")
	r.HandleFunc("/show/{id}", showEntry).Methods("GET")
	r.HandleFunc("/add", addEntry).Methods("POST")
	r.HandleFunc("/update/{id}", updateEntry).Methods("PUT")
	r.HandleFunc("/delete/{regNo}", delEntry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8006", r))

}
func dbConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "Sourav"
	dbPass := "rootpass11"
	dbName := "studentDB"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

}
func delEntry(writer http.ResponseWriter, request *http.Request) {
	db := dbConnection()
	stmt, err := db.Prepare("DELETE FROM tStudent WHERE RegNo=?")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(request)

	stmt.Exec(params["regNo"])
	result, err := db.Query("SELECT * FROM tStudent WHERE RegNo =" + params["regNo"])
	if err != nil {
		panic(err.Error())
	}
	var obj Student
	for result.Next() {

		result.Scan(&obj.Sno, &obj.RegNo, &obj.Name, &obj.Branch)
		//personSlice = append(personSlice, obj)

		fmt.Println(obj)
	}
	json.NewEncoder(writer).Encode(obj)

}

func updateEntry(writer http.ResponseWriter, request *http.Request) {

}

func addEntry(writer http.ResponseWriter, request *http.Request) {
	//writer.Header().Set("Content-Type","application/json")
	var studentObject Student
	_ = json.NewDecoder(request.Body).Decode(&studentObject)

	db := dbConnection()

	stmt, err := db.Prepare("INSERT INTO tStudent (RegNo,Name,Branch) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(studentObject.RegNo, studentObject.Name, studentObject.Branch)
	defer db.Close()
	defer stmt.Close()
	json.NewEncoder(writer).Encode(studentObject)

}

func showEntry(writer http.ResponseWriter, request *http.Request) {
	db := dbConnection()
	params := mux.Vars(request)
	result, err := db.Query("SELECT * FROM tStudent WHERE RegNo =" + params["id"])
	if err != nil {
		panic(err.Error())
	}
	var obj Student
	for result.Next() {

		result.Scan(&obj.Sno, &obj.RegNo, &obj.Name, &obj.Branch)
		//personSlice = append(personSlice, obj)

		fmt.Println(obj)
	}
	json.NewEncoder(writer).Encode(obj)

}

func showAll(writer http.ResponseWriter, request *http.Request) {
	db := dbConnection()
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
	defer result.Close()
	defer db.Close()
	json.NewEncoder(writer).Encode(&slice)
}

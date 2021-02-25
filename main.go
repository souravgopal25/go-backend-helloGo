package main

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"
)

//fmt is formmater package
//func is function

//Greet Struct(Model)
type Greet struct {
	ID   string
	NAME string
}

//Init greetVar as Slice Greet struct
var greetSlice []Greet
var id = 0

func main() {
	//Init Router
	r := mux.NewRouter()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//MockData
	greetSlice = append(greetSlice, Greet{ID: "1", NAME: "Sourav Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "2", NAME: "Shyam Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "3", NAME: "Ram Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "4", NAME: "Bantai Sharma"})

	//Route Handler/Endpoint
	r.HandleFunc("/api/greet", greet).Methods("GET")
	r.HandleFunc("/api/{name}", greetWithName).Methods("GET")
	r.HandleFunc("/api/{name}", greetUSer).Methods("GET")
	//listen&Serve
	log.Fatal(http.ListenAndServe(":8002", r))
	fmt.Println("Hello Sourav")
}

func greetUSer(writer http.ResponseWriter, request *http.Request) {
	/*writer.Header().Set("Content-Type","application/json")
	params:=mux.Vars(request)
	for i,s greetSlice{

	}*/

}

func greetWithName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id++
	greet := Greet{
		ID:   strconv.Itoa(id),
		NAME: params["name"],
	}

	json.NewEncoder(writer).Encode(greet)

}

func greet(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//params:=mux.Vars(request)
	/*for _,item:=range greetSlice{
		json.NewEncoder(writer).Encode(item)

	}*/
	json.NewEncoder(writer).Encode(&greetSlice)

}

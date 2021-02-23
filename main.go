package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "net/http"
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

func main() {
	//Init Router
	r := mux.NewRouter()
	//MockData
	greetSlice = append(greetSlice, Greet{ID: "1", NAME: "Sourav Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "2", NAME: "Shyam Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "3", NAME: "Ram Sharma"})
	greetSlice = append(greetSlice, Greet{ID: "4", NAME: "Bantai Sharma"})

	//Route Handler/Endpoint
	r.HandleFunc("/api/greet", greet).Methods("GET")
	r.HandleFunc("/api/{name}", greetWithName).Methods("GET")

	//listen&Serve
	log.Fatal(http.ListenAndServe(":8001", r))
	fmt.Println("Hello Sourav")
}

func greetWithName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	greet := Greet{
		ID:   "10",
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

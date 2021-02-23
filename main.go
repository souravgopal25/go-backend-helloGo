package main

import (
	_ "encoding/json"
	"github.com/gorilla/mux"
	_ "log"
	_ "net/http"
	_ "math/rand"
	_ "strconv"
	"fmt"
	_ "github.com/gorilla/mux"
)

//fmt is formmater package
//func is function

func main() {

	r := mux.NewRouter()

	fmt.Println("Hello Sourav")
}

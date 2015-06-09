package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	http.Handle("/", r)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("foo", "bar")
	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

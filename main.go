package main

import (
	"fmt"
	"log"
	"main/mongodb"
	"net/http"
)

const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/mongodb", handleMongodb)
	http.HandleFunc("/mongodb/something", handleMongodb)

	fmt.Printf("Starting server at port 4000\n")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}

func handleMongodb(w http.ResponseWriter, r *http.Request) {
	err, data := mongodb.Controller(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

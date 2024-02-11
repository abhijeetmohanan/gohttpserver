package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func groot() http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp := "Hostname = " + hostname
		w.Write([]byte(resp))

	}
	// pending calculate inbound request
	return http.HandlerFunc(handleFunc)

}
func main() {

	http.Handle("/", groot())
	log.Fatal(http.ListenAndServe(":8080", nil))

}

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

var hits int = 0

func counter(count *int) {
	*count = *count + 1
}

func groot() http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Print(err)
		}
		resp := "Hostname = " + hostname + "Hits on Server: " + strconv.Itoa(hits)
		w.Write([]byte(resp))
		counter(&hits)
		log.Printf("Hostname = %s : Hit Count : %d", hostname, hits)
	}

	return http.HandlerFunc(handleFunc)

}
func main() {

	// each request is printing twice 

	http.Handle("/", groot())
	log.Print("Starting Server at 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)

}

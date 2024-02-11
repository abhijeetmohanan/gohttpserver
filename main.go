package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"log"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	http.Handle("/", getRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

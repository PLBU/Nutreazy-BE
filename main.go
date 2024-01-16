package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pgURL, err := pq.ParseURL(os.Getenv("URL"))
	logFatal(err)
	
	handler := func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello, this is BE Go lang server!")
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
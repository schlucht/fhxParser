package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handl)
	http.ListenAndServe(":8080", nil)
}

func handl(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hallo Welt!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", n)
}

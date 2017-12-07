package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	port := func() string { if len(os.Args) > 1 { return os.Args[1] } else { return "3000" } }()
	log.Println(fmt.Sprint("Listening on ", port))
	http.ListenAndServe(fmt.Sprint(":", port), nil)
}

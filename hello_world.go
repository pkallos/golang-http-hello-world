package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)
    fmt.Printf("Starting server on port :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

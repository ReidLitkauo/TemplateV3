package main

import (
    //"fmt"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.ListenAndServe(":8080", nil)
}


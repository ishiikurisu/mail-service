package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", index)
    log.Print("Serving at 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    x := r.Form["from"][0]
    fmt.Fprintf(w, "Hello %s", x)
}

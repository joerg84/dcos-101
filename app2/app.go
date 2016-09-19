package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to DC/OS 101! \n")
    fmt.Fprintf(w, "Running on host '"+  os.Getenv("HOST") + "' and port '" + os.Getenv("PORT0") + "'")
}

func main() {
    fmt.Println("Starting DC/OS-101 App ")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + os.Getenv("PORT0") , nil)
)


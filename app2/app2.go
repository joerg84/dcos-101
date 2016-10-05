package main

import (
    "fmt"
    "net/http"
    "os"

    "gopkg.in/redis.v4"
)

func getKeys(w http.ResponseWriter) {
    // Connect to redis.
    client := redis.NewClient(&redis.Options{
        Addr:     "redis.marathon.l4lb.thisdcos.directory:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, _ := client.Ping().Result()
    fmt.Fprintf(w, "Pong " + pong +"\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to DC/OS 101! \n")
    fmt.Fprintf(w, "Running on node '"+  os.Getenv("HOST") + "' and port '" + os.Getenv("PORT0") + "' \n")
    getKeys(w)
}

func main() {
    fmt.Println("Starting DC/OS-101 App ")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + os.Getenv("PORT0") , nil)
}


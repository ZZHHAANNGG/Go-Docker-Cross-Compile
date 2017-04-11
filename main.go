package main

import (
       "io"
       "net/http"
       "os"
       "log"
)

func hello(w http.ResponseWriter, r *http.Request) {
     io.WriteString(w, "Hello world!djd")
}

func main() {
     log.SetOutput(os.Stdout)

     port := ":" + os.Getenv("PORT")
     http.HandleFunc("/", hello)

     log.Printf("\n Application is listening on %v\n", port)

     http.ListenAndServe(port, nil)
}

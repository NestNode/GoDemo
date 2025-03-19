package main

import (
    "fmt"
    "net/http"
    "github.com/sirupsen/logrus"
)

func main() {
    fmt.Println("test")
    
    log := logrus.New()
    log.Info("Starting server...")
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.WithField("path", r.URL.Path).Info("Request received")
        w.Write([]byte("Hello, World with Logging!"))
    })

    log.Fatal(http.ListenAndServe(":24059", nil))
}

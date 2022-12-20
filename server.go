package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    requestURL := "http://google.com"
    res, err := http.Get(requestURL)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        fmt.Fprintf(w, "call google failed: %s\n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    fmt.Printf("client: status code: %d\n", res.StatusCode)
    fmt.Fprintf(w, "%s\n", res)
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":80", nil)
}

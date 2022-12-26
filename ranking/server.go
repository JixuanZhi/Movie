
package main

import (
    "fmt"
    "net/http"
)


type Recommender struct {
	name string
}

func (reco *Recommender) handler() http.Handler{
    handler := http.NewServeMux()
    handler.HandleFunc("/hello", reco.hello)
    handler.HandleFunc("/headers", reco.headers)
    return handler
}


func (reco *Recommender) hello(w http.ResponseWriter, req *http.Request) {
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

func  (reco *Recommender) headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
	reco := Recommender{name: "test"}
	//std1 := Student{name: "Vani"}

    fmt.Println(reco)
    //http.HandleFunc("/hello", hello)
    //http.HandleFunc("/headers", headers)
    http.ListenAndServe(":80", reco.handler())
}
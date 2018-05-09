package main

import (
  //  "html/template"
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)



func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Hello\n")
    fmt.Fprintf(w,"Hello")
}

func main() {
    r := mux.NewRouter()
    //r.HandleFunc("/",indexHandler);
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.ListenAndServe(":8080",r);
}

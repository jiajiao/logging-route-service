package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloGO(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() 
	fmt.Println(r.Form) 
	fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello GO!") 
}

func main() {
	http.HandleFunc("/",sayHelloGO)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}
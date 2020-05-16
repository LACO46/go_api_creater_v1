package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "log"
	"net/http"
	"fmt"
	"reflect"
)

func main() {
	api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    api.SetApp(rest.AppSimple(hello))
    log.Fatal(http.ListenAndServe(":80", api.MakeHandler())) 
}

func hello(w rest.ResponseWriter, r *rest.Request) {
	var json map[string]string = map[string]string{"Body": "Hello World!"}
	w.WriteJson(json)
}
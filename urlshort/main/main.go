package main

import (
	"fmt"
	"urlshort"
	"net/http"
)

func main(){
	mux := defaultMux() 
		pathsToUrls:= map[string]string{
			"/urlshort-godoc":"https://godoc.....",
			"/yaml-godoc": "https://godoc......",
		}

		mapHandler:= urlshort.MapHandler(pathsToUrls,mux)
		yaml:=`
		-path: /urlshort
		 url: https://godoc.........
		-path: /urlshort-final
		 url: https://godoc........
		 `
		yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil{
			panic(err)
		}

		fmt.Println("Starting the server on : 8080")
		http.ListenAndServe(":8080", yamlHandler)
	}

	func defaultMux() *http.ServeMux{
		mux:= http.NewServeMux()
		mux.HandleFunc("/", hello)
		return mux
	}

	func hello(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w , "Hello, World")
	}
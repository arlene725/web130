package main

//Take the code at this github repo and have images serve from GCS
//https://github.com/GoesToEleven/html-css/tree/master/074_lbr-homage/02_below-the-fold

import("fmt"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request){

}

func main(){

	http.HandleFunc("/", foo)
	http.HandleFunc("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}
//Create a webpage that serves at localhost:8080
// and will display the name in the url when the url is localhost:8080/name
// - use req.URL.Path to do this
package main

import(
	"net/http"
	"strings"
	"io"
)

func printName(res http.ResponseWriter, req * http.Request){
	name:= strings.Split(req.URL.Path, "/")
	io.WriteString(res, name[1])
}

func main(){
	http.HandleFunc("/",printName)
	http.ListenAndServe(":8080", nil)
}

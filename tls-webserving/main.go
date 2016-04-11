package main
//Arlene Cazarez
//create a web page which serves at localhost over https using TLS
//$go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
//int the folder as this file

import("io"
	"net/http"

)

func tlsServe(res http.ResponseWriter, req *http.Request ){
	io.WriteString(res, "success")
}

func main(){
	http.HandleFunc("/", tlsServe)
	go http.ListenAndServe(":8085", http.RedirectHandler("https://localhost:443/", 301))
	http.ListenAndServeTLS(":443","cert.pem","key.pem", nil)
}
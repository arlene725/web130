//Create a webpage which uses a cookie to track the number
//of visits of a user. Display the number of visits.
//Make sure that the favicon.ico requests are not also incrementing the number of visits.

package main

import(
	"io"
	"net/http"
	"strconv"
)

func cookieF(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("cookie-use")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name:"cookie-use",
			Value: "0",
		}
	}
	num, _ :=strconv.Atoi(cookie.Value)
	num++
	cookie.Value=strconv.Itoa(num)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)

}

func favIgnore(res http.ResponseWriter, req * http.Request){
	//this function takes care to ignore favicon.ico requests that would increment the number
	//each visit
}

func main(){
	http.HandleFunc("/", cookieF)
	http.HandleFunc("/favicon.ico", favIgnore)
	http.ListenAndServe(":8080", nil)
}

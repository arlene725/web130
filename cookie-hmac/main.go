//Create a webpage which writes a cookie to the client's machine.
//Though this is NOT A BEST PRACTICE, you will store some session data in the cookie.
//Make sure you use HMAC to ensure that session data is not changed by a user.

package main

import ("net/http"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
)

func getCode(data string) string{
	h := hmac.New(sha256.New, []byte("key"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func cookiehmac(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name:"session",
			Value:"",
			//Secure: true,
			HttpOnly: true,
		}
	}
	if req.FormValue("name")!=""{
		needsSalt := req.FormValue("name")
		cookie.Value=needsSalt + " | " + getCode(needsSalt)
	}
	fmt.Fprint(res,`
		<!DOCTYPE html>
		<html>
			<body>
				<form method = "POST">
					`+cookie.Value+`
					<input type = "text" name = "name">
					<input type = "submit">
				</form>
			</body>
		</html>
	`)
}

func main(){
	http.HandleFunc("/", cookiehmac)
	http.ListenAndServe(":8080", nil)
}
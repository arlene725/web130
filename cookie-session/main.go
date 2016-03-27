//Arlene Cazarez
//Writes a cookie to the client's machine
//cookie should create a session
//use UUID, HttpOnly, and Secure
//comment Secure out

package main

import (//"io"
	"fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
)

func sessionCookie(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err!= nil{
		id, _:=uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly:true,
		}
		http.SetCookie(res, cookie)
	}
	//this next line prints out cookie id onto the webpage:
	//io.WriteString(res ,cookie.Value)

	//this next line prints cookie info to the terminal:
	fmt.Println("cookie info:", cookie)
}

func favIgnore(res http.ResponseWriter, req *http.Request){}

func main(){
	http.HandleFunc("/", sessionCookie)
	http.HandleFunc("/favicon.ico", favIgnore)
	http.ListenAndServe(":8080", nil)
}


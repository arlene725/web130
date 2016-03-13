package main

//PROJECT STEP 2 - have the application write a cookie called "session-fino" with a UUID.
// The cookie should serve HttpOnly and you should have the "Secure" flag set also though
// comment the "Secure" flag out as we're not using https.

import(
	"html/template"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"log"
)

func foo(w http.ResponseWriter, r *http.Request){

	tpl, err:= template.ParseFiles("temp.html")

	if err !=nil{
		log.Fatalln(err)
	}

	cookie, err := r.Cookie("session-fino")

	if err != nil{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	err = tpl.Execute(w, nil)
	if err != nil{
		log.Fatalln(err)
	}
}


func main(){

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

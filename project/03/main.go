package main

//PROJECT STEP 3 - continuing to build our application,
//create a template which is a form. The form should gather the user's name and age.
// Store the user's name and age in the cookie.

import(
"html/template"
"net/http"
"github.com/nu7hatch/gouuid"
"log"
)

func foo(w http.ResponseWriter, r *http.Request){

	tpl, err:= template.ParseFiles("index.html")

	if err !=nil{
		log.Fatalln(err)
	}

	name:= r.FormValue("Name:")
	age:= r.FormValue("Age:")
	cookie, err := r.Cookie("session-fino")

	if err != nil{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
		Name: "session-fino",
		Value: id.String()+"|"+name+age,
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

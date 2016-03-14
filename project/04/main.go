package main

//PROJECT STEP 4 - refactoring our application, create a new data type called "user" which has fields for the user's name and age.
//When you receive the user's name and age form submission, create a variable of type "user" then put those values from the form
//submission into the fields for that variable. Marshal your variable of type "user"  to JSON.
//Encode that JSON to base64.
//Store that value in the cookie.

import(
	"html/template"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"log"
	"encoding/base64"
	"encoding/json"
)

type user struct{
	Name string
	Age string
}

func foo(w http.ResponseWriter, r *http.Request){

	tpl, err:= template.ParseFiles("index.html")
	if err !=nil{
		log.Fatalln(err)
	}

	name:= r.FormValue("Name:")
	age:= r.FormValue("Age:")

	x:= user{
		Name: r.FormValue("Name:"),
		Age: r.FormValue("Age:"),
	}

	b, err := json.Marshal(x)
	if err != nil{
		log.Fatalln("error:", err)
	}
	y :=base64.StdEncoding.EncodeToString(b)

	cookie, err := r.Cookie("session-fino")

	if err != nil{
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			//storing the name and age into the cookie value
			Value: id.String()+"|"+name+age,
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	cookie.Value = y
	err = tpl.Execute(w, nil)

	if err != nil{
		log.Fatalln(err)
	}
}

func main(){

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

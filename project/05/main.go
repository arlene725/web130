package main

import("net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/base64"
	"encoding/json"
	"crypto/sha256"
	"crypto/hmac"
	"io"
	"fmt"
)


type User struct{
	Name string
	Age string
}

func getCode(data string) string{
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func servePage(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	name := req.FormValue("name")
	age := req.FormValue("age")

	X := User{
		Name: name,
		Age: age,
	}

	bs, err := json.Marshal(X)
	if err != nil{
		fmt.Println(err)
	}

	json := base64.StdEncoding.EncodeToString(bs)


	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	cookie = &http.Cookie{
		Name: "session",
		Value: id.String() + name + age + json + getCode(id.String()),
		HttpOnly: true,
	}
	http.SetCookie(res, cookie)
	tpl.Execute(res, nil)
}
}

func main() {

	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)

}

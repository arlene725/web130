//PROJECT STEP 8 - Allow the user to logout.
//Show a log-in button when the user is not logged-in.
//Show a log-out button only when the user is logged in.
package main

import("fmt"
	"net/http"
	"html/template"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"strings"
	"log"
)

type user struct{
	Name string
	Age string
	LogStatus bool
}

var tpl *template.Template

func index(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-id")
	if err != nil{
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if strings.Count(cookie.Value, "|") !=2{
		cookie= newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value){
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}
	u:= decodeUser(cookie)

	if req.Method == "POST"{
		u.LogStatus=true
		u.Name = req.FormValue("name")
		u.Age = req.FormValue("age")

		xs := strings.Split(cookie.Value, "|")
		id:= xs[0]

		cookie= currentVisitor(u, id)
		http.SetCookie(res, cookie)
	}

	fmt.Println("deBugging", u)
	temperr :=tpl.ExecuteTemplate(res, "index.html", u)
	if temperr != nil{
		http.Error(res, temperr.Error(), http.StatusInternalServerError)
	}
}


func login(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-id")
	if err != nil{
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if strings.Count(cookie.Value, "|") !=2{
		cookie= newVisitor()
		http.SetCookie(res, cookie)
	}

	if tampered(cookie.Value){
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}

	if req.Method == "POST" && req.FormValue("password") == "secret"{
		u:= decodeUser(cookie)
		u.LogStatus = true
		u.Name = req.FormValue("username")

		xs:= strings.Split(cookie.Value, "|")
		id:= xs[0]

		cookie:= currentVisitor(u, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)


}

func logout(res http.ResponseWriter, req *http.Request){
	cookie:= newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
}


func decodeUser(c *http.Cookie) user {

	xs:= strings.Split(c.Value, "|")

	usrData := xs[1]

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil{
		log.Println(err)}
	var u user
	err = json.Unmarshal(bs, &u)
	if err !=nil{
		fmt.Println(err)
		return user{}
	}

	return u
}


	func newVisitor() *http.Cookie{
		u := user{
		Name:   "",
		Age:  "",
		LogStatus: false,
		}
		bs, err :=json.Marshal(u)
		if err != nil{
		fmt.Println("error:", err)

		}

		id, _:= uuid.NewV4()
		return makeCookie(bs, id.String())
	}


	func currentVisitor(u user, id string) *http.Cookie{
		bs, err :=json.Marshal(u)
		if err != nil{
		fmt.Println("error:", err)

		}
		return makeCookie(bs, id)
	}

	func makeCookie(mm []byte, id string) *http.Cookie{
		b64 := base64.URLEncoding.EncodeToString(mm)
		code := getCode(b64)
		cookie:= &http.Cookie{
		Name: "session-id",
		Value: id + "|"+ b64 +"|"+code,
		HttpOnly: true,
		}

		return cookie
	}

	func getCode(data string) string {
		h := hmac.New(sha256.New, []byte("H3110w0rld"))
		io.WriteString(h, data)
		return fmt.Sprintf("%x", h.Sum(nil))
	}

	func tampered(s string) bool {
		xs := strings.Split(s, "|")
		usrData := xs[1]
		usrCode := xs[2]
		if usrCode != getCode(usrData) {
		return true
		}
		return false
	}

	func main(){
		tpl, _ = template.ParseGlob("templates/*.html")

		http.Handle("/favicon.ico", http.NotFoundHandler())
		http.HandleFunc("/", index)
		http.HandleFunc("/login", login)
		http.HandleFunc("/logout", logout)
		http.ListenAndServe(":8080", nil)
	}
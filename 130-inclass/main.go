package main

//instructions
//store a uuid in a cookie; value:uuid
//store the uuid in memcache; key:uuid, value:your name
//retrive the uuid from the cookie
//retirive the uuid & value from memcache

import ("fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine"
)

func init(){
	http.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req *http.Request){
	cookieId :=setcookie(res)
	setmemcache(cookieId, req)

	fmt.Println(cookieId)
}

func setcookie(res http.ResponseWriter) string{
	id, _:=uuid.NewV4()
	cookie := &http.Cookie{
		Name: "session",
		Value: id.String(),
		//Secure: true,
		HttpOnly:true,
	}
	http.SetCookie(res, cookie)
	return cookie.Value
}

func setmemcache(id string, req *http.Request){
	ctx:= appengine.NewContext(req)
	item1 := memcache.Item{
		Key: "Cazarez",
		Value: id,
	}
}

//GCS assignment
package main

import("google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"net/http"
)

const gcsBucket = ""

func init(){
	http.HandleFunc("/", handler)

}

func handler(res http.ResponseWriter, req *http.Request){
	ctx:= appengine.NewContext(req)

	if req.URL.Path != "/"{
		http.NotFound(res, req)
		return
	}

	html:=
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
		</form>

	if req.Method == "POST"{

		mpf, hdr, err := req.FormFile("")
	}
}
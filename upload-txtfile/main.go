//Create a webpage that serves a form and allows the user to upload a txt file.
// You do not need to check if the file is a txt; bad programming but
// just trust the user to follow the instructions. Once a user has uploaded a txt file,
// copy the text from the file and display it on the webpage.
// Use req.FormFile and io.Copy to do this
//

package main

import(
	"io"
	"net/http"
	"html/template"
	"log"
	"fmt"
)

func upload(res http.ResponseWriter, req *http.Request){

	tpl, err := template.ParseFiles("./temp.html")
	if err!=nil{
		log.Fatal("error:", err)
	}

	tpl.Execute(res, nil)

	if req.Method =="POST"{
		_, src, err:=req.FormFile("name")
		if err != nil{
			fmt.Println(err)
		}
		dst, err:= src.Open()
		if err != nil{
			fmt.Println(err)
		}

		io.Copy(res, dst)
	}

}

func main(){
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8080", nil)
}
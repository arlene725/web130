package main

import ("html/template"
	"net/http"
)
//import "fmt"

var tpl *template.Template

func init(){
	var err error //need to declare error because we declared tpl outside
			//and cant assign tpl again with :=
	tpl, err = template.ParseFiles("templates/index.gohtml")
	if err != nil {
		//log.Fatalln(err) //shuts down a web server and prints out error and time stamp
		panic(err)
	}

}

func main(){
	//http.Handle("/img", http.StripPrefix("/img", http.FileServer(http.Dir("pics"))))
	http.HandleFunc("/", index)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("public/css"))))
	http.Handle("/pics/", http.StripPrefix("/pics", http.FileServer(http.Dir("public/pics"))))
	http.ListenAndServe(":8080", nil)

}

func index (res http.ResponseWriter, req *http.Request){
	tpl.Execute(res, nil)
}
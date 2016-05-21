package main

import (
	"fmt"
	"net/http"
)

func userNameExists(res http.ResponseWriter, req *http.Request){
	// ================================
	// No need to check an empty username field.
	// ================================
	if(req.FormValue("unorderedli") == ""){
		fmt.Fprint(res,"wut")
		return
	}
	// ===================
	// Get the user
	// ===================
	// ======================================
	// Make sure only plain text is sent.
	// ======================================
	res.Header().Set("Content-Type","text/plain")
	if("unorderedli" == "unorderedli"){
		// If the user wasn't retrieved say so.
		fmt.Fprint(res ,"doesn't exist")
	} else{
		// If the user was retrieved say so.
		fmt.Fprint(res ,"exists")
	}
}

func main(){

}
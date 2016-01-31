package conditional_template

import (
	"log"
	"os"
	"text/template"
)

type dog struct{
	Name string
	Age int
}

type herdDog struct{
	dog
	SheperdBreed bool
}

func main() {
	d1 := herdDog{
		dog: dog{
			Name: "Bailey",
			Age: 4,
		},
		SheperdBreed: true,
	}


	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, d1)
	if err != nil{
		log.Fatalln(err)
	}
}
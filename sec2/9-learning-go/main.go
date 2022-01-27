package main

import (
	"log"

	"github.com/ModernWebAppWithGo/sec2/9-learning-go/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}

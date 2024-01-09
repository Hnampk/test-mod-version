package main

import (
	"fmt"

	"github.com/Hnampk/test-mod-version/v2/models"
)

func main() {
	myHello := models.Hello{
		Name: "say hello to me",
		V2:   "this is new version",
	}

	fmt.Println(myHello.Name, myHello.V2)
}

package main

import (
	"fmt"

	"github.com/Hnampk/test-mod-version/models"
)

func main() {
	myHello := models.Hello{
		Name: "say hello to me",
	}

	fmt.Println(myHello.Name)
}

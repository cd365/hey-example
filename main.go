package main

import (
	"fmt"
	"github.com/cd365/hey-example/examples"
)

func main() {

	db, err := examples.NewDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = db.RunTest(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success")
}

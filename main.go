package main

import "easy_demo/router"

func main() {

	router := router.Initrouter()

	err := router.Run(":9999")
	if err != nil {
		panic(err)
	}

}

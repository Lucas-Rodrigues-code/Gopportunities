package main

import (
	"github.com/Lucas-Rodrigues-code/Gopportunities/config"
	"github.com/Lucas-Rodrigues-code/Gopportunities/router"
)

func main() {
	err error := config.Init()
	if err != nil {
		panic(err)
		return
	}
	router.Initialize()

}

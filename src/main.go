package app

import (
	"fmt"
	"hello-gin/src/server"
)

func StartApplication() {
	fmt.Println("Hello World")

	server.Init()

}

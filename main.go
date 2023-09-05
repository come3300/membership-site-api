package main

import (
	"go_blog/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}

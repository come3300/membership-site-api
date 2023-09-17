package main

import (
	controller "go-jet-env/controllers"
)

func main() {
	
	router := controller.GetRouter()
	router.Run(":88")
}

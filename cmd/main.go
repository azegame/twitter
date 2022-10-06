package main

import (
	"twitter/controller"  
)


func main() {
	router := controller.GetRouter()
	router.Run()
}
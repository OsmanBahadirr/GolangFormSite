package main

import (
	"GolangForm/controller"
	"GolangForm/model"
)

func main() {
	model.Init()
	controller.Start()
}

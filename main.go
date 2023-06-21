package main

import (
	"GolangForm/controller"
	"GolangForm/model/core"
)

func main() {
	core.Init()
	controller.Start()

}

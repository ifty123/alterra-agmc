package main

import (
	"day2-task1/config"
	"day2-task1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}

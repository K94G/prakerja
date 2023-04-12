package main

import (
	"os"
	"prakerja_kg/config"
	"prakerja_kg/route"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "6000"
	}
	return port
}

func main() {
	config.InitDB()
	e := route.InitRoute()
	// e.Logger.Fatal(e.Start(":6000"))
	e.Logger.Fatal(e.Start(":" + getPort()))
}

package main

import (
	"awesomeProject/cmd/database"
	"awesomeProject/cmd/engine"
	"awesomeProject/cmd/routes"
)

func main() {
	// INIT DATABASE (sqlite3)
	database.InitDB()

	// INIT API ROUTING ENGINE
	engine.InitEngine()

	// INIT API ROUTES
	routes.InitRoutes()

	// START SERVER
	engine.RunEngine(":6900")
}

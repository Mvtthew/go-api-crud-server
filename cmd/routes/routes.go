package routes

import (
	"awesomeProject/cmd/controllers"
	"awesomeProject/cmd/engine"
)

func initPingRoutes() {
	engine.REngine.GET("/ping", controllers.HandlePingPong)
}

func initUsersRoutes()  {
	engine.REngine.GET("/users", controllers.GetAllUsers)
	engine.REngine.POST("/users", controllers.CreateUser)
}

func InitRoutes() {
	initPingRoutes()
	initUsersRoutes()
}

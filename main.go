package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/l-giuliani/networkEngine/controllers"
	"github.com/l-giuliani/networkEngine/services"
)

func main()  {
	services.InitSystem()

	engine := gin.Default()
	controllers.InitGameControllers(engine)
	engine.Run()
}
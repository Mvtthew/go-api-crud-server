package engine

import (
	"github.com/gin-gonic/gin"
)

var REngine *gin.Engine

func InitEngine()  {
	REngine = gin.Default()
}

func RunEngine(address string) {
	REngine.Run(address)
}

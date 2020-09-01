package main

import (
	gin_recovery "gin-recovery"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.New()
	gin_recovery.New(gin_recovery.RecoveryConfig{
		RecoveryHandler: HandlerFunc,
		ErrorReporting: &gin_recovery.ErrorReportingConfig{
			ProjectId:   "PROJECT_ID",
			ServiceName: "SERVICE_NAME",
		},
	})
	c.Use(gin_recovery.Recovery())
	c.GET("/panic", func(c *gin.Context) {
		panic("This is panic")
	})
	c.Run(":4040")
}

func HandlerFunc(ctx *gin.Context, err interface{}) {
	ctx.JSON(500, err)
}

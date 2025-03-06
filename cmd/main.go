package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vicjeremy/shorten-url/api"
)

func main() {
	api.InitDB()
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//r.Static("/web", "./web")
	//r.LoadHTMLGlob("website/*")
	shorten := r.Group("/shorten")
	{
		shorten.POST("/", api.CreateUrl)
		shorten.GET("/:code", api.GetUrl)
		shorten.PUT("/:code", api.UpdateUrl)
		shorten.DELETE("/:code", api.DeleteUrl)
		shorten.GET("/:code/stats", api.StatsUrl)
	}
	r.Run(":3000")
}

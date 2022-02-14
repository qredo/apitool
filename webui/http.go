package webui

import (
	"fmt"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", GetIndex)
	router.POST("/sign", Sign)
	router.POST("/send", Send)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	fmt.Printf("WebUI listening on http://127.0.0.1:%s\n", port)

	if err := router.Run(":" + port); err != nil {
		fmt.Println(err)
	}
}

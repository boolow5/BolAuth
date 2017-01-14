package main

import (
	"github.com/boolow5/BolAuth/conf"
	"github.com/boolow5/BolAuth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	conf.InitConfig()
	router := gin.Default()

	/*router.LoadHTMLGlob("templates/*")
	router.GET("/", func(this *gin.Context) {
		this.Redirect(http.StatusMovedPermanently, "/app")
	})
	router.GET("/app/*vuepath", controllers.Index)
	router.Static("/static", "./public/dist")*/

	router.GET("/", controllers.Index)

	apiV1 := router.Group("/api/v1")
	{
		//todos
		apiV1.GET("/todos", controllers.AllTodos)
		apiV1.POST("/todo", controllers.AddTodo)
		apiV1.DELETE("/todo", controllers.DeleteTodo)
		apiV1.PUT("/todo", controllers.UpdateTodo)

		//
		apiV1.GET("/roles", controllers.AllRoles)
		apiV1.POST("/role", controllers.AddRole)
		apiV1.DELETE("/role", controllers.DeleteRole)
		apiV1.PUT("/role", controllers.UpdateRole)
	}

	router.Run()
}

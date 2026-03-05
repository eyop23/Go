package routes

import "github.com/gin-gonic/gin"
import "github/eyop23/go_learn/controllers"
import "github/eyop23/go_learn/middleware"

func UserRoutes(r *gin.Engine){
   
   api:= r.Group("/api/user")
   {
      api.POST("/login",controllers.Login)
      api.POST("/register",controllers.Register)
   }
  protected := api.Group("/")
  protected.Use(middleware.AuthMiddleware());
  protected.GET("/", controllers.ListUser)
  protected.GET("/:id", controllers.GetUser)
}

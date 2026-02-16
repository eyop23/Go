package routes

import "github.com/gin-gonic/gin"
import "github/eyop23/go_learn/controllers"

func UserRoutes(r *gin.Engine){
   
   api:= r.Group("/api/user")
   {
   api.GET("/",controllers.ListUser)
   api.GET("/:id",controllers.GetUser)
   api.POST("/login",controllers.Login)
   api.POST("/register",controllers.Register)
   }
}

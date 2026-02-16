package main

import (
	"fmt"
	"os"

	"github/eyop23/go_learn/routes"
	"github/eyop23/go_learn/database"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)

func main(){

  godotenv.Load()

  database.ConnectPostgres()


   r:=gin.Default()

   routes.UserRoutes(r);

   port:= os.Getenv("PORT")

   fmt.Println("server listening on port:", port)

   r.Run(":" + port)

}
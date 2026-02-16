package database 

import (
	"log"
	"fmt"
	"os"
	"database/sql"

	_ "github.com/lib/pq"

)

var DB *sql.DB

func ConnectPostgres(){

	 db_user:= os.Getenv("DB_USER")
	 db_password:=os.Getenv("DB_PASSWORD")
	 database:=os.Getenv("DATABASE")


	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",db_user,db_password,database) 

	fmt.Println("connection string is :", connectionString)

	// connectionString:="postgres://postgres:postgres@localhost:5432/go_learn?sslmode=disable"

   db,err := sql.Open("postgres",connectionString)

   if err != nil {
	log.Fatal("Cannot open database:", err)
   }

   if err := db.Ping(); err != nil {
	 log.Fatal("Cannot connect to database:", err)
   }
   DB=db
   log.Println("connected")
}
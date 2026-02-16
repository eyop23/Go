package controllers

import (

	"net/http"
	"strconv"
	"fmt"
	"time"
	"database/sql"
	"errors"
	"os"



	"github/eyop23/go_learn/db"
	"github/eyop23/go_learn/database"
	"github/eyop23/go_learn/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"

	

)

	

func ListUser(c *gin.Context){
   q:=db.New(database.DB)

   users,err := q.ListUsers(c.Request.Context());
   if err != nil {
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound,gin.H{
			"error":"no user is found",
		})
		return

	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":err.Error(),
	})
	return
   }
   c.JSON(http.StatusOK,gin.H{
	"message":"ok",
	"users":users,
   })

}
func Register (c *gin.Context){
	var req dto.UserRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	q := db.New(database.DB);
	existingUser,err := q.GetUserByName(c.Request.Context(),req.Name)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existingUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to hash password"})
		return
	}

	user,err := q.CreateUser(c.Request.Context(), db.CreateUserParams{
		Name:req.Name,
		Age:req.Age,
		Password:string(hashedPassword),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	resp:=dto.UserResponse{
		ID:user.ID,
		Name:user.Name,
		Age:user.Age,
	}
	c.JSON(http.StatusCreated,resp)
}

func Login(c *gin.Context){
	var req dto.LoginRequest

	err := c.ShouldBindJSON(&req); if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	q:=db.New(database.DB)
   user,err :=q.GetUserByName(c.Request.Context(),req.Name)
   if err != nil {
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
   err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password))
   if err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	return
   }
   token,err := generateToken(user.ID)

   if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
	return
}
 c.JSON(http.StatusOK,dto.LoginResponse{
	Token:token,
 })

}
func GetUser(c *gin.Context){
	id_param := c.Param("id");

     id,err := strconv.ParseInt(id_param,10,32)

	 fmt.Printf("%T\n",id)
	 if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	 }

    q:=db.New(database.DB);

     user,err := q.GetUser(c.Request.Context(),int32(id));
	 if err != nil {
          
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound,gin.H{
				"error":"user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	 }

	 c.JSON(http.StatusOK,user)

}

func generateToken(userId int32) (string,error){
	secret:=os.Getenv("JWT_SECRET")
	expHours:=os.Getenv("JWT_EXPIRATION_HOURS")
	expTime:=time.Now().Add(24 * time.Hour)
	if expHours != "" {
		if h,err := strconv.Atoi(expHours); err == nil{
           expTime = time.Now().Add(time.Duration(h) * time.Hour)
		}
	}

	claims:=jwt.MapClaims{
		"user_id":userId,
		"exp":expTime.Unix(),
		"iat":time.Now().Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString([]byte(secret))
}
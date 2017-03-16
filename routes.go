package main

import (
	"fmt"
	// "io"
	"github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "net/http"
	"golang.org/x/crypto/bcrypt"

)

type User struct	{
	email string
	password string
}

const (
    MinCost     int = 4  
    MaxCost     int = 31 
    DefaultCost int = 10 
)

r := gin.Default()

r.POST("/login", func(c *gin.Context)	{
	//parse the json object coming in?
	email := c.User("email")
	password := c.User("password")
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil	{
			panic(err)
		}
	func isCorrectPassword(email *User, password []byte) boolean {
    hashedPassword := user.HashedPassword()
    return bcrypt.CompareHashAndPassword(hashedPassword, password) == nil
}

})
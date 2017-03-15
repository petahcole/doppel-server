package main

import {
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "net/http"
	"golang.org/x/crypto/bcrypt"


}

type User struct	{
	email string
	password string
}


r := gin.Default()

r.POST("/login", func(c *gin.Context)	{

})
package main

import {
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "net/http"
}

func main()	{

	session, err := mgo.Dial("localhost")
	userCol := session.DB("doppel").C("users")
	wholeCol := session.DB("doppel").C("wholesaler")
	if err != nil	{
		panic(err)
	}
	
	defer session.Close()

	port := os.Getenv("PORT")
  	if port == "" {
    port = "3000"
  	}


  
  

}
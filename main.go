package main

import (
	"fmt"
	"os"
	// "io"
	// "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
    // "net/http"
)

func main()	{

	session, err := mgo.Dial("localhost")
	// userCol := session.DB("doppel").C("users")
	// wholeCol := session.DB("doppel").C("wholesaler")
	if err != nil	{
		panic(err)
	}
	
	// session.Copy()
	defer session.Close()

	port := os.Getenv("PORT")
  	if port == "" {
    port = "3000"
  	}

	fmt.Println("hello from go")
  
  

}
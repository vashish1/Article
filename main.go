package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router :=gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	  router.Run()

}
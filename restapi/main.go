package main

import (
	"fmt"
	"log"
	"restapi/constants"
	"restapi/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	routes.AppRoutes(router)
	fmt.Println("server runnign on port",constants.Port)
	log.Fatal(router.Run(constants.Port))
}
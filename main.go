package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mari-muthu-k/gin-template/db"
	"github.com/mari-muthu-k/gin-template/routes"
	"github.com/mari-muthu-k/gin-template/utils"
	"gorm.io/gorm"
)



func init() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect mysql db
	err = db.ConnectRelationalDB("mysql", &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func main() {
	defer db.DisconnectRelationalDB()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(utils.GetCorsConfig())
	
	versionRouter := router.Group("/api/v1")
	{
		routes.NoAuthGroupRoutes(versionRouter)
	}
	

	if err := router.Run(); err != nil {
		log.Printf("error starting the server: %s\n", err)
	}
}

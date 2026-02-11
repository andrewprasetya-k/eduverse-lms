package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN is not set")
	}

	//db connection
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  dsn,
	// 	PreferSimpleProtocol: true, // Mengatasi error prepared statement pada Supabase Pooler
	// }), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database: " + err.Error())
	// }

	//initialize repo, service, handler

	//router setup
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// api := r.Group("/api")
	// {
	// 	//todo: routes group
	// }

	//run server
	r.Run(":8080")
}
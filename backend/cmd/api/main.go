package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//load env variables
	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN is not set")
	}

	//db connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Mengatasi error prepared statement pada Supabase Pooler
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	//initialize repo, service, handler
	schoolRepo := repository.NewSchoolRepository(db)
	schoolService := service.NewSchoolService(schoolRepo)
	schoolHandler := handler.NewSchoolHandler(schoolService)

	subjectRepo := repository.NewSubjectRepository(db)
	subjectService := service.NewSubjectService(subjectRepo, schoolRepo)
	subjectHandler := handler.NewSubjectHandler(subjectService)

	//router setup
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		// School Routes
		schools := api.Group("/schools")
		{
			schools.POST("", schoolHandler.CreateSchool)
			schools.GET("", schoolHandler.GetAllSchools)
			// schools.GET("/:id", schoolHandler.GetSchoolByID)
			schools.GET("/:schoolCode", schoolHandler.GetSchoolByCode)
			schools.PATCH("/:schoolCode", schoolHandler.UpdateSchool)
			schools.DELETE("/:schoolCode", schoolHandler.DeleteSchool)
		}

		// Subject Routes
		subjects := api.Group("/subjects")
		{
			subjects.POST("", subjectHandler.CreateSubject)
			subjects.GET("/:schoolCode", subjectHandler.GetAllSubjects)
			subjects.GET("/:schoolCode/:subjectCode", subjectHandler.GetSubjectByCode)
			subjects.PATCH("/:subjectCode", subjectHandler.UpdateSubject)
			subjects.DELETE("/:subjectCode", subjectHandler.DeleteSubject)
		}
	}

	//run server
	r.Run(":8080")
}
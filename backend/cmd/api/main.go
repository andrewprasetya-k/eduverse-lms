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

	// db connection
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

	academicYearRepo := repository.NewAcademicYearRepository(db)
	academicYearService := service.NewAcademicYearService(academicYearRepo, schoolService)
	academicYearHandler := handler.NewAcademicYearHandler(academicYearService)

	termRepo := repository.NewTermRepository(db)
	termService := service.NewTermService(termRepo)
	termHandler := handler.NewTermHandler(termService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	schoolUserRepo := repository.NewSchoolUserRepository(db)
	schoolUserService := service.NewSchoolUserService(schoolUserRepo)
	schoolUserHandler := handler.NewSchoolUserHandler(schoolUserService)


	//router setup
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		schoolAPI:=api.Group("/schools")
		{
			schoolAPI.POST("/", schoolHandler.CreateSchool)
			schoolAPI.GET("/", schoolHandler.GetSchools)
			schoolAPI.GET("/summary", schoolHandler.GetSchoolSummary)
			schoolAPI.GET("/check-code/:schoolCode", schoolHandler.CheckCodeAvailability)
			schoolAPI.GET("/:schoolCode", schoolHandler.GetSchoolByCode)
			schoolAPI.PATCH("/:schoolCode", schoolHandler.UpdateSchool)
			schoolAPI.PATCH("/restore/:schoolCode", schoolHandler.RestoreDeletedSchool)
			schoolAPI.DELETE("/:schoolCode", schoolHandler.DeleteSchool)
			schoolAPI.DELETE("/permanent/:schoolCode", schoolHandler.HardDeleteSchool)
		}

		academicYearAPI := api.Group("/academic-years")
		{
			academicYearAPI.POST("/", academicYearHandler.Create)
			academicYearAPI.GET("/", academicYearHandler.FindAll)
			academicYearAPI.GET("/:id", academicYearHandler.GetByID)
			academicYearAPI.GET("/school/:schoolCode", academicYearHandler.GetBySchool)
			academicYearAPI.PATCH("/:id", academicYearHandler.Update)
			academicYearAPI.PATCH("/activate/:id", academicYearHandler.Activate)
			academicYearAPI.PATCH("/deactivate/:id", academicYearHandler.Deactivate)
			academicYearAPI.DELETE("/:id", academicYearHandler.Delete)
		}

		termAPI := api.Group("/terms")
		{
			termAPI.POST("/", termHandler.Create)
			termAPI.GET("/", termHandler.FindAll)
			termAPI.GET("/:id", termHandler.GetByID)
			termAPI.GET("/academic-year/:academicYearId", termHandler.GetByAcademicYear)
			termAPI.PATCH("/:id", termHandler.Update)
			termAPI.PATCH("/activate/:id", termHandler.Activate)
			termAPI.PATCH("/deactivate/:id", termHandler.Deactivate)
			termAPI.DELETE("/:id", termHandler.Delete)
		}

		userAPI := api.Group("/users")
		{
			userAPI.POST("/", userHandler.Create)
			userAPI.GET("/", userHandler.FindAll)
			userAPI.GET("/:id", userHandler.GetByID)
			userAPI.PATCH("/:id", userHandler.Update)
			userAPI.DELETE("/:id", userHandler.Delete)
		}

		schoolUserAPI := api.Group("/school-users")
		{
			schoolUserAPI.POST("/enroll", schoolUserHandler.Enroll)
			schoolUserAPI.GET("/school/:schoolId", schoolUserHandler.GetMembersBySchool)
			schoolUserAPI.GET("/user/:userId", schoolUserHandler.GetSchoolsByUser)
			schoolUserAPI.DELETE("/:id", schoolUserHandler.Unenroll)
		}
	}

	//run server
	r.Run(":8080")
}
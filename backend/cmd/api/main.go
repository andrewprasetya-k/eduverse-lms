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
	schoolUserService := service.NewSchoolUserService(schoolUserRepo, schoolService)
	schoolUserHandler := handler.NewSchoolUserHandler(schoolUserService, schoolService)

	subjectRepo := repository.NewSubjectRepository(db)
	subjectService := service.NewSubjectService(subjectRepo, schoolService)
	subjectHandler := handler.NewSubjectHandler(subjectService, schoolService)

	rbacRepo := repository.NewRBACRepository(db)
	rbacService := service.NewRBACService(rbacRepo, schoolService)
	rbacHandler := handler.NewRBACHandler(rbacService)

	classRepo := repository.NewClassRepository(db)
	classService := service.NewClassService(classRepo, schoolService)
	classHandler := handler.NewClassHandler(classService)

	subjectClassRepo := repository.NewSubjectClassRepository(db)
	subjectClassService := service.NewSubjectClassService(subjectClassRepo)
	subjectClassHandler := handler.NewSubjectClassHandler(subjectClassService)

	enrollmentRepo := repository.NewEnrollmentRepository(db)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)

	mediaRepo := repository.NewMediaRepository(db)
	mediaService := service.NewMediaService(mediaRepo)
	mediaHandler := handler.NewMediaHandler(mediaService)

	attachmentRepo := repository.NewAttachmentRepository(db)
	attachmentService := service.NewAttachmentService(attachmentRepo)

	materialRepo := repository.NewMaterialRepository(db)
	materialService := service.NewMaterialService(materialRepo, attachmentService, mediaRepo)
	materialHandler := handler.NewMaterialHandler(materialService)

	feedRepo := repository.NewFeedRepository(db)
	feedService := service.NewFeedService(feedRepo, attachmentService)
	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	feedHandler := handler.NewFeedHandler(feedService, commentService)
	commentHandler := handler.NewCommentHandler(commentService)

	assignmentRepo := repository.NewAssignmentRepository(db)
	assignmentService := service.NewAssignmentService(assignmentRepo, attachmentService)
	assignmentHandler := handler.NewAssignmentHandler(assignmentService)

	logRepo := repository.NewLogRepository(db)
	logService := service.NewLogService(logRepo)
	logHandler := handler.NewLogHandler(logService)


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
			userAPI.PATCH("/change-password/:id", userHandler.ChangePassword)
			userAPI.DELETE("/:id", userHandler.Delete)
		}

		schoolUserAPI := api.Group("/school-users")
		{
			schoolUserAPI.POST("/enroll", schoolUserHandler.Enroll)
			schoolUserAPI.GET("/school/:schoolCode", schoolUserHandler.GetMembersBySchool)
			schoolUserAPI.GET("/user/:userId", schoolUserHandler.GetSchoolsByUser)
			schoolUserAPI.DELETE("/:userId", schoolUserHandler.Unenroll)
		}

		subjectAPI := api.Group("/subjects")
		{
			subjectAPI.POST("/", subjectHandler.Create)
			subjectAPI.GET("/", subjectHandler.FindAll)
			subjectAPI.GET("/:id", subjectHandler.GetByID)
			subjectAPI.GET("/school/:schoolCode", subjectHandler.GetBySchool)
			subjectAPI.GET("/school/:schoolCode/:subjectCode", subjectHandler.GetByCode)
			subjectAPI.PATCH("/:id", subjectHandler.Update)
			subjectAPI.DELETE("/:id", subjectHandler.Delete)
		}

		rbacAPI := api.Group("/rbac")
		{
			// Roles
			rbacAPI.POST("/roles", rbacHandler.CreateRole)
			rbacAPI.GET("/roles/school/:schoolCode", rbacHandler.GetRolesBySchool)
			rbacAPI.GET("/roles/:id", rbacHandler.GetRoleByID)
			rbacAPI.PATCH("/roles/:id", rbacHandler.UpdateRole)
			rbacAPI.DELETE("/roles/:id", rbacHandler.DeleteRole)
			rbacAPI.POST("/roles/permissions/:id", rbacHandler.SetRolePermissions)

			// Permissions
			rbacAPI.POST("/permissions", rbacHandler.CreatePermission)
			rbacAPI.GET("/permissions", rbacHandler.GetAllPermissions)

			// User Roles (Assignments)
			rbacAPI.POST("/user-roles", rbacHandler.AssignRole)
			rbacAPI.DELETE("/user-roles", rbacHandler.RemoveRole)
			rbacAPI.GET("/user-roles/user/:schoolUserId", rbacHandler.GetUserRoles)
			rbacAPI.PATCH("/user-roles/user/:schoolUserId", rbacHandler.UpdateUserRoles)
		}

		classAPI := api.Group("/classes")
		{
			classAPI.POST("/", classHandler.Create)
			classAPI.GET("/", classHandler.FindAll)
			classAPI.GET("/:id", classHandler.GetByID)
			classAPI.PATCH("/:id", classHandler.Update)
			classAPI.DELETE("/:id", classHandler.Delete)
		}

		subjectClassAPI := api.Group("/subject-classes")
		{
			subjectClassAPI.POST("/assign", subjectClassHandler.Assign)
			subjectClassAPI.GET("/class/:classId", subjectClassHandler.GetByClass)
			subjectClassAPI.GET("/:id", subjectClassHandler.GetByID)
			subjectClassAPI.DELETE("/:id", subjectClassHandler.Unassign)
		}

		enrollmentAPI := api.Group("/enrollments")
		{
			enrollmentAPI.POST("/enroll", enrollmentHandler.Enroll)
			enrollmentAPI.GET("/class/:classId", enrollmentHandler.GetByClass)
			enrollmentAPI.GET("/member/:schoolUserId", enrollmentHandler.GetByMember)
			enrollmentAPI.DELETE("/:id", enrollmentHandler.Unenroll)
		}

		mediaAPI := api.Group("/medias")
		{
			mediaAPI.POST("/metadata", mediaHandler.RecordMetadata)
			mediaAPI.GET("/:id", mediaHandler.GetByID)
			mediaAPI.DELETE("/:id", mediaHandler.Delete)
		}

		materialAPI := api.Group("/materials")
		{
			materialAPI.POST("/", materialHandler.Create)
			materialAPI.GET("/", materialHandler.FindAll)
			materialAPI.GET("/:id", materialHandler.GetByID)
			materialAPI.POST("/progress", materialHandler.UpdateProgress)
		}

		feedAPI := api.Group("/feeds")
		{
			feedAPI.POST("/", feedHandler.Create)
			feedAPI.GET("/class/:classId", feedHandler.GetByClass)
		}

		commentAPI := api.Group("/comments")
		{
			commentAPI.POST("/", commentHandler.Create)
			commentAPI.GET("/", commentHandler.GetBySource)
			commentAPI.DELETE("/:id", commentHandler.Delete)
		}

		assignmentAPI := api.Group("/assignments")
		{
			assignmentAPI.POST("/categories", assignmentHandler.CreateCategory)
			assignmentAPI.POST("/", assignmentHandler.CreateAssignment)
			assignmentAPI.GET("/class/:classId", assignmentHandler.GetByClass)
			assignmentAPI.POST("/submit", assignmentHandler.Submit)
			assignmentAPI.POST("/assess", assignmentHandler.Assess)
		}

		logAPI := api.Group("/logs")
		{
			logAPI.GET("/school/:schoolId", logHandler.GetBySchool)
		}
	}

	//run server
	r.Run(":8080")
}
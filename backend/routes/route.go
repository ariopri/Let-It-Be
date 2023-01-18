package routes

import (
	"net/http"
	"time"

	"github.com/ariopri/Let-It-Be/tree/main/backend/controllers"
	kursusRepo "github.com/ariopri/Let-It-Be/tree/main/backend/kursus/repository"
	kursusUseCase "github.com/ariopri/Let-It-Be/tree/main/backend/kursus/usecase"
	"github.com/ariopri/Let-It-Be/tree/main/backend/middlewares"
	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	userRepo "github.com/ariopri/Let-It-Be/tree/main/backend/users/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	kursusHandler "github.com/ariopri/Let-It-Be/tree/main/backend/kursus/handler"
	userHandler "github.com/ariopri/Let-It-Be/tree/main/backend/users/delivery/http"

	authHandler "github.com/ariopri/Let-It-Be/tree/main/backend/controllers/auth/delivery/http"
	authRepo "github.com/ariopri/Let-It-Be/tree/main/backend/controllers/auth/repository"
)

func SetupRoutes() *gin.Engine {
	models.ConnectDatabase()
	r := gin.Default()
	//user cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/informasi", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.InformasiData)
	})
	r.GET("/modul", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Kelas)
	})
	r.GET("/detail", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.DetailModul)
	})
	r.GET("/faq", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.FAQList)
	})
	public := r.Group("/")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/admin/register", controllers.RegisterAdmin)
	public.POST("/reset/password", controllers.ResetPassword)
	//=============================Middlewares for Client======================================================
	client := r.Group("/client")
	client.Use(middlewares.JwtAuthMiddleware())
	client.GET("/user", controllers.CurrentUser)
	client.PUT("/update/profile", controllers.UpdateProfile)

	//=============================Middlewares for Admin======================================================
	admin := r.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddlewareAdmin())
	admin.GET("/user", controllers.CurrentUser)
	admin.PUT("/update/profile", controllers.UpdateProfile)

	//=============================Middlewares for Auth======================================================
	authRepo := authRepo.NewAuthRepo(models.DB)
	authHandler.NewAuthHandler(r, authRepo)
	//=============================Middlewares for Kursus======================================================
	ctxTimeout := time.Duration(2) * time.Second
	userRepo := userRepo.NewUserRepo(models.DB)
	userHandler.NewSiswaHandler(r, userRepo)

	kursusRepo := kursusRepo.NewKursusRepo(models.DB)
	kursusUseCase := kursusUseCase.NewKursusUseCase(kursusRepo, userRepo, ctxTimeout)
	kursusHandler.NewKursusHandler(r, kursusUseCase)

	return r
}

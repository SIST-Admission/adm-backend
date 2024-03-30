package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SIST-Admission/adm-backend/src/controllers"
	"github.com/SIST-Admission/adm-backend/src/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configFilePath *string
var env string

func start() {
	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()
	loadConfig(configFilePath)

	// Set Server Mode
	if env == "prod." {
		// gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(middlewares.CORSMiddleware())
	loadRoutes(engine, viper.GetString(env+"server.basePath"))
	startServer(engine, viper.GetString(env+"server.port"))
}

func startServer(engine *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	go func() {
		logrus.Info("Starting server on port ", port)
		logrus.Info("Server host ", viper.GetString(env+"server.host"))
		logrus.Info("Client host ", viper.GetString(env+"client.host"))
		logrus.Info("Server base path ", viper.GetString(env+"basePath"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutting down server...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Default().Println("Server exiting")
}

func loadConfig(configFilePath *string) {
	viper.SetConfigName("app")
	viper.AddConfigPath("..")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	env = viper.GetString("env") + "."
	logrus.Info("Environment: ", env)
}

func loadRoutes(engine *gin.Engine, basePath string) {

	// Controllers
	applicationsController := controllers.ApplicationsController{}
	userController := controllers.UserController{}
	paymentsController := controllers.PaymentsController{}
	batchesController := controllers.BatchesController{}
	meritListsController := controllers.MeritListsController{}

	// Application Routes "/{basePath}"
	app := engine.Group(basePath)
	{
		app.GET("/", middlewares.Auth, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "User is authenticated",
			})
		})
		app.GET("/auth", middlewares.Auth, userController.LoggedInUser)

		// User Routes "/{basePath}/users"
		users := app.Group("/users")
		{
			users.POST("/", userController.RegisterUser)
			users.POST("/login", userController.LoginUser)
			users.GET("/logout", userController.LogoutUser)
		}

		// User Routes "/{basePath}/applications"
		applications := app.Group("/applications")
		applications.Use(middlewares.Auth)
		{
			applications.GET("/stats", applicationsController.GetApplicationStats)
			applications.POST("getAllApplications", applicationsController.GetAppApplications)
			applications.GET("/:appId", applicationsController.GetApplication)
			applications.GET("/getApplicationByUser", applicationsController.GetApplicationByUser)
			applications.POST("/start", applicationsController.StartApplication)
			applications.POST("/basicDetails", applicationsController.SaveBasicDetails)
			applications.POST("/academicDetails", applicationsController.SaveAcademicDetails)
			applications.POST("/submitApplication", applicationsController.SubmitApplication)
			applications.PUT("/updateDocumentStatus", middlewares.AdminAuth, applicationsController.UpdateDocumentStatus)
			applications.PUT("/updateApplicationStatus", middlewares.AdminAuth, applicationsController.UpdateApplicationStatus)
		}

		// Documents Routes "/{basePath}/documents"
		meritList := app.Group("/meritLists")
		meritList.Use(middlewares.Auth, middlewares.AdminAuth)
		{
			meritList.POST("/createList", meritListsController.CreateMeritList)
			meritList.POST("/addStudent", meritListsController.AddStudents)
			meritList.POST("/getAllMeritLists", meritListsController.GetAllMeritLists)
			meritList.POST("/getUnListedCandidates", meritListsController.GetUnListedCandidates)
			meritList.POST("/getListedCandidates", meritListsController.GetListedCandidates)
		}

		batches := app.Group("/batches")
		batches.Use(middlewares.Auth)
		{
			batches.GET("/getAllBatches", batchesController.GetBatches)
		}

		// Payments Routes "/{basePath}/payments"
		payments := app.Group("/payments")
		payments.Use(middlewares.Auth)
		{
			payments.POST("/getOrder", paymentsController.GetOrder)
			payments.GET("/transactions", paymentsController.GetTransactions)
			payments.POST("/getAdmissionOrder", paymentsController.GetAdmissionOrder)
		}

		hooks := app.Group("/hooks")
		{
			hooks.POST("/verifyPayment", paymentsController.VerifyPayment)
		}

	}
}

func main() {
	start()
}

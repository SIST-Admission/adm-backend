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
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var configFilePath *string

func start() {
	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()
	loadConfig(configFilePath)

	// Set Server Mode
	if viper.GetString("env") == "prod" {
		// gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	loadRoutes(engine, viper.GetString("server.basePath"))
	startServer(engine, viper.GetString("server.port"))
}

func startServer(engine *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	go func() {
		log.Default().Println("Starting server on port", port)
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
		log.Default().Fatalf("Server forced to shutdown: %v", err)
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
}

func loadRoutes(engine *gin.Engine, basePath string) {

	// Controllers
	applicationController := controllers.ApplicationController{}
	userController := controllers.UserController{}

	// Application Routes "/{basePath}"
	application := engine.Group(basePath)
	{
		application.GET("/ping", applicationController.Ping)

		// User Routes "/{basePath}/users"
		users := application.Group("/users")
		{
			users.POST("/", userController.RegisterUser)
		}

	}
}

func main() {
	start()
}

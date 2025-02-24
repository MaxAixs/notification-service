package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"notification-service/cmd/server"
	"notification-service/internal/config"
	"notification-service/notification/handler"
	"notification-service/notification/repository"
	"notification-service/notification/service"
	"notification-service/pkg/analyticService/client"
	"notification-service/pkg/analyticService/worker"
	"notification-service/pkg/database"
	"notification-service/pkg/mailgun"
	"os"
	"os/signal"
	"time"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := config.InitCfg(); err != nil {
		logrus.Fatalf("init config failed: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file failed: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("server.port")
		logrus.Infof("SERVER_PORT not set, default server port: %s", port)
	}

	db, err := database.NewPostgresDB(database.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	mgDomain := os.Getenv("MAILGUN_DOMAIN")
	mgAPIToken := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun(mgDomain, mgAPIToken)

	repo := repository.NewRepository(db)
	services := service.NewNotificationService(repo, mg)
	handlers := handler.NewHandler(services)

	srv := &server.Server{}
	srvCfg := server.ConfigServer{
		Port:              config.GetPort(),
		ReadHeaderTimeout: viper.GetDuration("server.read_header_timeout"),
		WriteTimeout:      viper.GetDuration("server.write_timeout"),
		IdleTimeout:       viper.GetDuration("server.idle_timeout"),
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := srv.RunServer(srvCfg, handlers.MapRoutes()); err != nil {
			logrus.Fatalf("failed to start notification server: %v", err)
		}
	}()

	logrus.Println("notification server started")

	gRPCClient, err := client.NewAnalyticsClient(viper.GetString("analytics.grpc_port"))
	if err != nil {
		logrus.Fatalf("failed to create analytics client: %v", err)
	}
	defer gRPCClient.CloseConnection()

	analyticWorker := worker.NewAnalyticWorker(services, gRPCClient)

	go func() {
		if err := analyticWorker.Start(ctx); err != nil {
			logrus.Fatalf("failed to start analytics worker: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	cancel()

	shutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutDownCancel()

	if err := srv.ShutDown(shutDownCtx); err != nil {
		logrus.Fatalf("failed to shutdown notification server: %v", err)
	}

	logrus.Println("notification server shutdown")
}

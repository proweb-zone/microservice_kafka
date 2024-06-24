package main

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"microservice/kafka/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// type msgTopic struct {
// 	ID     string  `json:"id"`
// 	MSG  string  `json:"msg"`
// }

func main(){

cfg := config.MustLoad()

log := setupLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// router := gin.Default()
	// router.POST("/topic/:topic_name", addMsgToTopic)
	// router.Run("localhost:3003")
}

// func addMsgToTopic(c *gin.Context){
// 	topicName := c.Param("topic_name")

// 	var newMsgTopic msgTopic

//     if err := c.BindJSON(&newMsgTopic); err != nil {
//         return
//     }

// 	repo.AddMsgToTopic(topicName, newMsgTopic)

// 	c.IndentedJSON(http.StatusOK, "write to topic")
// }

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

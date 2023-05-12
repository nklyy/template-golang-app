package main

import (
	"errors"
	"log"
	"syscall"
	"template-golang-app/config"
	"template-golang-app/pkg/logger"
	"template-golang-app/services/health"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init logger
	newLogger, err := logger.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("can't create logger: %v", err)
	}

	zapLogger, err := newLogger.SetupZapLogger()
	if err != nil {
		log.Fatalf("can't setup zap logger: %v", err)
	}
	defer func(zapLogger *zap.SugaredLogger) {
		err := zapLogger.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			log.Fatalf("can't setup zap logger: %v", err)
		}
	}(zapLogger)

	// // Connect to database
	// db, ctx, cancel, err := mongodb.NewConnection(cfg)
	// if err != nil {
	// 	zapLogger.Fatalf("failed to connect to mongodb: %s", err)
	// }
	// defer mongodb.Close(db, ctx, cancel)

	// // Ping db
	// err = mongodb.Ping(db, ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zapLogger.Info("DB connected successfully")

	// Handlers
	healthHandler := health.NewHandler()

	// Create config variable
	config := fiber.Config{
		ServerHeader: "Template Golang App", // add custom server header
	}

	// Create fiber app
	app := fiber.New(config)

	// Set-up Route
	app.Route("/api/v1", func(router fiber.Router) {
		healthHandler.SetupRoutes(router)
	})

	// Handle 404 page
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{"error": "page not found"})
	})

	// Start App
	zapLogger.Infof("Starting HTTP server on port: %v", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project/config"
	"project/internal/handler"
	"project/internal/middleware"
	"project/internal/services"
	"project/pkg/logger"
)

func main() {
	// 1. Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Инициализация логгера
	log := logger.New()

	// 3. Инициализация сервисов
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Инициализация сервиса с передачей API ключа
	predictionService := services.NewPredictionService(httpClient, cfg)
	apiHandler := handler.NewAPIHandler(predictionService, log)

	// 4. Настройка роутера
	mux := http.NewServeMux()

	// Регистрация обработчиков
	mux.HandleFunc("/predict/hba1c", apiHandler.HandleHBA1C)
	mux.HandleFunc("/predict/tg", apiHandler.HandleTG)
	mux.HandleFunc("/predict/hdl", apiHandler.HandleHDL)
	mux.HandleFunc("/predict/ldl", apiHandler.HandleLDL)
	mux.HandleFunc("/predict/ferr", apiHandler.HandleFERR)
	mux.HandleFunc("/predict/ldll", apiHandler.HandleLDLL)

	// 5. Применение middleware
	authMiddleware := middleware.AuthMiddleware(cfg, log)
	handlerWithAuth := authMiddleware(mux)

	// 6. Настройка сервера
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      handlerWithAuth,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 7. Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting server on :%s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server error: %v", err)
		}
	}()

	<-done
	log.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server shutdown error: %v", err)
	} else {
		log.Info("Server stopped gracefully")
	}
}

package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// TODO: go-fiber + https://docs.gofiber.io/api/middleware/adaptor/

func Run() {

	// *****
	// Реализация Clean Architecture
	// *****
	// myRepository := repository.New()
	// myService := service.New(myRepository)
	// myController := controller.New(myService)
	// myController.Serve()

	// *****
	// Реализация Graceful Shutdown
	// *****

	// Создаем контекст для отслеживания сигнала завершения работы
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем сервер
	server := &http.Server{
		Addr:    ":8888",
		Handler: nil, // Ваш обработчик запросов
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	log.Println("Сервер запущен")

	// Ожидаем сигнал завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Получен сигнал завершения работы")

	// Устанавливаем таймаут для graceful shutdown
	timeout := 5 * time.Second
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	// Останавливаем сервер с использованием контекста
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("Сервер остановлен")
}

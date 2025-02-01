//Практика работы с плавным завершением сервиса

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	endpoint         = "/ping"
	clientAnswer     = "pong"
	hostAdress       = "localhost:8080"
	gShutdownTimeout = 10 // second
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(endpoint, handlePing)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	server := &http.Server{
		Addr:    hostAdress,
		Handler: mux,
	}

	go startServer(server)
	log.Println("Сервер слушает запросы")

	//главная горутина здесь остановилась
	//код в main продолжит выполнение, если получит сигнал от ОС
	sig := <-sigs

	err := gracefulShutdown(server, sig)
	if err != nil {
		log.Fatal("Сервис завершён с ошибкой: ", err)
	}

	log.Println("Сервис завершён безопасно")
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен запрос: %s", r.URL)

	time.Sleep(3000 * time.Millisecond)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(clientAnswer))
	log.Println("Запрос обработан")
}

func startServer(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("ошибка работы сервера:", err)
	}
}

func gracefulShutdown(srv *http.Server, sig os.Signal) error {
	log.Println("от ОС поступил сигнал завершения сервиса:", sig)

	maxDelay := time.Duration(gShutdownTimeout * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), maxDelay)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("работа сервера окончена таймаутом graceful-shutdown: %s", err)
	}

	return nil
}

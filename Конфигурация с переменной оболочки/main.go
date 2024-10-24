//Тренируюсь конфиругировать приложение на переменных окружения

package main

import (
	"log"
	"net/http"
	"os"
)

type app struct {
	port string
}

func mustNew() app {
	return app{
		port: port(),
	}
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Отсутствует переменная окружения PORT")
	}
	log.Println("Загрузили переменную оболочки PORT: ", port)
	return port
}

func main() {
	a := mustNew()
	err := a.run()
	if err != nil {
		log.Fatal(err)
	}
}

func (a app) run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет, Golang-инженеры!"))
	})
	err := http.ListenAndServe(a.port, nil)
	if err != nil {
		return err
	}
	return nil
}

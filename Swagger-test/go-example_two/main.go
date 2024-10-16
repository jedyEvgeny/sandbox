//Тестирую работу со Swagger
//Описание см. здесь: https://dzen.ru/a/ZwTRpba9sACJjBwj
package main

import (
	h "example_swagger_test/handlers"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title						Я, Golang-инженер
// @version					1.0
// @description				Тестируем бэкенд с фронтендом по CRUD-запросам
// @contact.name				Евгений
// @contact.url				https://github.com/jedyEvgeny
// @contact.email				KEF1991@yandex.ru
// @license.name				MIT
// @license.url				http://opensource.org/licenses/MIT
// @host						localhost:8080
// @BasePath					/home
// @accept						json
// @produce					json text/plain
// @schemes					http https
// @externalDocs.description	Резерв для дополнительного описания API
// @externalDocs.url			https://t.me/+ZGac_D1V4wFjYzRi
// @x-name						{"environment": "production", "version": "1.0.0", "team": "backend"}
// @tag.name					items
// @tag.description			Операции с товарами
// @tag.docs.url				https://t.me/EvKly
// @tag.docs.description		Консультация по работе с товарами
func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))

	http.HandleFunc("/home/create_item", h.HandlerCreateItem)
	http.HandleFunc("/home/update_item/", h.HandlerUpdateItem)
	http.HandleFunc("/home/delete_item/", h.HandleDeleteItem)
	http.HandleFunc("/home/item/", h.HandleItem)

	http.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "./docs/swagger.json")
	})
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Запустили сервер")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

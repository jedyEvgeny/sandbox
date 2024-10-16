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

// 2. Обновить товар (PUT)
// Запрос:

// Метод: PUT
// URL: /home/update_item/СтарыйТовар (где СтарыйТовар - название товара, который обновляется)
// Тело запроса:
// json

// Copy
// {
//     "product": "НовыйТовар",
//     "price": 150
// }
// Ответ:

// Статус: 200 OK (если успешно)
// Тело ответа:
// json

// Copy
// {
//     "product": "НовыйТовар",
//     "price": 150
// }
// 3. Обновить цену товара (PATCH)
// Запрос:

// Метод: PATCH
// URL: /home/update_item/НазваниеТовара (где НазваниеТовара - название товара, чья цена обновляется)
// Тело запроса:
// json

// Copy
// {
//     "price": 200
// }
// Ответ:

// Статус: 200 OK (если успешно)
// Тело ответа:
// json

// Copy
// {
//     "product": "НазваниеТовара",
//     "price": 200
// }
// 4. Удалить товар (DELETE)
// Запрос:

// Метод: DELETE
// URL: /home/delete_item/НазваниеТовара (где НазваниеТовара - имя товара, который нужно удалить)
// Тело запроса: отсутствует.
// Ответ:

// Статус: 204 No Content (если успешно удалён)
// Тело ответа: отсутствует.

// 5. Показать товар по наименованию (GET)
// Запрос:

// Метод: GET
// URL: /home/item/НазваниеТовара (где НазваниеТовара - имя товара, который нужно показать)
// Тело запроса: отсутствует.
// Ответ:

// Статус: 200 OK (если товар найден)
// Тело ответа:
// json

// Copy
// {
//     "product": "НазваниеТовара",
//     "price": 200
// }
// Или:
// Статус: 404 Not Found (если товар не найден)
// Тело ответа:
// json

// Copy
// {
//     "error": "Товар не найден"
// }

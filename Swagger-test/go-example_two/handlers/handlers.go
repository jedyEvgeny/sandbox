package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	dir  = "./data"
	perm = 0644
)

//	@Summary		Создать новый товар
//	@Description	Создает новый товар с указанным именем и ценой.
//	@Description	Наименование и цена передаются в теле в json-объекте.
//	@Description	Если товар не удается создать, возвращает ошибку.
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			item	body		Item		true	"Создаем новый товар"
//	@Success		201		{object}	Resourse	"Товар успешно создан"
//	@Failure		400		{object}	nil			"Ошибка валидации данных"
//	@Failure		405		{object}	nil			"Метод не разрешен"
//	@Failure		500		{object}	nil			"Ошибка сервера"
//	@Router			/home/create_item [post]
func HandlerCreateItem(w http.ResponseWriter, r *http.Request) {
	log.Printf(msgNewRequest, r.URL, r.Method, r.Proto)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if r.Method != http.MethodPost {
		log.Printf(errMethod, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := fmt.Sprintf("получен метод %s, ожидался %s", r.Method, http.MethodPost)
		w.Write([]byte(msg))
		return
	}
	i := Item{}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		log.Printf(errDecodeJson, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fName, err := saveItem(i)
	if err != nil {
		log.Printf(errSaveItem, err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("не смогли добавить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fName))
}

func HandlerUpdateItem(w http.ResponseWriter, r *http.Request) {
	log.Printf(msgNewRequest, r.URL, r.Method, r.Proto)
	switch r.Method {
	case http.MethodPatch:
		patchUpdates(w, r)
	case http.MethodPut:
		putUpdates(w, r)
	default:
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		log.Printf(errMethod, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := fmt.Sprintf("получен метод %s, ожидался %s", r.Method, http.MethodPatch)
		w.Write([]byte(msg))
	}
}

//	@Summary		Удалить товар
//	@Description	Удаляет товар по наименованию
//	@Description	Наименование товара передаётся в эндпоинте
//	@Description	Если товар не удается удалить, возвращает ошибку.
//	@Tags			items
//	@Produce		text/plain
//	@Success		204	{object}	nil		"Товар успешно удалён"
//	@Failure		404	{object}	string	"Нечего удалять: товар не найден"
//	@Failure		405	{object}	string	"Метод не разрешен"
//	@Failure		500	{object}	string	"Ошибка сервера"
//	@Param			id	path		string	true	"Удаляем товар"
//	@Router			/home/delete_item/{id} [delete]
func HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	log.Printf(msgNewRequest, r.URL, r.Method, r.Proto)
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	if r.Method != http.MethodDelete {
		log.Printf(errMethod, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := fmt.Sprintf("получен метод %s, ожидался %s", r.Method, http.MethodDelete)
		w.Write([]byte(msg))
		return
	}
	status, err := deleteItem(r)
	if err != nil && status == 500 {
		log.Printf(errDeleteItem, err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("не смогли удалить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	if err != nil && status == 404 {
		log.Printf(errDeleteItem, err)
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("не смогли обнаружить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//	@Summary		Получаем информацию о товаре
//	@Description	Получает наименование и цену товара.
//	@Description	Наименование товара передаётся как часть эндпоинта.
//	@Description	Если товар не удается найти, возвращает ошибку.
//	@Tags			items
//	@Produce		json
//	@Success		200	{object}	Item	"Информация о товаре получена"
//	@Failure		404	{object}	nil		"Товар не найден"
//	@Failure		405	{object}	nil		"Метод не разрешен"
//	@Failure		500	{object}	nil		"Ошибка сервера"
//	@Param			id	path		string	true	"Обновляем существующий товар"
//	@Router			/home/item/{id} [get]
func HandleItem(w http.ResponseWriter, r *http.Request) {
	log.Printf(msgNewRequest, r.URL, r.Method, r.Proto)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if r.Method != http.MethodGet {
		log.Printf(errMethod, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := fmt.Sprintf("получен метод %s, ожидался %s", r.Method, http.MethodGet)
		w.Write([]byte(msg))
		return
	}
	jsonData, status, err := item(r)
	if err != nil && status == 500 {
		log.Printf(errDeleteItem, err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("не смогли удалить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	if err != nil && status == 404 {
		log.Printf(errDeleteItem, err)
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("не смогли обнаружить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

//	@Summary		Обновляем стоимость товара
//	@Description	Обновляет стоимость товара.
//	@Description	Наименование товара передаётся как часть эндпоинта.
//	@Description	Стоимость передаётся в теле в json-объекта.
//	@Description	Если товар не удается обновить, возвращает ошибку.
//	@Tags			items
//	@Accept			json
//	@Produce		text/plain
//	@Param			item	body		ItemPrice	true	"Создаем новый товар"
//	@Success		204		{object}	nil			"Цена обновлена"
//	@Failure		400		{object}	nil			"Ошибка валидации данных"
//	@Failure		405		{object}	nil			"Метод не разрешен"
//	@Failure		500		{object}	nil			"Ошибка сервера"
//	@Param			id		path		string		true	"Обновляем существующий товар"
//	@Router			/home/create_item/{id} [patch]
func patchUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	i := Item{}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		log.Printf(errDecodeJson, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = updateItem(i, r)
	if err != nil {
		log.Printf(errSaveItem, err)
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("не смогли обновить цену: %v", err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//	@Summary		Обновляем товар
//	@Description	Обновляет наименование и стоимость товара.
//	@Description	Текущее наименование товара передаётся как часть эндпоинта.
//	@Description	Новое наименование товара и стоимость передаётся в теле запроса json-объектом.
//	@Description	Если товар не удается обновить, возвращает ошибку.
//	@Tags			items
//	@Accept			json
//	@Produce		text/plain
//	@Param			item	body		Item	true	"Создаем новый товар"
//	@Success		204		{object}	nil		"Наименование и цена товара обновлены"
//	@Failure		400		{object}	nil		"Ошибка валидации данных"
//	@Failure		405		{object}	nil		"Метод не разрешен"
//	@Failure		500		{object}	nil		"Ошибка сервера"
//	@Param			id		path		string	true	"Обновляем существующий товар"
//	@Router			/home/create_item/{id} [put]
func putUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	i := Item{}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		log.Printf(errDecodeJson, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = totalUpdateItem(i, r)
	if err != nil {
		log.Printf(errSaveItem, err)
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("не смогли обновить товар: %v", err)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func saveItem(i Item) (string, error) {
	if ok := checkTegs(i); !ok {
		return "", fmt.Errorf(errAllFields, i)
	}
	fPath := filepath.Join(dir, i.Product+".txt")
	if _, err := os.Stat(fPath); err == nil {
		return "", fmt.Errorf(errAlreadyExists, i.Product, err)
	}
	if err := os.MkdirAll(dir, perm); err != nil {
		return "", fmt.Errorf(errCreateDir, err)
	}
	priceStr := priceToString(i.Price)
	err := os.WriteFile(fPath, []byte(priceStr), perm)
	if err != nil {
		return "", fmt.Errorf(errSaveRes, err)
	}
	log.Printf(msgSaveSucsecc, fPath)
	return fPath, nil
}

func updateItem(i Item, r *http.Request) error {
	product := strings.TrimPrefix(r.URL.Path, "/home/update_item/")
	i.Product = product
	if ok := checkTegs(i); !ok {
		return fmt.Errorf(errAllFields, i)
	}
	fPath := filepath.Join(dir, i.Product+".txt")
	if _, err := os.Stat(fPath); err != nil {
		return fmt.Errorf(errNotExists, i.Product, err)
	}
	priceStr := priceToString(i.Price)
	err := os.WriteFile(fPath, []byte(priceStr), perm)
	if err != nil {
		return fmt.Errorf(errSaveRes, err)
	}
	log.Printf("цена товара %s обновлена на %.2f\n", i.Product, i.Price)
	return nil
}

func totalUpdateItem(i Item, r *http.Request) error {
	oldProduct := strings.TrimPrefix(r.URL.Path, "/home/update_item/")
	if ok := checkTegs(i); !ok {
		return fmt.Errorf(errAllFields, i)
	}
	fPathOld := filepath.Join(dir, oldProduct+".txt")
	if _, err := os.Stat(fPathOld); err != nil {
		return fmt.Errorf(errNotExists, i.Product, err)
	}
	priceStr := priceToString(i.Price)
	err := os.WriteFile(fPathOld, []byte(priceStr), perm)
	if err != nil {
		return fmt.Errorf(errSaveRes, err)
	}
	fPathNew := filepath.Join(dir, i.Product+".txt")
	err = os.Rename(fPathOld, fPathNew)
	if err != nil {
		return fmt.Errorf(errChangeName, err)
	}
	log.Printf("старый товар %s обновлён новым товаром %s\n", oldProduct, i.Product)
	return nil
}

func deleteItem(r *http.Request) (int, error) {
	product := strings.TrimPrefix(r.URL.Path, "/home/delete_item/")
	fPath := filepath.Join(dir, product+".txt")
	if _, err := os.Stat(fPath); err != nil {
		return http.StatusNotFound, fmt.Errorf(errNotExists, product, err)
	}
	err := os.Remove(fPath)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf(errSaveRes, err)
	}
	log.Printf("товар %s удалён\n", product)
	return http.StatusOK, nil
}

func item(r *http.Request) ([]byte, int, error) {
	product := strings.TrimPrefix(r.URL.Path, "/home/item/")
	fPath := filepath.Join(dir, product+".txt")
	if _, err := os.Stat(fPath); err != nil {
		return nil, http.StatusNotFound, fmt.Errorf(errNotExists, product, err)
	}
	price, err := price(fPath)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не смогли получить цену: %w", err)
	}
	i := Item{
		Product: product,
		Price:   price,
	}
	dataJson, err := json.Marshal(i)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(errSaveRes, err)
	}
	log.Printf("товар %s подготовлен к отправке клиенту\n", product)
	return dataJson, http.StatusOK, nil
}

func checkTegs(i Item) bool {
	ok := true
	if i.Price <= 0 {
		log.Println("ошибка данных: цена должна быть больше нуля")
		ok = false
	}
	if i.Product == "" {
		log.Println("ошибка данных: отсутствует наименование продукта")
		ok = false
	}
	return ok
}

func priceToString(price float64) string {
	return strconv.FormatFloat(price, 'f', 2, 64)
}

func price(fPath string) (float64, error) {
	content, err := os.ReadFile(fPath)
	if err != nil {
		return 0.0, fmt.Errorf(errReadItem, err)
	}
	price, err := strconv.ParseFloat(string(content), 64)
	if err != nil {
		return 0.0, fmt.Errorf("не смогли преобразовать цену в float64: %w", err)
	}
	return price, nil
}

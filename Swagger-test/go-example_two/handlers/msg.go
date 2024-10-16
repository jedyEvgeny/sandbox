package handlers

const (
	errDecodeJson    = "не смогли декодировать json: %v\n"
	errSaveItem      = "не смогли сохранить информацию: %v\n"
	errUpdateItem    = "не смогли обновить цену: %v\n"
	errCreateDir     = "не удалось создать каталог: %w"
	errSaveRes       = "не удалось сохранить инфо: %w"
	errMethod        = "несоответствие метода: %s"
	errAllFields     = "заполнены не все поля: %v"
	errAlreadyExists = "товар %s уже существует: %w"
	errNotExists     = "товар %s ещё не создан: %w"
	errChangeName    = "не смогли переименовать товар: %w"
	errDeleteItem    = "не смогли удалить товар: %v"
	errReadItem      = "не удалось прочитать файл: %w"
)

const (
	msgNewRequest  = "Получен запрос:\n\t- Url: %s\n\t- Метод: %s\n\t- По протоколу: %s\n"
	msgSaveSucsecc = "Успешное сохранение по пути: \n\t- %s"
)

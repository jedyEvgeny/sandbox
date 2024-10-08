-- 1. Создаем новую таблицу с нужной схемой
CREATE TABLE IF NOT EXISTS new_hosts (
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    host TEXT,
    country TEXT, -- Новое поле
    region TEXT, -- Новое поле
    City TEXT, -- Новое поле
    ISP TEXT, -- Новое поле
    Timezone TEXT, -- Новое поле
    timeAdd DATETIME DEFAULT CURRENT_TIMESTAMP -- Указываем, что временная метка ставится автоматически
);

-- 2. Копируем данные из старой таблицы в новую
INSERT INTO new_hosts (id, ip, host, timeAdd)
SELECT id, ip, host, timeAdd FROM hosts;

-- 3. Удаляем старую таблицу
DROP TABLE hosts;

-- 4. Переименовываем новую таблицу
ALTER TABLE new_hosts RENAME TO hosts;

-- 5. Создаем индекс
CREATE INDEX IF NOT EXISTS idx_items_name ON hosts(host);


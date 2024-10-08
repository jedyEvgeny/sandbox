-- 1. Создаем новую таблицу со схемой как у предыдущей
CREATE TABLE IF NOT EXISTS new_hosts (
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    host TEXT,
    timeAdd DATETIME
);

-- 2. Копируем данные из старой таблицы hosts
INSERT INTO new_hosts (id, ip, host, timeAdd)
SELECT id, ip, host, timeAdd FROM hosts;

-- 3. Удаляем старую таблицу
DROP TABLE IF EXISTS hosts;

-- 4. Переименовываем новую таблицу
ALTER TABLE new_hosts RENAME TO hosts;

-- 5. Создаем индекс, который был в старой таблице
CREATE INDEX IF NOT EXISTS idx_items_name ON hosts(host);


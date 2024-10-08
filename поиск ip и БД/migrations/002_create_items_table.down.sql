-- 1. Создаем старую таблицу с первоначальной схемой
CREATE TABLE IF NOT EXISTS new_hosts (
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    host FLOAT, -- Возвращаем тип данных на FLOAT
    timeAdd DATETIME
);

-- 2. Копируем данные обратно из таблицы hosts
INSERT INTO new_hosts (id, ip, host, timeAdd)
SELECT id, ip, host, timeAdd FROM hosts;

-- 3. Удаляем новую таблицу
DROP TABLE IF EXISTS hosts;

-- 4. Переименовываем новую таблицу
ALTER TABLE new_hosts RENAME TO hosts;

-- 5. Создаем индекс, если он был в старой таблице
CREATE INDEX IF NOT EXISTS idx_items_name ON hosts(host);


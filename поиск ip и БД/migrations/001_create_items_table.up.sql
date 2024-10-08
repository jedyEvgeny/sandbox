CREATE TABLE IF NOT EXISTS hosts (
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    host FLOAT,
    timeAdd DATETIME
);
CREATE INDEX IF NOT EXISTS 
idx_items_name ON hosts(host);


//Цель программы - научиться пользоваться миграциями
//Через флаг -ip или -domain вводятся параметры для поиска инфо
//Программа ищет по ip гео-данные, сохраняет их в БД 
//Через СУБД SQLite и выводить последние 10 записей БД в терминал

package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type GeoResponse struct {
	IP         string `json:"query"`
	Host       string
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	ISP        string `json:"isp"`
	Timezone   string `json:"timezone"`
	Provider   string `json:"as"`
}

const (
	errEmptyFlags    = "не введён ip-адрес или домен"
	errNotIpOrDomain = "введённые данные не распознаны как ip-адрес или домен: %w"
	errStmt          = "не смогли подготовить sql-запрос: %w"
	errExec          = "не смогли выполнить sql-запрос: %w"
	errRes           = "не смогли получить результат выполнения sql-запроса: %w"
	errResAffected   = "не выполнены изменения в БД: %w"

	geoAPI = "ip-api.com"
)

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatal("не удалось создать БД")
	}
	defer func() { _ = db.Close() }()

	possibleIP, possibleAddr := mustParseData()

	ip, host, err := IPAndHost(possibleIP, possibleAddr)
	if err != nil {
		log.Fatal(err)
	}

	g, err := geoInfo(ip, host)
	if err != nil {
		log.Fatal(err)
	}
	g.printGeoInfo()

	err = g.insertInfo(db)
	if err != nil {
		log.Fatal(err)
	}

	err = printLastTenEntries(db)
	if err != nil {
		log.Fatal(err)
	}
}

func geoInfo(ip, host string) (GeoResponse, error) {
	u := url.URL{
		Scheme: "http",
		Host:   geoAPI,
		Path:   path.Join("json", ip),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return GeoResponse{}, fmt.Errorf("не смогли получить ответ API: %w", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return GeoResponse{}, fmt.Errorf("не удалось связаться с сервером: %w", err)
	}
	defer func() { resp.Body.Close() }()
	geo := GeoResponse{}
	err = json.NewDecoder(resp.Body).Decode(&geo)
	if err != nil {
		return GeoResponse{}, err
	}
	geo.Host = host
	return geo, nil
}

func (g GeoResponse) printGeoInfo() {
	s := fmt.Sprintf(`
	IP: %s
	Хост: %s
	Страна: %s
	Регион: %s
	Город: %s
	Компания: %s
	Провайдер: %s
	Часовой пояс: %s
`, g.IP, g.Host, g.Country, g.RegionName,
		g.City, g.ISP, g.Provider, g.Timezone)
	fmt.Print(s)
}

func mustParseData() (string, string) {
	if len(os.Args) != 3 {
		log.Fatal(errEmptyFlags)
	}
	ip := flag.String("ip", "", "ip-адрес сайта")
	addr := flag.String("domain", "", "домен сайта")

	flag.Parse()

	if *ip == "" && *addr == "" {
		log.Fatal(errEmptyFlags)
	}
	return *ip, *addr
}

func IPAndHost(possibleIP, possibleDomain string) (string, string, error) {
	var ip []net.IP
	var addr []string
	var err error
	if possibleIP != "" {
		addr, err = net.LookupAddr(possibleIP)
		if err != nil {
			return "", "", fmt.Errorf("не распознан IP %s: %w", possibleIP, err)
		}
		ip = append(ip, net.ParseIP(possibleIP))
	}

	if possibleDomain != "" {
		ip, err = net.LookupIP(possibleDomain)
		if err != nil {
			return "", "", fmt.Errorf("не распознан домен %s: %w", possibleDomain, err)
		}
		addr = append(addr, possibleDomain)
	}
	fmt.Printf("IP-адрес: %s\nХост: %s\n", ip, addr)

	return ip[0].String(), addr[0], nil
}

func initDatabase() (*sql.DB, error) {
	dbName := "internet_resuorses.db?_timeout=50"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть БД: %w", err)
	}

	start := time.Now()
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось выполнить пинг: %w", err)
	}
	end := time.Now()
	log.Println("Пинг выполнен за:", end.Sub(start))

	err = reserveDatabase()
	if err != nil {
		return nil, fmt.Errorf("ошибка создания резервной копии БД %w", err)
	}

	migrationPath := "file://migrations"
	dbFullName := "sqlite3://" + dbName

	m, err := migrate.New(migrationPath, dbFullName)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать мигратор: %w", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("не удалось применить миграции: %w", err)
	}
	if err == migrate.ErrNoChange {
		log.Println("Схема БД в актуальном состоянии")
	}
	if err != migrate.ErrNoChange {
		log.Println("миграции применены успешно!")
	}
	cancelMigrations(m)
	return db, nil
}

func cancelMigrations(m *migrate.Migrate) {
	fmt.Print("Откатить последнюю миграцию? (y/n): ")
	var ans string
	if _, _ = fmt.Scan(&ans); ans == "y" {
		err := m.Steps(-1)
		if err != nil {
			log.Fatalf("Ошибка при откате миграции: %v", err)
		}
		log.Println("Успех в откате последней миграции!")
	}
}

func reserveDatabase() error {
	var response string
	fmt.Print("Хотите сделать резервную копию БД? (y/n): ")
	_, _ = fmt.Scan(&response)
	if response != "y" {
		return nil
	}
	dbName := "internet_resuorses.db"
	backupDir := "backup_DB"
	err := os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("не удалось создать директорию %s: %w", backupDir, err)
	}
	timeTemplate := "20060102_150405"
	backupFileName := time.Now().Format(timeTemplate) + "_backup_DB.db"
	backupFilePath := filepath.Join(backupDir, backupFileName)
	sourceFile, err := os.Open(dbName)
	if err != nil {
		return fmt.Errorf("не удалось открыть БД для резервного копирования: %w", err)
	}
	defer func() { _ = sourceFile.Close() }()
	destFile, err := os.Create(backupFilePath)
	if err != nil {
		return fmt.Errorf("не удалось создать резервную копию БД: %w", err)
	}
	defer func() { _ = destFile.Close() }()
	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("ошибка при копировании данных в резервную БД: %w", err)
	}
	log.Println("Резервная копия успешно создана:", backupFilePath)
	return nil
}

func (g GeoResponse) insertInfo(db *sql.DB) error {
	request := `
	INSERT INTO hosts
	(ip, host, country, region, City, ISP, Timezone)
	VALUES (:ip, :host, :country, :region, :city, :isp, :timezone)
	`
	sqlStmt, err := db.Prepare(request)
	if err != nil {
		return fmt.Errorf(errStmt, err)
	}
	defer func() { _ = sqlStmt.Close() }()
	startInsert := time.Now()
	res, err := sqlStmt.Exec(
		sql.Named("ip", g.IP),
		sql.Named("host", g.Host),
		sql.Named("country", g.Country),
		sql.Named("region", g.RegionName),
		sql.Named("city", g.City),
		sql.Named("isp", g.ISP),
		sql.Named("timezone", g.Timezone),
	)
	endInsert := time.Now()
	if err != nil {
		return fmt.Errorf(errExec, err)
	}
	resAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf(errRes, err)
	}
	if resAffected == 0 {
		return fmt.Errorf(errResAffected, err)
	}

	log.Println("информация внесена в БД за время:", endInsert.Sub(startInsert))
	return nil
}

func printLastTenEntries(db *sql.DB) error {
	request := `
        SELECT id, ip, host
		FROM hosts 
        ORDER BY timeAdd DESC 
        LIMIT 10
    `
	sqlStmt, err := db.Prepare(request)
	if err != nil {
		return fmt.Errorf(errStmt, err)
	}
	defer func() { _ = sqlStmt.Close() }()

	startExec := time.Now()
	rows, err := sqlStmt.Query()
	endExec := time.Now()
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос для получения последних записей: %w", err)
	}
	defer func() { rows.Close() }()

	fmt.Printf("Информация из БД:\n%s\n", strings.Repeat("-", 30))
	for rows.Next() {
		var (
			id   int
			ip   string
			host string
		)

		if err := rows.Scan(&id, &ip, &host); err != nil {
			return fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		fmt.Printf("ID: %d, IP: %s, хост: %s\n", id, ip, host)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("ошибка при итерации по строкам: %w", err)
	}
	fmt.Println(strings.Repeat("-", 30))
	log.Println("Данные прочитаны за: ", endExec.Sub(startExec))
	return nil
}

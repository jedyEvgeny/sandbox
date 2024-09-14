//Ищем IP или DNS

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const geoAPI = "http://ip-api.com/json/"

type GeoResponse struct {
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	ISP        string `json:"isp"`
	Query      string `json:"query"`
	Timezone   string `json:"timezone"`
}

func getGeoInfo(ip string) (*GeoResponse, error) {
	resp, err := http.Get(geoAPI + ip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geo GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, err
	}
	return &geo, nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Введите через флаг -dns или -ip сайта")
		os.Exit(1)
	}
	ip := mustFindIP()

	geo, err := getGeoInfo(ip)
	if err != nil {
		log.Printf("не удалось извлечь гео-информацию %s: %v\n", ip, err)
	}
	str := fmt.Sprintf("Общая информация:\n\t- Страна: %s\n\t- Город: %s\n\t- Компания: %s\n",
		geo.Country, geo.City, geo.ISP,
	)
	fmt.Print(str)
}

func mustFindIP() string {
	ip := flag.String("ip", "", "ip-адрес сайта")
	addr := flag.String("dns", "", "адрес сайта")
	flag.Parse()
	if ip == nil && addr == nil {
		log.Fatal("не введены через флаг -dns или -ip сайта")
	}
	if *ip != "" {
		address, err := net.LookupAddr(*ip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("IP Address: %s\n", *ip)
		fmt.Printf("Hostname: %s\n", address[0])
	}

	if *addr != "" {
		ip, err := net.LookupIP(*addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("IP Address: %s\n", ip)
		fmt.Printf("Hostname: %s\n", *addr)
	}
	return *ip
}

// Запуск бенчмарка go test -bench=. -benchmem
package main

import (
	"log"
	"os"
	"testing"
)

var (
	testOutput = output{
		proto:  "HTTP/1.1",
		method: "POST",
		query:  "https://some-dns.tech/api/Payout/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		body: []byte(`{
    "sum": 500000,
    "paymentMethod": "БАНК",
    "transferIdentifier": "1111111111",
    "fiat": "VND",
    "MerchantOrderIntId": "xxxxxxxxxxxxxxxxxx",
    "bankName": 79
}`),
		header: "Authorization: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\r\nContent-Type: application/json",
	}

	testInput = input{
		proto:  "HTTP/1.1",
		status: "200 OK",
		host:   "pay-hub.tech",
		header: "Connection: keep-alive\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Thu, 28 Aug 2025 09:47:31 GMT\r\nServer: nginx/1.24.0 (Ubuntu)\r\nX-Rate-Limit-Limit: 1m\r\nX-Rate-Limit-Remaining: 999\r\nX-Rate-Limit-Reset: 2025-08-28T09:48:31.1004079Z",
	}
)

func BenchmarkLogRequest(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testOutput.logRequest()
	}
}

func BenchmarkLogResponse(b *testing.B) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput.body = []byte(data)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testInput.logResponse()
	}
}

func BenchmarkLogRequestOptimized(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testOutput.logRequestOptimized()
	}
}

func BenchmarkLogResponseOptimized(b *testing.B) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput.body = []byte(data)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testInput.logResponseOptimized()
	}
}
  

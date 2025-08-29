package main

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	output struct {
		proto, method, query, header string
		body                         []byte
	}
	input struct {
		proto, status, host, header string
		body                        []byte
	}
)

func main() {
	o := output{
		proto:  "HTTP/1.1",
		method: "POST",
		query:  "https://pay-hub.tech/api/AvailableOrder/Payout/example-id",
		header: "Authorization: Bearer token\r\nContent-Type: application/json",
	}

	i := input{
		proto:  "HTTP/1.1",
		status: "400 Bad Request",
		host:   "pay-hub.tech",
		body:   []byte(`{"error": "Upcoming balance"}`),
		header: "Content-Type: application/json",
	}

	fmt.Println(o.logRequest())
	fmt.Println(i.logResponse())
	fmt.Println(o.logRequestOptimized())
	fmt.Println(i.logResponseOptimized())
}

// Логгер отправки запроса - текущая версия
func (o *output) logRequest() string {
	dump := new(bytes.Buffer)
	dump.WriteString(fmt.Sprintf("Making %s %s request to: %s", o.proto, o.method, o.query))
	dump.WriteString("\r\n")

	dump.WriteString(o.header)
	dump.WriteString("\r\n")

	if len(o.body) > 0 {
		dump.Write(o.body)
		dump.WriteString("\r\n")
	}

	return dump.String()
}

// Логгер получения ответа большого объёма - текущая версия
func (i *input) logResponse() string {
	dump := new(bytes.Buffer)
	dump.WriteString(fmt.Sprintf("Received %s response with status: [%s] from: %s", i.proto, i.status, i.host))
	dump.WriteString("\r\n")

	dump.WriteString(i.header)
	dump.WriteString("\r\n")

	if len(i.body) > 0 {
		dump.Write(i.body)
		dump.WriteString("\r\n")
	}

	return dump.String()
}

// Логер отправки запроса - оптимизация
func (o *output) logRequestOptimized() string {
	var dump strings.Builder
	dump.WriteString(fmt.Sprintf("Making %s %s request to: %s\n", o.proto, o.method, o.query))
	dump.WriteString("\r\n")

	dump.WriteString(o.header)
	dump.WriteString("\r\n")

	if len(o.body) > 0 {
		dump.Write(o.body)
		dump.WriteString("\r\n")
	}

	return dump.String()
}

// Логер получения ответа большого объёма - оптимизация
func (i *input) logResponseOptimized() string {
	var dump strings.Builder
	dump.WriteString(fmt.Sprintf("Received %s response with status: [%s] from: %s\n", i.proto, i.status, i.host))
	dump.WriteString("\r\n")

	dump.WriteString(i.header)
	dump.WriteString("\r\n")

	if len(i.body) > 0 {
		dump.Write(i.body)
		dump.WriteString("\r\n")
	}

	return dump.String()
}

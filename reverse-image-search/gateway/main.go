package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/web"
)

const (
	DEFAULT_API_URL     = "http://127.0.0.1:5000"
	DEFAULT_CONSOLE_URL = "http://127.0.0.1:8090"
	DEFAULT_PORT        = "3000"
)

func main() {
	var apiUrl = config.SetDataFromEnv("GATE_PROXY_API", DEFAULT_API_URL)
	var configUrl = config.SetDataFromEnv("GATE_PROXY_CONSOLE", DEFAULT_CONSOLE_URL)
	var appPort = config.SetDataFromEnv("GATE_APP_PORT", DEFAULT_PORT)

	fmt.Println("Proxy API Addr:", DEFAULT_API_URL)
	fmt.Println("Proxy Console API Addr:", DEFAULT_CONSOLE_URL)
	fmt.Println("Web Server port:", DEFAULT_PORT)

	web.API(apiUrl, configUrl, appPort)
}

package main

import http "github.com/roblesoft/stockapi/internal/controller/http"

func main() {
	server := http.NewServer()
	server.SetupRouter()
	server.Run()
}

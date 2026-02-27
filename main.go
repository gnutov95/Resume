package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Fly передаёт порт через переменную PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port
	fmt.Printf("Сервер запущен на %s\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

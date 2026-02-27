package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	// Обработка всех запросов
	http.Handle("/", fs)

	port := "8080"
	fmt.Printf("Сервер запущен на http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}

}

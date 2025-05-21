package main

import (
	"log"
	"messenger_user/database"
	"net/http"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Ошибка инициализации базы: %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

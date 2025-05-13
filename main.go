package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет, мир!")
	})
	log.Println("Сервер запущен на :8074")
	log.Fatal(http.ListenAndServe(":8074", nil))
}

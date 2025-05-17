package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет, весь мир!")
		fmt.Fprintln(w, "Привет, весь мир!")
		fmt.Fprintln(w, "Привет, весь мир!")
		fmt.Fprintln(w, "Работает! я спать.")
		fmt.Fprintln(w, "Работает! добра утра.")
	})
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

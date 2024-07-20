package main

import (
	"fmt"
	db "go-demo-service/config"
	"go-demo-service/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.Insert(w, r)
		}
	})
	db.Init()
	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Println("failed to start server", err)
	}
	fmt.Println("server started on 8080")

}

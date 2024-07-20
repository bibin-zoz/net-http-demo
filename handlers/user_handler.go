package handlers

import (
	"encoding/json"
	"fmt"
	models "go-demo-service/Models"
	db "go-demo-service/config"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	user := models.Admin{}

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	result := db.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("error to save data", result.Error)
	}
	fmt.Println("successfully isnerted data")

	json.NewEncoder(w).Encode(user.Email)
	w.WriteHeader(http.StatusOK)
}

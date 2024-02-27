package main

//Імпорт
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Коментарі

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	//Встановлює обробник запитів для шляху "/time"
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC3339)
		response := TimeResponse{Time: currentTime}
		jsonResponse, err := json.Marshal(response)
		//Перевірка на помилку
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})
	//виведення адреси у консоль
	fmt.Println("Сервер запущено на http://localhost:8795")
	//старт сервера
	http.ListenAndServe(":8795", nil)
}

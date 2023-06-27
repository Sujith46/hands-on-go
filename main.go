package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"message"`
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/post", postHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, World!"}
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody Message
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := Message{Text: "Received: " + requestBody.Text}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("JSON error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"time"
)

type HelloResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(Hello))
	mux.Handle("/hello2", http.HandlerFunc(Hello2))

	fmt.Println("build server. port is 8000...")
	err := http.ListenAndServe(":8000", handlers.CORS()(mux))
	if err != nil {
		fmt.Printf("server error: %+v\n", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	res := &HelloResponse{
		ID:      1,
		Message: "this is GO server!",
	}

	time.Sleep(10 * time.Second)

	b, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		fmt.Printf("return response failed. %+v\n", err)
	}
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	res := &HelloResponse{
		ID:      1,
		Message: "this is GO server2!",
	}

	time.Sleep(5 * time.Second)

	b, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		fmt.Printf("return response failed. %+v\n", err)
	}
}

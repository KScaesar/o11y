package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	sleepTime := time.Duration(rand.Intn(6)) * time.Second
	time.Sleep(sleepTime)
	fmt.Fprint(w, "Hello, Golang\n")
}

func main() {
	http.HandleFunc("/go/hello", helloHandler)
	fmt.Println("golang server is listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("golang ListenAndServe fail:", err)
	}
	fmt.Println("golang server stop")
}

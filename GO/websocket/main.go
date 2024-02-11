package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()
	fmt.Printf("server starting...")
	setupAPI(ctx)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err")
		log.Fatal("ListenAndServe: ", err)
	}
}

func setupAPI(ctx context.Context) {

	//ctx := context.Background()
	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	fmt.Printf("WebSocket endpoint registered. \n")
	http.HandleFunc("/login", manager.loginHandler)
	fmt.Printf("Login endpoint registered. \n")
	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.clients))
	})
	fmt.Printf("Debug endpoint registered. \n")
}

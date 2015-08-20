package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start() {
	initEnvironment()

	setRouteMap()

	startServer()
}

func initEnvironment() {
	fmt.Printf("%v init environment\n", time.Now())
}

func setRouteMap() {
	http.HandleFunc("/account", accountHandle)
	http.HandleFunc("/friend", friendHandle)
}

func startServer() {
	server := &http.Server{
		Addr:           ":9191",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

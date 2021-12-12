package main

import (
	"github.com/go-tour/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewCusRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

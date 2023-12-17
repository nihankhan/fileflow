package main

import (
	"fmt"
	"log"
	"net/http"

	internal "github.com/nihankhan/fileflow/internal/routers"
)

func main() {
	r := internal.Routers()
	server := &http.Server{
		Addr:    ":7000",
		Handler: r,
	}

	fmt.Println("FileFox server is Running on 127.0.0.1:7000")

	done := make(chan struct{})

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}

		done <- struct{}{}
	}()

	<-done
}

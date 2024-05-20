package main

import (
	"fmt"
	"net/http"

	"github.com/wbrijesh/trying-trpc/internal/haberdasherserver"
	"github.com/wbrijesh/trying-trpc/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}

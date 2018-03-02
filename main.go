package main

import (
	"home-collect/route"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":3000", route.MuxRouter().InitRouter()); err != nil {
		log.Fatal(err)
	}
}

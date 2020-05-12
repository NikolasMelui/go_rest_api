package main

import (
	"log"

	"github.com/nikolasmelui/go_rest_api/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

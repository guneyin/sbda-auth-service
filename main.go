package main

import (
	"fmt"
	"github.com/guneyin/sbda-auth-service/service"
	"log"
)

func main() {
	svc, err := service.NewService()
	if err != nil {
		fmt.Println(err)
	}

	err = svc.Register()
	if err != nil {
		fmt.Println(err)
	}

	if err = svc.Serve(); err != nil {
		log.Fatal(err)
	}
}

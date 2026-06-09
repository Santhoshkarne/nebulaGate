package main

import (
	"fmt"
	"log"

	"github.com/SaisrikarVollala/nebulagate/internal/config"
)

func main() {

	servers, err := config.LoadServers("configs/servers.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range servers {
		fmt.Println(s.ID, s.URL)
	}
}

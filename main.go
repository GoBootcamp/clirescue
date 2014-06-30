package main

import (
	"log"
	"os"

	"github.com/GoBootcamp/clirescue/trackerapi"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Command required")
	}
	switch os.Args[1] {
	case "login":
		err := trackerapi.CacheCredentials()
		if err != nil {
			panic(err)
		}
	default:
		log.Fatal("Unknown Command: ", os.Args[1])
	}
}

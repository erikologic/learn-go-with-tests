package main

import (
	"log"
	"net/http"

	poker "github.com/mrenrich84/learn-go-with-tests/playerServer"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}

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

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}

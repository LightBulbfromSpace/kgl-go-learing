package main

import (
	"fmt"
	poker "github.com/LightBulbfromSpace/build_an_application"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeDB()

	fmt.Println("Let's play poker")
	fmt.Println(`Type "{Name} wins" to record a win`)

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}

package main

import (
	"os"

	"github.com/thanawatjack/kawaii-shop-tutorial/config"
	"github.com/thanawatjack/kawaii-shop-tutorial/modules/servers"
	databases "github.com/thanawatjack/kawaii-shop-tutorial/pkg/databases/migrations"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.Newserver(cfg, db).Start()

}

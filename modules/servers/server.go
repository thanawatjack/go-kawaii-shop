package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/thanawatjack/kawaii-shop-tutorial/config"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	cfg config.IConfig
	db  *sqlx.DB
}

func Newserver(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Print("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// listen to host:port

	log.Printf("server is starting on %v ", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())

}

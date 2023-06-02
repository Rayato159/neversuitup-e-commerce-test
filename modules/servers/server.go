package servers

import (
	"encoding/json"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IServer interface {
	StartUsersServer()
	StartProductsServer()
	GetServer() *server
}

type server struct {
	App *fiber.App
	Cfg config.IConfig
	Db  *sqlx.DB
}

func NewServer(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		Cfg: cfg,
		Db:  db,
		App: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) GetServer() *server { return s }

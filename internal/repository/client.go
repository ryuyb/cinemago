package repository

import (
	"cinemago/internal/config"
	"cinemago/internal/model/ent"
	"context"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

var (
	globalClient Client
)

type Client struct {
	*ent.Client
	Driver *entSql.Driver
}

func NewClient() (*Client, func(), error) {
	databaseCfg := config.GetConfig().Database

	var (
		drv *entSql.Driver
		err error
	)
	switch databaseCfg.Driver {
	case "sqlite3", "sqlite":
		drv, err = entSql.Open(dialect.SQLite, databaseCfg.Source)
	case "pgx", "postgres":
		drv, err = entSql.Open("pgx", databaseCfg.Source)
	default:
		log.Fatalf("Unknown database driver: %s", databaseCfg.Driver)
	}
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	entClient := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			log.Debug(a...)
		}),
	)

	if databaseCfg.Migrate {
		if err := entClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed migrate database schema: %v", err)
		}
	}

	cleanup := func() {
		log.Info("closing the database resources")
		if err := entClient.Close(); err != nil {
			log.Errorf("failed close database resources: %v", err)
		}
	}

	globalClient = Client{
		Client: entClient,
		Driver: drv,
	}

	return &globalClient, cleanup, nil
}

func GetClient() *Client {
	return &globalClient
}

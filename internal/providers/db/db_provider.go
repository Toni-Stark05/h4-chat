package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"h4-chat/internal/config"
	"log"
	"log/slog"
)

type PGProvider struct {
	DB *sql.DB
}

func NewPGProvider(i *do.Injector) (*PGProvider, error) {
	cfg := do.MustInvoke[*config.Config](i)
	logger := do.MustInvoke[*slog.Logger](i)

	dsn := getDns(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("could not connect to db", "error", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("could not connect db: ", "error: ", err)
	}

	logger.Info("connected to db")
	return &PGProvider{DB: db}, nil
}

func getDns(c *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
	)
}

func (p *PGProvider) Close() {
	p.DB.Close()
}

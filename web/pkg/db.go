package pkg

import (
	"_schemas/ent"
	"context"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func InitDB(logger *zap.Logger) *ent.Client {
	// client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		logger.Fatal("failed opening connection to mysql", zap.Error(err))
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	return client
}

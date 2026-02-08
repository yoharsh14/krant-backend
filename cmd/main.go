package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"
	"yoharsh14/krant-backend/internal/env"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cfg := config{
		addr: ":8082",
		db:   dbConfig{dsn: env.GetString("MONGO_DRIVER", "mongodb+srv://harshdambhareh53_db_user:root@cluster0.jvytesx.mongodb.net/?appName=Cluster0")},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.db.dsn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)

	if err != nil {
		log.Panic("Error in connecting with database", err)
		os.Exit(1)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	logger.Info("Connected with database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     client,
	}

	if err := api.run(api.mount(ctx)); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}

}

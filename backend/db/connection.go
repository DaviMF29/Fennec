package db

import (
	"context"
	"fmt"
	"time"

	"github.com/DaviMF29/wombat/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (*mongo.Client, error) {
	conf := config.GetDB()

	if conf.URI == "" {
		return nil, fmt.Errorf("MongoDB URI não configurada")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB Atlas: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("erro ao verificar a conexão com o MongoDB Atlas: %v", err)
	}

	return client, nil
}

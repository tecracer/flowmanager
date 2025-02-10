package fm_connect

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/connect"
)

func InitClient(ctx context.Context) *connect.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client := connect.NewFromConfig(cfg)
	return client
}

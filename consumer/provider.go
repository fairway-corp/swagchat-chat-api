package consumer

import (
	"context"

	"github.com/swagchat/chat-api/config"
)

type provider interface {
	SubscribeMessage() error
	UnsubscribeMessage() error
}

func Provider(ctx context.Context) provider {
	cfg := config.Config()

	var p provider
	switch cfg.Consumer.Provider {
	case "nsq":
		p = &nsqProvider{
			ctx: ctx,
		}
	case "kafka":
		p = &kafkaProvider{
			ctx: ctx,
		}
	default:
		p = &noopProvider{
			ctx: ctx,
		}
	}

	return p
}

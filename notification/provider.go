package notification

import (
	"context"

	"github.com/swagchat/chat-api/config"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

// [APNS] https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html
// [FCM] https://firebase.google.com/docs/cloud-messaging/concept-options

type MessageInfo struct {
	Text  string
	Badge int
}

type NotificationResult struct {
	Data  interface{}
	Error error
}

type NotificationChannel chan NotificationResult

type provider interface {
	CreateTopic(string) NotificationChannel
	DeleteTopic(string) NotificationChannel
	CreateEndpoint(userID string, platform scpb.Platform, token string) NotificationChannel
	DeleteEndpoint(string) NotificationChannel
	Subscribe(string, string) NotificationChannel
	Unsubscribe(string) NotificationChannel
	Publish(string, string, *MessageInfo) NotificationChannel
}

func Provider(ctx context.Context) provider {
	cfg := config.Config()
	var p provider

	switch cfg.Notification.Provider {
	case "awsSns":
		p = &awssnsProvider{
			ctx:                   ctx,
			region:                cfg.Notification.AmazonSNS.Region,
			accessKeyId:           cfg.Notification.AmazonSNS.AccessKeyID,
			secretAccessKey:       cfg.Notification.AmazonSNS.SecretAccessKey,
			roomTopicNamePrefix:   cfg.Notification.RoomTopicNamePrefix,
			applicationArnIos:     cfg.Notification.AmazonSNS.ApplicationArnIos,
			applicationArnAndroid: cfg.Notification.AmazonSNS.ApplicationArnAndroid,
		}
	default:
		p = &noopProvider{
			ctx: ctx,
		}
	}

	return p
}

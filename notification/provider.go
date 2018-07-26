package notification

import (
	"context"
	"net/http"

	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/utils"
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
	CreateEndpoint(string, int32, string) NotificationChannel
	DeleteEndpoint(string) NotificationChannel
	Subscribe(string, string) NotificationChannel
	Unsubscribe(string) NotificationChannel
	Publish(context.Context, string, string, *MessageInfo) NotificationChannel
}

func Provider() provider {
	cfg := utils.Config()
	var p provider

	switch cfg.Notification.Provider {
	case "awsSns":
		p = &awssnsProvider{
			region:                cfg.Notification.AmazonSNS.Region,
			accessKeyId:           cfg.Notification.AmazonSNS.AccessKeyID,
			secretAccessKey:       cfg.Notification.AmazonSNS.SecretAccessKey,
			roomTopicNamePrefix:   cfg.Notification.RoomTopicNamePrefix,
			applicationArnIos:     cfg.Notification.AmazonSNS.ApplicationArnIos,
			applicationArnAndroid: cfg.Notification.AmazonSNS.ApplicationArnAndroid,
		}
	default:
		p = &notuseProvider{}
	}

	return p
}

func createProblemDetail(msg string, err error) *model.ProblemDetail {
	return &model.ProblemDetail{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   err,
	}
}

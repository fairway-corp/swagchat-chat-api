package datastore

import (
	"github.com/swagchat/chat-api/model"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

type selectMessagesOptions struct {
	roomID          string
	roleIDs         []int32
	limitTimestamp  int64
	offsetTimestamp int64
	orders          []*scpb.OrderInfo
}

type SelectMessagesOption func(*selectMessagesOptions)

func SelectMessagesOptionFilterByRoomID(roomID string) SelectMessagesOption {
	return func(ops *selectMessagesOptions) {
		ops.roomID = roomID
	}
}

func SelectMessagesOptionFilterByRoleIDs(roleIDs []int32) SelectMessagesOption {
	return func(ops *selectMessagesOptions) {
		ops.roleIDs = roleIDs
	}
}

func SelectMessagesOptionOrders(orders []*scpb.OrderInfo) SelectMessagesOption {
	return func(ops *selectMessagesOptions) {
		ops.orders = orders
	}
}

func SelectMessagesOptionLimitTimestamp(limitTimestamp int64) SelectMessagesOption {
	return func(ops *selectMessagesOptions) {
		ops.limitTimestamp = limitTimestamp
	}
}

func SelectMessagesOptionOffsetTimestamp(offsetTimestamp int64) SelectMessagesOption {
	return func(ops *selectMessagesOptions) {
		ops.offsetTimestamp = offsetTimestamp
	}
}

type messageStore interface {
	createMessageStore()

	InsertMessage(message *model.Message) error
	SelectMessages(limit, offset int32, opts ...SelectMessagesOption) ([]*model.Message, error)
	SelectMessage(messageID string) (*model.Message, error)
	SelectCountMessages(opts ...SelectMessagesOption) (int64, error)
	UpdateMessage(message *model.Message) error
}

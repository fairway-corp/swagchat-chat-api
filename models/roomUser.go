package models

import (
	"net/http"

	"github.com/fairway-corp/swagchat-api/utils"
)

type RoomUser struct {
	RoomId      string         `json:"roomId,omitempty" db:"room_id"`
	UserId      string         `json:"userId,omitempty" db:"user_id"`
	UnreadCount *int64         `json:"unreadCount,omitempty" db:"unread_count"`
	MetaData    utils.JSONText `json:"metaData,omitempty" db:"meta_data"`
	Created     int64          `json:"created,omitempty" db:"created"`
}

type ErrorRoomUser struct {
	UserId string         `json:"userId,omitempty"`
	Error  *ProblemDetail `json:"error"`
}

type ResponseRoomUser struct {
	RoomUsers []RoomUser      `json:"roomUsers,omitempty"`
	Errors    []ErrorRoomUser `json:"errors,omitempty"`
}

type RequestRoomUserIds struct {
	UserIds []string `json:"userIds,omitempty"`
}

type RoomUsers struct {
	RoomUsers []*RoomUser `json:"roomUsers"`
}

func (rus *RequestRoomUserIds) IsValid() *ProblemDetail {
	if len(rus.UserIds) == 0 {
		return &ProblemDetail{
			Title:     "Request parameter error. (Create room's user list)",
			Status:    http.StatusBadRequest,
			ErrorName: ERROR_NAME_INVALID_PARAM,
			InvalidParams: []InvalidParam{
				InvalidParam{
					Name:   "userIds",
					Reason: "Not set.",
				},
			},
		}
	}

	return nil
}

func (rus *RequestRoomUserIds) RemoveDuplicate() {
	rus.UserIds = utils.RemoveDuplicate(rus.UserIds)
}
package model

import (
	"net/http"

	scpb "github.com/swagchat/protobuf"
)

type CreateRoomUsersRequest struct {
	scpb.CreateRoomUsersRequest
	Room *Room
}

func (crur *CreateRoomUsersRequest) GenerateRoomUsers() []*RoomUser {
	roomUsers := make([]*RoomUser, len(crur.UserIDs))
	for i, userID := range crur.UserIDs {
		ru := &RoomUser{}
		ru.RoomID = crur.RoomID
		ru.UserID = userID
		ru.UnreadCount = int32(0)
		ru.Display = crur.Display
		roomUsers[i] = ru
	}
	return roomUsers
}

func (crur *CreateRoomUsersRequest) Validate() *ErrorResponse {
	if crur.Room.Type == scpb.RoomType_OneOnOne {
		if len(crur.UserIDs) != 1 {
			invalidParams := []*scpb.InvalidParam{
				&scpb.InvalidParam{
					Name:   "userIds",
					Reason: "In case of 1-on-1 room type, only one user can be specified for userIDs.",
				},
			}
			return NewErrorResponse("Failed to create a user.", invalidParams, http.StatusBadRequest, nil)
		}
	}
	return nil
}

type GetUserIdsOfRoomUserRequest struct {
	scpb.GetUserIdsOfRoomUserRequest
}

type UpdateRoomUserRequest struct {
	scpb.UpdateRoomUserRequest
}

type RoomUser struct {
	scpb.RoomUser
}

func (ru *RoomUser) UpdateRoomUser(req *UpdateRoomUserRequest) {
	if req.UnreadCount != nil {
		ru.UnreadCount = *req.UnreadCount
	}

	if req.Display != nil {
		ru.Display = *req.Display
	}
}

type DeleteRoomUsersRequest struct {
	scpb.DeleteRoomUsersRequest
	Room *Room
}

func (drur *DeleteRoomUsersRequest) Validate() *ProblemDetail {
	if drur.Room.Type == scpb.RoomType_OneOnOne {
		if len(drur.Room.Users)-len(drur.UserIDs) != 1 {
			return &ProblemDetail{
				Message: "Invalid params",
				InvalidParams: []*InvalidParam{
					&InvalidParam{
						Name:   "userIds",
						Reason: "In case of 1-on-1 room type, only one user must be specified.",
					},
				},
				Status: http.StatusBadRequest,
			}
		}
	}
	return nil
}

// type RoomUsers struct {
// 	RoomID    string
// 	UserIDs   []string
// 	Display   bool
// 	Room      *Room
// 	RoomUsers []*RoomUser
// }

// type RoomUser struct {
// 	RoomID      string `json:"roomId,omitempty" db:"room_id"`
// 	UserID      string `json:"userId,omitempty" db:"user_id"`
// 	UnreadCount int32  `json:"unreadCount,omitempty" db:"unread_count"`
// 	Display     bool   `json:"display,omitempty" db:"display"`
// }

// type ErrorRoomUser struct {
// 	UserId string         `json:"userId,omitempty"`
// 	Error  *ProblemDetail `json:"error"`
// }

// type ResponseRoomUser struct {
// 	RoomUsers []RoomUser      `json:"roomUsers,omitempty"`
// 	Errors    []ErrorRoomUser `json:"errors,omitempty"`
// }

// type RequestRoomUserIDs struct {
// 	UserIDs []string `json:"userIds,omitempty" db:"-"`
// }

// ImportFromPbCreateUserRolesRequest import from CreateUserRolesRequest proto
// func (rus *RoomUsers) ImportFromPbCreateUserRolesRequest(req *scpb.CreateRoomUsersRequest) {
// 	rus.RoomID = req.RoomID
// 	rus.UserIDs = utils.RemoveDuplicate(req.UserIDs)
// 	rus.Display = *req.Display
// }

// func (rus *RoomUsers) Validate() *ProblemDetail {
// 	if len(rus.UserIDs) == 0 {
// 		return &ProblemDetail{
// 			Message: "Invalid params",
// 			InvalidParams: []*InvalidParam{
// 				&InvalidParam{
// 					Name:   "userIds",
// 					Reason: "Not set.",
// 				},
// 			},
// 			Status: http.StatusBadRequest,
// 		}
// 	}

// 	if rus.Room.Type == OneOnOne {
// 		if len(rus.UserIDs) == 2 {
// 			return &ProblemDetail{
// 				Message: "Invalid params",
// 				InvalidParams: []*InvalidParam{
// 					&InvalidParam{
// 						Name:   "userIds",
// 						Reason: "In case of 1-on-1 room type, It can only update once.",
// 					},
// 				},
// 				Status: http.StatusBadRequest,
// 			}
// 		}
// 	}

// 	return nil
// }

// // ImportFromPbUpdateUserRoleRequest import from UpdateUserRoleRequest proto
// func (ru *RoomUser) ImportFromPbUpdateUserRoleRequest(req *scpb.UpdateRoomUserRequest) {
// 	ru.RoomID = req.RoomID
// 	ru.UserID = req.UserID
// 	ru.UnreadCount = *req.UnreadCount
// }

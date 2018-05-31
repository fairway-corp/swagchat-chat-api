package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/swagchat/chat-api/utils"
)

const (
	MessageTypeText           = "text"
	MessageTypeImage          = "image"
	MessageTypeFile           = "file"
	MessageTypeIndicatorStart = "indicatorStart"
	MessageTypeIndicatorEnd   = "indicatorEnd"
	MessageTypeUpdateRoomUser = "updateRoomUser"
)

type Messages struct {
	Messages []*Message `json:"messages" db:"-"`
	AllCount int64      `json:"allCount" db:"all_count"`
}

type Message struct {
	ID               uint64         `json:"-" db:"id"`
	MessageID        string         `json:"messageId" db:"message_id,notnull"`
	SuggestMessageID string         `json:"suggestMessageId" db:"suggest_message_id,notnull"`
	RoomID           string         `json:"roomId" db:"room_id,notnull"`
	UserID           string         `json:"userId" db:"user_id,notnull"`
	Type             string         `json:"type,omitempty" db:"type"`
	EventName        string         `json:"eventName,omitempty" db:"-"`
	Payload          utils.JSONText `json:"payload" db:"payload"`
	Role             *Role          `json:"role,omitempty" db:"role,notnull"`
	Created          int64          `json:"created" db:"created,notnull"`
	Modified         int64          `json:"modified" db:"modified,notnull"`
	Deleted          int64          `json:"-" db:"deleted,notnull"`
}

func (m *Message) MarshalJSON() ([]byte, error) {
	l, _ := time.LoadLocation("Etc/GMT")
	return json.Marshal(&struct {
		MessageID        string         `json:"messageId"`
		SuggestMessageID string         `json:"suggestMessageId"`
		RoomID           string         `json:"roomId"`
		UserID           string         `json:"userId"`
		Type             string         `json:"type"`
		EventName        string         `json:"eventName,omitempty"`
		Payload          utils.JSONText `json:"payload"`
		Role             *Role          `json:"role"`
		Created          string         `json:"created"`
		Modified         string         `json:"modified"`
	}{
		MessageID:        m.MessageID,
		SuggestMessageID: m.SuggestMessageID,
		RoomID:           m.RoomID,
		UserID:           m.UserID,
		Type:             m.Type,
		EventName:        m.EventName,
		Payload:          m.Payload,
		Role:             m.Role,
		Created:          time.Unix(m.Created, 0).In(l).Format(time.RFC3339),
		Modified:         time.Unix(m.Modified, 0).In(l).Format(time.RFC3339),
	})
}

type ResponseMessages struct {
	MessageIds []string         `json:"messageIds,omitempty"`
	Errors     []*ProblemDetail `json:"errors,omitempty"`
}

type PayloadText struct {
	Text string `json:"text"`
}

type PayloadImage struct {
	Mime         string `json:"mime"`
	Filename     string `json:"filename"`
	SourceUrl    string `json:"sourceUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type PayloadTextSuggest struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
}

type PayloadButtons struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PayloadConfirm struct {
	Text string `json:"text"`
}

type PayloadList struct {
	Text string `json:"text"`
}

type PayloadCarousel struct {
	Columns []PayloadCarouselColumn `json:"columns"`
}

type PayloadCarouselColumn struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PayloadLocation struct {
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PayloadUsers struct {
	Users []string `json:"users"`
}

func (m *Message) IsValid() *ProblemDetail {
	if m.MessageID != "" && !utils.IsValidID(m.MessageID) {
		return &ProblemDetail{
			Title:  "Request error",
			Status: http.StatusBadRequest,
			InvalidParams: []InvalidParam{
				InvalidParam{
					Name:   "messageId",
					Reason: "messageId is invalid. Available characters are alphabets, numbers and hyphens.",
				},
			},
		}
	}

	if m.Payload == nil {
		return &ProblemDetail{
			Title:  "Request error",
			Status: http.StatusBadRequest,
			InvalidParams: []InvalidParam{
				InvalidParam{
					Name:   "payload",
					Reason: "payload is empty.",
				},
			},
		}
	}

	if m.Type == MessageTypeText {
		var pt PayloadText
		json.Unmarshal(m.Payload, &pt)
		if pt.Text == "" {
			return &ProblemDetail{
				Title:  "Request error",
				Status: http.StatusBadRequest,
				InvalidParams: []InvalidParam{
					InvalidParam{
						Name:   "payload",
						Reason: "Text type needs text.",
					},
				},
			}
		}
	}

	if m.Type == MessageTypeImage {
		var pi PayloadImage
		json.Unmarshal(m.Payload, &pi)
		if pi.Mime == "" {
			return &ProblemDetail{
				Title:  "Request error",
				Status: http.StatusBadRequest,
				InvalidParams: []InvalidParam{
					InvalidParam{
						Name:   "payload",
						Reason: "Image type needs mime.",
					},
				},
			}
		}

		if pi.SourceUrl == "" {
			return &ProblemDetail{
				Title:  "Request error",
				Status: http.StatusBadRequest,
				InvalidParams: []InvalidParam{
					InvalidParam{
						Name:   "payload",
						Reason: "Image type needs sourceUrl.",
					},
				},
			}
		}
	}

	return nil
}

func (m *Message) BeforeSave() {
	if m.MessageID == "" {
		m.MessageID = utils.GenerateUUID()
	}

	if m.Role == nil {
		general := RoleGeneral
		m.Role = &general
	}

	nowTimestamp := time.Now().Unix()
	if m.Created == 0 {
		m.Created = nowTimestamp
	}
	m.Modified = nowTimestamp
}

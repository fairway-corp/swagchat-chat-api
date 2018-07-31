package datastore

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/utils"
)

var (
	ctx context.Context
)

func TestMain(m *testing.M) {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	cfg := utils.Config()
	cfg.Logger.EnableConsole = false
	cfg.Datastore.SQLite.OnMemory = true
	Provider(ctx).Connect(cfg.Datastore)
	Provider(ctx).CreateTables()

	nowTimestamp := time.Now().Unix()

	var newUser *model.User
	userRoles := make([]*model.UserRole, 20, 20)

	for i := 1; i <= 10; i++ {
		userID := fmt.Sprintf("datastore-user-id-%04d", i)

		newUser = &model.User{}
		newUser.UserID = userID
		newUser.MetaData = []byte(`{"key":"value"}`)
		newUser.LastAccessed = nowTimestamp + int64(i)
		newUser.Created = nowTimestamp + int64(i)
		newUser.Modified = nowTimestamp + int64(i)
		err := Provider(ctx).InsertUser(newUser)
		if err != nil {
			fmt.Errorf("Failed to insert user on main test")
		}

		newUserRole := &model.UserRole{}
		newUserRole.UserID = userID
		newUserRole.Role = 1
		userRoles[i-1] = newUserRole
	}
	for i := 11; i <= 20; i++ {
		userID := fmt.Sprintf("datastore-user-id-%04d", i)

		newUser = &model.User{}
		newUser.UserID = fmt.Sprintf("datastore-user-id-%04d", i)
		newUser.MetaData = []byte(`{"key":"value"}`)
		newUser.LastAccessed = nowTimestamp
		newUser.Created = nowTimestamp + int64(i)
		newUser.Modified = nowTimestamp + int64(i)
		err := Provider(ctx).InsertUser(newUser)
		if err != nil {
			fmt.Errorf("Failed to insert user on main test")
		}

		newUserRole := &model.UserRole{}
		newUserRole.UserID = userID
		newUserRole.Role = 2
		userRoles[i-1] = newUserRole
	}

	err := Provider(ctx).InsertUserRoles(userRoles)
	if err != nil {
		fmt.Errorf("Failed to insert user roles on main test")
	}

	var newRoom *model.Room
	for i := 1; i <= 10; i++ {
		newRoom = &model.Room{}
		newRoom.RoomID = fmt.Sprintf("datastore-room-id-%04d", i)
		newRoom.UserID = fmt.Sprintf("datastore-user-id-%04d", i)
		newRoom.Type = 1
		newRoom.MetaData = []byte(`{"key":"value"}`)
		newRoom.LastMessageUpdated = nowTimestamp + int64(i)
		newRoom.Created = nowTimestamp + int64(i)
		newRoom.Modified = nowTimestamp + int64(i)
		err := Provider(ctx).InsertRoom(newRoom)
		if err != nil {
			fmt.Errorf("Failed to insert room on main test")
		}
	}
	for i := 11; i <= 20; i++ {
		newRoom = &model.Room{}
		newRoom.RoomID = fmt.Sprintf("datastore-room-id-%04d", i)
		newRoom.UserID = fmt.Sprintf("datastore-user-id-%04d", i)
		newRoom.Type = 2
		newRoom.MetaData = []byte(`{"key":"value"}`)
		newRoom.LastMessageUpdated = nowTimestamp
		newRoom.Created = nowTimestamp + int64(i)
		newRoom.Modified = nowTimestamp + int64(i)
		err := Provider(ctx).InsertRoom(newRoom)
		if err != nil {
			fmt.Errorf("Failed to insert room on main test")
		}
	}

	time.Sleep(1 * time.Second)

	code := m.Run()
	Provider(ctx).Close()
	os.Exit(code)
}

package datastore

import "github.com/swagchat/chat-api/model"

func (p *mysqlProvider) createRoomStore() {
	rdbCreateRoomStore(p.database)
}

func (p *mysqlProvider) InsertRoom(room *model.Room, opts ...RoomOption) (*model.Room, error) {
	return rdbInsertRoom(p.database, room, opts...)
}

func (p *mysqlProvider) SelectRoom(roomID string) (*model.Room, error) {
	return rdbSelectRoom(p.database, roomID)
}

func (p *mysqlProvider) SelectRooms() ([]*model.Room, error) {
	return rdbSelectRooms(p.database)
}

func (p *mysqlProvider) SelectUsersForRoom(roomID string) ([]*model.UserForRoom, error) {
	return rdbSelectUsersForRoom(p.database, roomID)
}

func (p *mysqlProvider) SelectCountRooms() (int64, error) {
	return rdbSelectCountRooms(p.database)
}

func (p *mysqlProvider) UpdateRoom(room *model.Room) (*model.Room, error) {
	return rdbUpdateRoom(p.database, room)
}

func (p *mysqlProvider) UpdateRoomDeleted(roomID string) error {
	return rdbUpdateRoomDeleted(p.database, roomID)
}

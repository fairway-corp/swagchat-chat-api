package datastore

import (
	"fmt"
	"time"

	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/model"
)

func rdbCreateRoomStore(db string) {
	master := RdbStore(db).master()

	tableMap := master.AddTableWithName(model.Room{}, tableNameRoom)
	tableMap.SetKeys(true, "id")
	for _, columnMap := range tableMap.Columns {
		if columnMap.ColumnName == "room_id" {
			columnMap.SetUnique(true)
		}
	}
	err := master.CreateTablesIfNotExists()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while creating room table. %v.", err))
		return
	}
}

func rdbInsertRoom(db string, room *model.Room, opts ...RoomOption) error {
	master := RdbStore(db).master()

	opt := roomOptions{}
	for _, o := range opts {
		o(&opt)
	}

	trans, err := master.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while inserting room. %v.", err))
		return err
	}

	err = trans.Insert(room)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while inserting room. %v.", err))
		return err
	}

	for _, ru := range opt.users {
		err = trans.Insert(ru)
		if err != nil {
			trans.Rollback()
			logger.Error(fmt.Sprintf("An error occurred while inserting room. %v.", err))
			return err
		}
	}

	err = trans.Commit()
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while inserting room. %v.", err))
		return err
	}

	return nil
}

func rdbSelectRooms(db string, limit, offset int32, opts ...RoomOption) ([]*model.Room, error) {
	replica := RdbStore(db).replica()

	opt := roomOptions{}
	for _, o := range opts {
		o(&opt)
	}

	var rooms []*model.Room
	query := fmt.Sprintf(`SELECT
	room_id,
	user_id,
	name,
	picture_url,
	information_url,
	meta_data,
	type,
	last_message,
	last_message_updated,
	created,
	modified
	FROM %s
	WHERE deleted = 0`, tableNameRoom)
	params := make(map[string]interface{})

	query = fmt.Sprintf("%s ORDER BY", query)
	if opt.orders == nil {
		query = fmt.Sprintf("%s created DESC", query)
	} else {
		i := 1
		for k, v := range opt.orders {
			query = fmt.Sprintf("%s %s %s", query, k, v.String())
			if i < len(opt.orders) {
				query = fmt.Sprintf("%s,", query)
			}
			i++
		}
	}

	query = fmt.Sprintf("%s LIMIT :limit OFFSET :offset", query)
	params["limit"] = limit
	params["offset"] = offset
	println(query)
	_, err := replica.Select(&rooms, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting rooms. %v.", err))
		return nil, err
	}

	return rooms, nil
}

func rdbSelectRoom(db, roomID string) (*model.Room, error) {
	replica := RdbStore(db).replica()

	var rooms []*model.Room
	query := fmt.Sprintf("SELECT * FROM %s WHERE room_id=:roomId AND deleted=0;", tableNameRoom)
	params := map[string]interface{}{"roomId": roomID}
	_, err := replica.Select(&rooms, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting room. %v.", err))
		return nil, err
	}

	if len(rooms) == 1 {
		return rooms[0], nil
	}

	return nil, nil
}

func rdbSelectUsersForRoom(db, roomID string) ([]*model.UserForRoom, error) {
	replica := RdbStore(db).replica()

	var users []*model.UserForRoom
	query := fmt.Sprintf(`SELECT
u.user_id,
u.name,
u.picture_url,
u.information_url,
u.meta_data,
u.can_block,
u.last_accessed,
u.created,
u.modified,
ru.display AS ru_display
FROM %s AS ru 
LEFT JOIN %s AS u ON ru.user_id = u.user_id 
WHERE ru.room_id = :roomId AND u.deleted = 0 
ORDER BY u.created;`, tableNameRoomUser, tableNameUser)
	params := map[string]interface{}{"roomId": roomID}
	_, err := replica.Select(&users, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting users for room. %v.", err))
		return nil, err
	}

	return users, nil
}

func rdbSelectCountRooms(db string, opts ...RoomOption) (int64, error) {
	replica := RdbStore(db).replica()

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE deleted = 0;", tableNameRoom)
	count, err := replica.SelectInt(query)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting room count. %v.", err))
		return 0, err
	}

	return count, nil
}

func rdbUpdateRoom(db string, room *model.Room) error {
	master := RdbStore(db).master()

	if room.Deleted != 0 {
		return rdbUpdateRoomDeleted(db, room.RoomID)
	}

	_, err := master.Update(room)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while updating room. %v.", err))
		return err
	}

	return nil
}

func rdbUpdateRoomDeleted(db, roomID string) error {
	master := RdbStore(db).master()
	trans, err := master.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while deleting room. %v.", err))
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE room_id=:roomId;", tableNameRoomUser)
	params := map[string]interface{}{
		"roomId": roomID,
	}
	_, err = trans.Exec(query, params)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting room. %v.", err))
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET deleted=:deleted WHERE room_id=:roomId;", tableNameSubscription)
	params = map[string]interface{}{
		"roomId":  roomID,
		"deleted": time.Now().Unix(),
	}
	_, err = trans.Exec(query, params)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting room. %v.", err))
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET deleted=:deleted WHERE room_id=:roomId;", tableNameRoom)
	params = map[string]interface{}{
		"roomId":  roomID,
		"deleted": time.Now().Unix(),
	}
	_, err = trans.Exec(query, params)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting room. %v.", err))
		return err
	}

	err = trans.Commit()
	if err != nil {
		err := trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting room. %v.", err))
		return err
	}

	return nil
}

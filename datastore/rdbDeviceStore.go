package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/tracer"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

func rdbCreateDeviceStore(ctx context.Context, db string) {
	span := tracer.Provider(ctx).StartSpan("rdbCreateDeviceStore", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	master := RdbStore(db).master()

	tableMap := master.AddTableWithName(model.Device{}, tableNameDevice)
	tableMap.SetUniqueTogether("user_id", "platform")
	for _, columnMap := range tableMap.Columns {
		if columnMap.ColumnName == "token" || columnMap.ColumnName == "notification_device_id" {
			columnMap.SetUnique(true)
		}
	}
	err := master.CreateTablesIfNotExists()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while creating device table. %v.", err))
		return
	}
}

func rdbInsertDevice(ctx context.Context, db string, device *model.Device) error {
	span := tracer.Provider(ctx).StartSpan("rdbInsertDevice", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	master := RdbStore(db).master()

	if err := master.Insert(device); err != nil {
		logger.Error(fmt.Sprintf("An error occurred while inserting device. %v.", err))
		return err
	}

	return nil
}

func rdbSelectDevices(ctx context.Context, db string, opts ...SelectDevicesOption) ([]*model.Device, error) {
	span := tracer.Provider(ctx).StartSpan("rdbSelectDevices", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	replica := RdbStore(db).replica()

	opt := selectDevicesOptions{}
	for _, o := range opts {
		o(&opt)
	}

	if opt.userID == "" && opt.platform == scpb.Platform_PlatformNone && opt.token == "" {
		return nil, errors.New("Be sure to specify either userId or platform or token")
	}

	var devices []*model.Device
	query := fmt.Sprintf("SELECT * FROM %s WHERE ", tableNameDevice)
	params := map[string]interface{}{}

	if opt.userID != "" {
		query = fmt.Sprintf("%s user_id=:userId AND", query)
		params["userId"] = opt.userID
	}

	if opt.platform != scpb.Platform_PlatformNone {
		query = fmt.Sprintf("%s platform=:platform AND", query)
		params["platform"] = opt.platform
	}

	if opt.token != "" {
		query = fmt.Sprintf("%s token=:token AND", query)
		params["token"] = opt.token
	}

	query = query[0 : len(query)-len(" AND")]

	_, err := replica.Select(&devices, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting devices. %v.", err))
		return nil, err
	}

	return devices, nil
}

func rdbSelectDevice(ctx context.Context, db, userID string, platform scpb.Platform) (*model.Device, error) {
	span := tracer.Provider(ctx).StartSpan("rdbSelectDevice", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	replica := RdbStore(db).replica()

	var devices []*model.Device
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=:userId AND platform=:platform;", tableNameDevice)
	params := map[string]interface{}{
		"userId":   userID,
		"platform": platform,
	}
	_, err := replica.Select(&devices, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting device. %v.", err))
		return nil, err
	}

	if len(devices) == 1 {
		return devices[0], nil
	}

	return nil, nil
}

func rdbSelectDevicesByUserID(ctx context.Context, db, userID string) ([]*model.Device, error) {
	span := tracer.Provider(ctx).StartSpan("rdbSelectDevicesByUserID", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	replica := RdbStore(db).replica()

	var devices []*model.Device
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=:userId;", tableNameDevice)
	params := map[string]interface{}{
		"userId": userID,
	}
	_, err := replica.Select(&devices, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting devices by userId. %v.", err))
		return nil, err
	}

	return devices, nil
}

func rdbSelectDevicesByToken(ctx context.Context, db, token string) ([]*model.Device, error) {
	span := tracer.Provider(ctx).StartSpan("rdbSelectDevicesByToken", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	replica := RdbStore(db).replica()

	var devices []*model.Device
	query := fmt.Sprintf("SELECT * FROM %s WHERE token=:token;", tableNameDevice)
	params := map[string]interface{}{
		"token": token,
	}
	_, err := replica.Select(&devices, query, params)
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting device by token. %v.", err))
		return nil, err
	}

	return devices, nil
}

func rdbUpdateDevice(ctx context.Context, db string, device *model.Device) error {
	span := tracer.Provider(ctx).StartSpan("rdbUpdateDevice", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	master := RdbStore(db).master()
	trans, err := master.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while updating device. %v.", err))
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET deleted=? WHERE user_id=? AND platform=?;", tableNameSubscription)
	_, err = trans.Exec(query, time.Now().Unix(), device.UserID, device.Platform)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while updating device. %v.", err))
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET token=?, notification_device_id=? WHERE user_id=? AND platform=?;", tableNameDevice)
	_, err = trans.Exec(query, device.Token, device.NotificationDeviceID, device.UserID, device.Platform)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while updating device. %v.", err))
		return err
	}

	err = trans.Commit()
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while updating device. %v.", err))
		return err
	}

	return nil
}

func rdbDeleteDevice(ctx context.Context, db, userID string, platform scpb.Platform) error {
	span := tracer.Provider(ctx).StartSpan("rdbDeleteDevice", "datastore")
	defer tracer.Provider(ctx).Finish(span)

	master := RdbStore(db).master()
	trans, err := master.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("An error occurred while deleting device. %v.", err))
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET deleted=:deleted WHERE user_id=:userId AND platform=:platform;", tableNameSubscription)
	params := map[string]interface{}{
		"userId":   userID,
		"platform": platform,
		"deleted":  time.Now().Unix(),
	}
	_, err = trans.Exec(query, params)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting device. %v.", err))
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE user_id=:userId AND platform=:platform;", tableNameDevice)
	params = map[string]interface{}{
		"userId":   userID,
		"platform": platform,
	}
	_, err = trans.Exec(query, params)
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting device. %v.", err))
		return err
	}

	err = trans.Commit()
	if err != nil {
		trans.Rollback()
		logger.Error(fmt.Sprintf("An error occurred while deleting device. %v.", err))
		return err
	}

	return nil
}

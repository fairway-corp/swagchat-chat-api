package datastore

import (
	"github.com/fairway-corp/swagchat-api/utils"
	gorp "gopkg.in/gorp.v2"
)

var (
	dbMap                   *gorp.DbMap
	TABLE_NAME_API          = utils.Cfg.Datastore.TableNamePrefix + "api"
	TABLE_NAME_USER         = utils.Cfg.Datastore.TableNamePrefix + "user"
	TABLE_NAME_ROOM         = utils.Cfg.Datastore.TableNamePrefix + "room"
	TABLE_NAME_ROOM_USER    = utils.Cfg.Datastore.TableNamePrefix + "room_user"
	TABLE_NAME_MESSAGE      = utils.Cfg.Datastore.TableNamePrefix + "message"
	TABLE_NAME_DEVICE       = utils.Cfg.Datastore.TableNamePrefix + "device"
	TABLE_NAME_SUBSCRIPTION = utils.Cfg.Datastore.TableNamePrefix + "subscription"
)

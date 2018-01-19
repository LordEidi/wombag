package wombag
/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017-18 by SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 ** LordLightningBolt@swordlord.com
 **
-----------------------------------------------------------------------------*/
import (
	"github.com/jinzhu/gorm"
	"log"
	"swordlord.com/wombag/model"
	"time"
)

var db gorm.DB

func InitDatabase() {

	dialect := GetStringFromConfig("db.dialect")
	args := GetStringFromConfig("db.args")

	database, err := gorm.Open(dialect, args)
	if err != nil {
		log.Fatalf("failed to connect database, %s", err)
		panic("failed to connect database")
	}

	gorm.DefaultCallback.Update().Register("update_upd_dat", updateCreated)

	db = *database

	db.SingularTable(true)
	db.LogMode(true)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Entry{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.EntryTag{})
}

func updateCreated(scope *gorm.Scope) {
	if scope.HasColumn("UpdDat") {
		scope.SetColumn("UpdDat", time.Now())
	}
}

func CloseDB() {

	db.Close()
}

func GetDB() *gorm.DB {

	return &db
}

package model
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
	"time"
	"github.com/jinzhu/gorm"
)

// TODO: Add CrtUsr and UpdUsr
type EntryTag struct {
	EntryTagId	uint `gorm:"primary_key"`
	EntryId	uint
	Entry Entry `gorm:"ForeignKey:entry_id"`
	TagId 	uint
	Tag Tag `gorm:"ForeignKey:tag_id"`
	CrtDat	time.Time `sql:"DEFAULT:current_timestamp"`
	UpdDat	time.Time `sql:"DEFAULT:current_timestamp"`
}

func (m *EntryTag) BeforeUpdate(scope *gorm.Scope) (err error) {

	scope.SetColumn("UpdDat", time.Now())
	return  nil
}




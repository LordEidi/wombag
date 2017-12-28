package model
/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017 by SwordLord - the coding crew - http://www.swordlord.com
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
)

type Device struct {
	Id    string `gorm:"primary_key"`
	Token 	string
	User	User `gorm:"ForeignKey:Name"`
	UserName string
	AccessToken	string
	ValidUntil	time.Time `sql:"DEFAULT:current_timestamp"`
	CrtDat	time.Time `sql:"DEFAULT:current_timestamp"`
	UpdDat	time.Time `sql:"DEFAULT:current_timestamp"`
}

/*
func (s *User) BeforeUpdate() (err error) {
	s.UpdDat = time.Now()
	return
}
*/

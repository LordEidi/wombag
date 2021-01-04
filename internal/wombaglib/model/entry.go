package model

/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017-20 by SwordLord - the coding crew - http://www.swordlord.com
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
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

// TODO: Add CrtUsr and UpdUsr
type Entry struct {
	EntryId        uint   `gorm:"primary_key"`
	Owner          string `sql:"NOT NULL;DEFAULT:admin;"`
	URL            string `sql:"NOT NULL;"`
	PreviewPic     string
	Domain         string
	Title          string `sql:"NOT NULL;"`
	Content        string `sql:"type:blob"`
	Language       string
	MimeType       string
	PreviewPicture string
	Tags           []Tag     `gorm:"-"`
	TagsAsString   string    `gorm:"-"`
	Starred        bool      `sql:"NOT NULL;DEFAULT:false"`
	Archived       bool      `sql:"NOT NULL;DEFAULT:false"`
	CrtDat         time.Time `sql:"DEFAULT:current_timestamp"`
	CrtUsr         string    `sql:"NOT NULL;DEFAULT:admin"`
	UpdDat         time.Time `sql:"DEFAULT:current_timestamp"`
	UpdUsr         string
}

func (m *Entry) BeforeUpdate(stx *gorm.DB) (err error) {

	m.UpdDat = time.Now()
	return nil
}

func (e Entry) GetContentJSON() string {

	b, err := json.Marshal(e.Content)
	if err != nil {
		return ""
	}
	// remove leading "
	jsonified := string(b)
	return strings.Trim(jsonified, "\"")

}

func (e Entry) GetTitleJSON() string {
	b, err := json.Marshal(e.Title)
	if err != nil {
		return ""
	}
	// remove leading "
	jsonified := string(b)
	return strings.Trim(jsonified, "\"")
}

func (e Entry) GetTags() string {

	// caching mechanism
	if e.TagsAsString != "" {

		if e.TagsAsString == "empty" {

			return ""
		} else {

			return e.TagsAsString
		}
	}

	//e.Tags = `{"id":12,"label":"dd","slug":"dd"}`

	ret := ""

	for _, tag := range e.Tags {

		if len(ret) > 0 {
			ret += ","
		}
		ret += fmt.Sprintf(`{"id":"%v","label":"%s","slug":"%s"}`, tag.TagId, tag.Label, tag.Slug)
	}

	return ret
}

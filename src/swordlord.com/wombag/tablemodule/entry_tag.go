package tablemodule
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
	"log"
	"swordlord.com/wombag"
	"swordlord.com/wombag/model"
)

func GetTagsPerEntry(entryId uint) []model.EntryTag {

	var rows []model.EntryTag

	query := wombag.GetDB()

	query.Where("Id=?", entryId)

	return rows
}

func AddTagsToEntry() {

	// tags 	string 	false 	tag1,tag2,tag3 	a comma-separated list of tags.
}

// TODO: function to remove tag from one entry, or from all (or some)
func DeleteTagPerEntry(entryID uint, tagID uint) {

	db := wombag.GetDB()

	retDB := db.Where("EntryID = ? AND TagID = ", entryID, tagID).Delete(model.EntryTag{})

	if retDB.Error != nil {
		log.Fatal(retDB.Error)
		return
	}

	log.Printf("Rows affected: %s\n", retDB.RowsAffected)
}

func DeleteEntryTag(tagId uint) {

	db := wombag.GetDB()

	retDB := db.Where("TagID = ", tagId).Delete(model.EntryTag{})

	if retDB.Error != nil {
		log.Fatal(retDB.Error)
		return
	}

	log.Printf("Rows affected: %s\n", retDB.RowsAffected)
}

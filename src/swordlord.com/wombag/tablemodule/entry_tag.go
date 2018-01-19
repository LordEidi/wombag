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
	"fmt"
	"log"
	"swordlord.com/wombag"
	"swordlord.com/wombag/model"
)

func GetTagsPerEntry(entryId uint) []model.Tag {

	var rows []model.Tag

	query := wombag.GetDB()

	rs := query.Joins("JOIN entry_tag on entry_tag.tag_id=tag.tag_id").Where("entry_tag.entry_id = ?", entryId).Group("tag.tag_id").Find(&rows)

	if rs.Error != nil {
		log.Printf("Error with Tags for Entry %q: %s\n", entryId, rs.Error)
		log.Fatal(rs.Error)
	}

	return rows
}

func AddTagToEntry(EntryId uint, tag string) (model.Tag, error) {

	db := wombag.GetDB()

	t := model.Tag{Slug: tag, Label: tag}

	retDB := db.Where(model.Tag{Slug: tag, Label: tag}).FirstOrInit(&t)
	//retDB := db.Create(&t)

	if retDB.Error != nil {
		log.Printf("Error with Tag %q: %s\n", tag, retDB.Error)
		log.Fatal(retDB.Error)
		return model.Tag{}, retDB.Error
	}

	db.Save(&t)

	fmt.Printf("Tag %s added.\n", tag)

	te := model.EntryTag{EntryId: EntryId, TagId: t.TagId}

	retDB2 := db.Where(model.EntryTag{EntryId: EntryId, TagId: t.TagId}).FirstOrInit(&te)
	//retDB2 := db.Create(&te)

	if retDB2.Error != nil {
		log.Printf("Error with EntryTag %q: %s\n", te, retDB2.Error)
		log.Fatal(retDB2.Error)
		return model.Tag{}, retDB2.Error
	}

	db.Save(&te)

	fmt.Printf("EntryTag Entry: %s Tag %s added.\n", EntryId, tag)

	return t, nil
}

// TODO: function to remove tag from one entry, or from all (or some)
func DeleteTagPerEntry(entryID uint, tagID uint) {

	db := wombag.GetDB()

	retDB := db.Where("entry_id = ? AND tag_id = ?", entryID, tagID).Delete(model.EntryTag{})

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

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

func ListTags() {

	db := wombag.GetDB()

	var rows []*model.Tag

	db.Find(&rows)

	var tags [][]string

	for _, tag := range rows {

		tags = append(tags, []string{tag.Id, tag.Slug, tag.Label, tag.CrtDat.Format("2006-01-02 15:04:05"), tag.UpdDat.Format("2006-01-02 15:04:05")})
	}

	wombag.WriteTable([]string{"Id", "Slug", "Label", "CrtDat", "UpdDat"}, tags)
}


func GetTagsTyped(filter *Filter) []model.Tag {

	var rows []model.Tag

	query := wombag.GetDB().Order("Label")

	query.Find(&rows)

	return rows
}

func AddTag(label string, slug string) (model.Tag, error) {

	db := wombag.GetDB()

	tag := model.Tag{Slug: slug, Label: label}
	retDB := db.Create(&tag)

	if retDB.Error != nil {
		log.Printf("Error with Tag %q: %s\n", tag, retDB.Error)
		log.Fatal(retDB.Error)
		return model.Tag{}, retDB.Error
	}

	fmt.Printf("Tag %s added.\n", label)

	return tag, nil
}

func UpdateTag(id uint, label string, slag string) {

	db := wombag.GetDB()

	retDB := db.Model(&model.Tag{}).Where("Id=?", id).Update("Label", label)

	if retDB.Error != nil {
		log.Printf("Error with Tag %q: %s\n", id, retDB.Error)
		log.Fatal(retDB.Error)
		return
	}

	fmt.Printf("Tag %s updated.\n", label)
}

// TODO: function to remove tag from one entry, or from all (or some)
func DeleteTag(id string) {

	db := wombag.GetDB()

	tag := &model.Tag{}

	retDB := db.Where("id = ?", id).First(&tag)

	if retDB.Error != nil {
		log.Printf("Error with Tag %q: %s\n", id, retDB.Error)
		log.Fatal(retDB.Error)
		return
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("Tag not found: %s\n", tag)
		log.Fatal("Tag not found: " + id + "\n")
		return
	}

	log.Printf("Deleting Tag: %s", &tag.Id)

	db.Delete(&tag)

	fmt.Printf("Tag %s deleted.\n", id)
}

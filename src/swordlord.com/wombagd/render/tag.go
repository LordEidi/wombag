package render

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
	"net/http"
	"swordlord.com/wombag/model"
	"text/template"
)

type TagsJSON struct {
	Tags []model.Tag
}

func (r TagsJSON) Render(w http.ResponseWriter) (err error) {

	if err = writeTagsJSON(w, r); err != nil {
		fmt.Printf("Error while rendering %v\n", err)
	}
	return
}

func (r TagsJSON) WriteHeader(w http.ResponseWriter) {
	writeHeader(w, jsonContentType)
}

func writeTagsJSON(w http.ResponseWriter, tags TagsJSON) error {

	var t *template.Template

	writeHeader(w, jsonContentType)

	t = template.Must(template.New("tags.tmpl").ParseFiles("./templates/tags.tmpl"))

	err := t.Execute(w, tags)

	if err != nil {
		return err
	}

	return nil
}

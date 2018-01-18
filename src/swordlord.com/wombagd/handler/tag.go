package handler
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
	"swordlord.com/wombag/tablemodule"
	"swordlord.com/wombagd/render"
	"swordlord.com/wombagd/respond"
)

func OnGetTagsForEntry(w http.ResponseWriter, req *http.Request){

	respond.NotImplementedYet(w)
}

func OnAddTagsToEntry(w http.ResponseWriter, req *http.Request){

	// POST /api/entries/{entry}/tags.{_format}
	// entry (int), tags (string) tag1,tag2,tag3

	var form QueryParams

	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: " + err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	// split tags if not empty
	// for each tag, check if exists, if not, add to tags, link to this entry

	respond.NotImplementedYet(w)
}

func OnDeleteTagOnEntriesBySlug(w http.ResponseWriter, req *http.Request){

	// DELETE /api/tag/label.{_format} -> tag (string)
	respond.NotImplementedYet(w)
}

func OnDeleteTagOnEntry(w http.ResponseWriter, req *http.Request){

	// DELETE /api/entries/{entry}/tags/{tag}.{_format} -> tag (int), entry (int)
	respond.NotImplementedYet(w)
}

func OnDeleteTagOnEntriesById(w http.ResponseWriter, req *http.Request){

	// tag by int
	respond.NotImplementedYet(w)
}

func OnRetrieveAllTags(w http.ResponseWriter, req *http.Request){

	form := tablemodule.NewFilter()
	// This will infer what binder to use depending on the content-type header.
	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: " + err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	tags := render.TagsJSON{}
	tags.Tags = tablemodule.GetTagsTyped(&form)

	respond.Render(w,http.StatusOK, tags)
}

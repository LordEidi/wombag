package handler

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
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"wombag/internal/wombaglib"
	"wombag/internal/wombaglib/render"
	"wombag/internal/wombaglib/tablemodule"
)

func OnGetTagsForEntry(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	sId := vars["entry"]

	EntryId, err := strconv.Atoi(sId)

	tags := render.TagsJSON{}

	if err == nil {
		tags.Tags = tablemodule.GetTagsPerEntry(uint(EntryId))
	}

	wombaglib.Render(w, http.StatusOK, tags)
}

func OnAddTagsToEntry(w http.ResponseWriter, req *http.Request) {

	// POST /api/entries/{entry}/tags.{_format}
	// entry (int), tags (string) tag1,tag2,tag3

	var form QueryParams

	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		wombaglib.WithMessage(w, http.StatusBadRequest, "An Error occured: "+err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		wombaglib.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	vars := mux.Vars(req)
	sId := vars["entry"]

	EntryId, err := strconv.Atoi(sId)

	// split tags if not empty
	// for each tag, check if exists, if not, add to tags, link to this entry

	tags := strings.Split(form.Tags, ",")
	tagsJson := render.TagsJSON{}

	for _, tag := range tags {

		a, err := tablemodule.AddTagToEntry(uint(EntryId), tag)

		if err == nil {
			tagsJson.Tags = append(tagsJson.Tags, a)
		}

	}

	wombaglib.Render(w, http.StatusOK, tagsJson)
}

func OnDeleteTagOnEntriesBySlug(w http.ResponseWriter, req *http.Request) {

	// DELETE /api/tag/label.{_format} -> tag (string)
	wombaglib.NotImplementedYet(w)
}

func OnDeleteTagOnEntry(w http.ResponseWriter, req *http.Request) {

	// DELETE /api/entries/{entry}/tags/{tag}.{_format} -> tag (int), entry (int)
	vars := mux.Vars(req)
	sEntryId := vars["entry"]
	sTagId := vars["tag"]

	entryId, _ := strconv.Atoi(sEntryId)
	tagId, _ := strconv.Atoi(sTagId)

	tablemodule.DeleteTagPerEntry(uint(entryId), uint(tagId))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Deleted EntryTag\n")
}

func OnDeleteTagOnEntriesById(w http.ResponseWriter, req *http.Request) {

	// tag by int
	wombaglib.NotImplementedYet(w)
}

func OnRetrieveAllTags(w http.ResponseWriter, req *http.Request) {

	form := tablemodule.NewFilter()
	// This will infer what binder to use depending on the content-type header.
	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		wombaglib.WithMessage(w, http.StatusBadRequest, "An Error occured: "+err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		wombaglib.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	tags := render.TagsJSON{}
	tags.Tags = tablemodule.GetTagsTyped(&form)

	wombaglib.Render(w, http.StatusOK, tags)
}

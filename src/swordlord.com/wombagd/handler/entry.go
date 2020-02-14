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
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"swordlord.com/wombag/model"
	"swordlord.com/wombag/tablemodule"
	"swordlord.com/wombagd/lib"
	"swordlord.com/wombagd/render"
	"swordlord.com/wombagd/respond"
)

func OnCreateEntry(w http.ResponseWriter, req *http.Request) {

	var form QueryParams

	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: "+err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	device := req.Context().Value(lib.AuthDevice).(model.Device)

	entry, err := tablemodule.AddEntry(device, form.Url)
	if err != nil {
		respond.WithMessage(w, http.StatusInternalServerError, "An Error occured: "+err.Error())
	}

	// TODO get correct entry from update...
	entryJSON := render.EntryJSON{}
	entryJSON.Entry = entry
	respond.Render(w, http.StatusOK, entryJSON)
}

func OnRetrieveEntries(w http.ResponseWriter, req *http.Request) {
	/*
		vars := mux.Vars(req)

		fmt.Fprintf(w, "Read: %v\n", vars["category"])
	*/

	form := tablemodule.NewFilter()

	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: "+err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	device := req.Context().Value(lib.AuthDevice).(model.Device)

	entries := render.EntriesJSON{}
	entries.SetEntries(tablemodule.GetEntriesTyped(device, &form))
	entries.Limit = form.PerPage
	entries.Page = form.Page

	respond.Render(w, http.StatusOK, entries)
}

func OnDeleteEntry(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	sId := vars["entry"]

	EntryId, err := strconv.Atoi(sId)

	if err == nil {
		tablemodule.DeleteEntry(uint(EntryId))
	}
}

func OnGetEntry(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	sId := vars["entry"]

	id, err := strconv.Atoi(sId)

	if err != nil {
		id = 0
	}

	device := req.Context().Value(lib.AuthDevice).(model.Device)

	entry := render.EntryJSON{}
	entry.Entry = tablemodule.GetEntryTyped(device, id)
	respond.Render(w, http.StatusOK, entry)
}

func OnChangeEntry(w http.ResponseWriter, req *http.Request) {

	var form QueryParams
	// This will infer what binder to use depending on the content-type header.
	err := bind(&form, req)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured: "+err.Error())
		return
	}

	isvalid := isValid(form)
	if !isvalid {
		respond.WithMessage(w, http.StatusBadRequest, "An Error occured")
		return
	}

	vars := mux.Vars(req)

	sId := vars["entry"]
	sTitle := form.Title

	device := req.Context().Value(lib.AuthDevice).(model.Device)

	// todo what if you are not authorised?
	tablemodule.UpdateEntry(device, sId, form.Starred != 0, form.Archive != 0, sTitle)

	id, err := strconv.Atoi(sId)

	if err != nil {
		id = 0
	}

	entry := render.EntryJSON{}
	// todo what if you are not authorised?
	entry.Entry = tablemodule.GetEntryTyped(device, id)
	respond.Render(w, http.StatusOK, entry)
}

func OnGetEntryFormatted(w http.ResponseWriter, req *http.Request) {

	respond.NotImplementedYet(w)
}

func OnReloadEntry(w http.ResponseWriter, req *http.Request) {

	respond.NotImplementedYet(w)
}

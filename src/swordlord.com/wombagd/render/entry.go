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
	"swordlord.com/wombag"
	"swordlord.com/wombag/model"
	"text/template"
)

type EntryJSON struct {
	Entry model.Entry
}

type EntriesJSON struct {
	Entries []model.Entry
	Page    int
	Size    int
	Limit   int
	Server  string
	Port    string
}

func (es *EntriesJSON) SetEntries(e []model.Entry) {

	es.Entries = make([]model.Entry, len(e))
	copy(es.Entries, e)
	es.Size = len(es.Entries)
}

func (es EntriesJSON) GetEntries() []model.Entry {

	return es.Entries
}

func (r EntryJSON) Render(w http.ResponseWriter) (err error) {

	if err = writeEntryJSON(w, r); err != nil {
		fmt.Printf("Error while rendering %v\n", err)
	}
	return
}

func (r EntriesJSON) Render(w http.ResponseWriter) (err error) {

	r.Server = wombag.GetStringFromConfig("www.host")
	r.Port = wombag.GetStringFromConfig("www.port")

	if err = writeEntriesJSON(w, r); err != nil {
		fmt.Printf("Error while rendering %v\n", err)
	}
	return
}

func (r EntryJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (r EntriesJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func writeEntryJSON(w http.ResponseWriter, entry EntryJSON) error {

	var t *template.Template

	writeContentType(w, jsonContentType)

	t = template.Must(template.New("entry.tmpl").ParseFiles("./templates/entry.tmpl"))

	err := t.Execute(w, entry)

	if err != nil {
		return err
	}

	return nil
}

func writeEntriesJSON(w http.ResponseWriter, entries EntriesJSON) error {

	var t *template.Template

	writeContentType(w, jsonContentType)

	t = template.Must(template.New("entries.tmpl").ParseFiles("./templates/entries.tmpl"))

	err := t.Execute(w, entries)

	if err != nil {
		return err
	}

	return nil
}

/*
func WriteJSON(w http.ResponseWriter, obj interface{}) error {


	var t *template.Template

	writeContentType(w, jsonContentType)

	switch v := obj.(type) {
		case int:
			// v is an int here, so e.g. v + 1 is possible.
			fmt.Printf("Integer: %v", v)
		case float64:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			fmt.Printf("Float64: %v", v)
		case string:
			// v is a string here, so e.g. v + " Yeah!" is possible.
			fmt.Printf("String: %v", v)
		default:
			// And here I'm feeling dumb. ;)
			fmt.Printf("I don't know, ask stackoverflow.")
	}

	t = template.Must(template.New("entry.tmpl").ParseFiles("./templates/entry.tmpl"))
	t = template.Must(template.New("entries.tmpl").ParseFiles("./templates/entries.tmpl"))

	err := t.Execute(w, obj)

	if err != nil {
		return err
	}

	return nil
}
*/

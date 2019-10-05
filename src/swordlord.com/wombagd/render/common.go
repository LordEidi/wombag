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
	"net/http"
)

//var jsonContentType = []string{"application/json; charset=utf-8"}
var jsonContentType = "application/json; charset=utf-8"

//var jsonContentType = "application/json"

//func writeHeader(w http.ResponseWriter, value []string) {
func writeHeader(w http.ResponseWriter, value string) {
	// TODO set or no replace?
	/*
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = value
		}
	*/
	w.Header().Set("Server", "Wombag")
	w.Header().Set("Content-Type", value)
}

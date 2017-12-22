package lib
/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017 by SwordLord - the coding crew - http://www.swordlord.com
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
	"github.com/gin-gonic/gin"
	"strconv"
	"swordlord.com/wombag/tablemodule"
	"swordlord.com/wombagd/render"
)

type QueryParams struct {

	Url			string	`form:"url" json:"url"` 		// Url for the entry.
	Title		string 	`form:"title" json:"title"` 	// Optional, we'll get the title from the page.
	Tags 		string 	`form:"tags" json:"tags"` 		// tag1,tag2,tag3 	a comma-separated list of tags.
	Starred 	int 	`form:"starred" json:"starred"` // 1 or 0 	entry already starred
	Archive 	int 	`form:"archive" json:"archive"` // 1 or 0 	entry already archived
}


func OnRoot(c *gin.Context){

	c.JSON(200, gin.H{"message": "welcome to the root"})
}

func OnRemoveAnnotation(c *gin.Context){

}

func OnUpdateAnnotation(c *gin.Context){

}

func OnRetrieveAnnotation(c *gin.Context){

}

func OnCreateNewAnnotation(c *gin.Context){

}

func OnCreateEntry(c *gin.Context){

	var form QueryParams
	// This will infer what binder to use depending on the content-type header.
	err := c.Bind(&form)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		c.JSON(300, gin.H{"An Error occured": err})
	}

	entry, err := tablemodule.AddEntry(form.Url)
	if err != nil {
		c.JSON(500, gin.H{"An Error occured": err})
	}

	// TODO get correct entry from update...
	entryJSON := render.EntryJSON{}
	entryJSON.Entry = entry
	c.Render(200, entryJSON)
}

func OnRetrieveEntries(c *gin.Context){

	filter := tablemodule.NewFilter()
	// This will infer what binder to use depending on the content-type header.
	err := c.Bind(&filter)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		c.JSON(300, gin.H{"error": err})
	}

	entries := render.EntriesJSON{}
	entries.Entries = tablemodule.GetEntriesTyped(&filter)
	entries.Limit = filter.PerPage
	entries.Page = filter.Page

	c.Render(200, entries)
}

func OnDeleteEntry(c *gin.Context){

	s_id := c.Param("entry")

	EntryId, err := strconv.Atoi(s_id)

	if err == nil {
		tablemodule.DeleteEntry(uint(EntryId))
	}
}

func OnGetEntry(c *gin.Context){

}

func OnChangeEntry(c *gin.Context){

	var form QueryParams
	// This will infer what binder to use depending on the content-type header.
	err := c.Bind(&form)

	if err != nil {
		fmt.Printf("Error when binding %v\n", err)
		c.JSON(300, gin.H{"error": err})
	}

	s_id := c.Param("entry")

	tablemodule.UpdateEntry(s_id, form.Starred != 0, form.Archive != 0)

	id, err := strconv.Atoi(s_id)

	if err != nil {
		id = 0
	}

	entry := render.EntryJSON{}
	entry.Entry = tablemodule.GetEntryTyped(id)
	c.Render(200, entry)
}

func OnGetEntryFormatted(c *gin.Context){

}

func OnReloadEntry(c *gin.Context){

}

func OnRetrieveTagsForEntry(c *gin.Context){

}

func OnAddTagsToEntry(c *gin.Context){

}

func OnDeleteTagsOnEntry(c *gin.Context){

}

func OnDeleteTagOnEntry(c *gin.Context){

}

func OnRetrieveAllTags(c *gin.Context){

}

func OnRemoveTagFromEveryEntry(c *gin.Context){

}

func OnRetrieveVersionNumber(c *gin.Context){

}

func OnOAuth(c *gin.Context){

	wtext := render.WombagText{}
	c.Render(200, wtext)
}
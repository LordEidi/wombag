package main
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
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/pprof"
	"swordlord.com/wombag"
	"swordlord.com/wombagd/lib"
)

func main() {

	//
	wombag.InitConfig()

	//
	wombag.InitDatabase()
	defer wombag.CloseDB()

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	//authorized := r.Group("/", lib.BasicAuth())

	/*
	authorized.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	*/

	r.GET("/", lib.OnRoot)

	api := r.Group("/api/", lib.ServiceOAuth())

	api.DELETE("/annotations/:annotation", lib.OnRemoveAnnotation) // DELETE
	api.PUT("/annotations/:annotation", lib.OnUpdateAnnotation) // PUT
	api.GET("/annotations/:annotation", lib.OnRetrieveAnnotation) // GET
	api.POST("/annotations/:entry", lib.OnCreateNewAnnotation) // POST
	api.GET("/entries/", lib.OnRetrieveEntries) // GET
	api.POST("/entries/", lib.OnCreateEntry) // POST
	api.DELETE("/entries/:entry", lib.OnDeleteEntry) // DELETE
	api.GET("/entries/:entry", lib.OnGetEntry) // GET
	api.PATCH("/entries/:entry", lib.OnChangeEntry) // PATCH
	api.GET("/entries/:entry/export", lib.OnGetEntryFormatted) // GET
	api.PATCH("/entries/:entry/reload", lib.OnReloadEntry) // PATCH
	api.GET("/entries/:entry/tags", lib.OnRetrieveTagsForEntry) // GET
	api.POST("/entries/:entry/tags", lib.OnAddTagsToEntry) // POST
	api.DELETE("/entries/:entry/tags/:tag", lib.OnDeleteTagsOnEntry) // DELETE
	api.DELETE("/tag/label", lib.OnDeleteTagOnEntry) // DELETE
	api.GET("/tags", lib.OnRetrieveAllTags) // GET
	//api.DELETE("/tags/label", lib.OnRemoveTagsFromEveryEntry) // DELETE
	api.DELETE("/tags/:tag", lib.OnRemoveTagFromEveryEntry) // DELETE
	api.GET("/version", lib.OnRetrieveVersionNumber) // GET

	r.POST("oauth/v2/token", lib.OnOAuth)

	//r.GET('/', onHitRoot);
	//r.bypassed.add(onBypass);

	// TODO in DEV Mode only
	r.GET("/debug/pprof/block", pprofHandler(pprof.Index))
	r.GET("/debug/pprof/heap", pprofHandler(pprof.Index))
	r.GET("/debug/pprof/profile", pprofHandler(pprof.Profile))
	r.POST("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
	r.GET("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
	r.GET("/debug/pprof/trace", pprofHandler(pprof.Trace))

	host := wombag.GetStringFromConfig("www.host")
	port := wombag.GetStringFromConfig("www.port")

	fmt.Printf("wombagd running on %v:%v\n", host, port)
	fmt.Printf("** add entry  : curl -X POST 'http://%s:%s/api/entries/' --data 'url=http://test' -H 'Content-Type:application/x-www-form-urlencoded'\n", host, port)
	fmt.Printf("** get entries: curl -X GET 'http://%s:%s/api/entries/?page=1&perPage=20'\n", host, port)
	fmt.Printf("** get entry  : curl -X GET 'http://%s:%s/api/entries/1'\n", host, port)
	fmt.Printf("** patch entry: curl -X PATCH 'http://%s:%s/api/entries/1' --data 'archive=1&starred=1' -H 'Content-Type:application/x-www-form-urlencoded'\n", host, port)

	r.Run(host + ":" + port) // listen and serve
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

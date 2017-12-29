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

	env := wombag.GetStringFromConfig("env")

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// what to do when a user hits the root
	r.GET("/", lib.OnRoot)

	// group to wrap all services which need authentication (authentication bearer http header)
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
	// this one does not like the one below, sine both have the same path, left here for completeness of the API
	// api.DELETE("/tags/label", lib.OnRemoveTagsFromEveryEntry) // DELETE
	api.DELETE("/tags/:tag", lib.OnRemoveTagFromEveryEntry) // DELETE
	api.GET("/version", lib.OnRetrieveVersionNumber) // GET

	// endpoint which is used to ask for a access token
	r.POST("oauth/v2/token", lib.OnOAuth)

	// TODO set a bypass function when no path triggers
	//r.GET('/', onHitRoot);
	//r.bypassed.add(onBypass);

	host := wombag.GetStringFromConfig("www.host")
	port := wombag.GetStringFromConfig("www.port")

	if env == "dev" {

		// give the user the possibility to trace and profile the app
		r.GET("/debug/pprof/block", pprofHandler(pprof.Index))
		r.GET("/debug/pprof/heap", pprofHandler(pprof.Index))
		r.GET("/debug/pprof/profile", pprofHandler(pprof.Profile))
		r.POST("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
		r.GET("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
		r.GET("/debug/pprof/trace", pprofHandler(pprof.Trace))

		// give the user some hints on what URLs she could test
		fmt.Printf("wombagd running on %v:%v\n", host, port)

		fmt.Printf("** get token  : curl -X POST 'http://%s:%s/oauth/v2/token' -F 'client_id=id' -F 'client_secret=secret' -F 'grant_type=password' -F 'password=pwd' -F 'username=uid' -H 'Content-Type:application/x-www-form-urlencoded'\n", host, port)
		fmt.Printf("** add entry  : curl -X POST 'http://%s:%s/api/entries/' --data 'url=http://test' -H 'Content-Type:application/x-www-form-urlencoded' -H 'Authorization: Bearer (access token)'\n", host, port)
		fmt.Printf("** get entries: curl -X GET 'http://%s:%s/api/entries/?page=1&perPage=20' -H 'Authorization: Bearer (access token)\n", host, port)
		fmt.Printf("** get entry  : curl -X GET 'http://%s:%s/api/entries/1' -H 'Authorization: Bearer (access token)\n", host, port)
		fmt.Printf("** patch entry: curl -X PATCH 'http://%s:%s/api/entries/1' --data 'archive=1&starred=1' -H 'Content-Type:application/x-www-form-urlencoded' -H 'Authorization: Bearer (access token)\n", host, port)

	}

	// have fun with wombagd
	r.Run(host + ":" + port)
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

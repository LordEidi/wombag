package main
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
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/negroni"
	"net/http"
	"swordlord.com/wombag"
	"swordlord.com/wombagd/handler"
	"swordlord.com/wombagd/lib"
)

func main() {

	// Initialise env and params
	wombag.InitConfig()

	// Initialise database
	// todo: make sure database is working as expected, chicken out otherwise
	wombag.InitDatabase()
	defer wombag.CloseDB()

	env := wombag.GetStringFromConfig("env")

	n := negroni.New(negroni.NewRecovery(), negroni.HandlerFunc(lib.OAuthMiddleware), negroni.NewLogger())

	gr := mux.NewRouter().StrictSlash(false)

	n.UseHandler(gr)

	// what to do when a user hits the root
	gr.HandleFunc("/", handler.OnRoot).Methods("GET")

	// annotations are very special, they are using the web authentication method vs oauth
	gr.HandleFunc("/annotations/{annotation}", handler.OnRemoveAnnotation).Methods("DELETE")
	gr.HandleFunc("/annotations/{annotation}", handler.OnUpdateAnnotation).Methods("PUT")
	gr.HandleFunc("/annotations/{annotation}", handler.OnRetrieveAnnotation).Methods("GET")
	gr.HandleFunc("/annotations/{annotation}", handler.OnCreateNewAnnotation).Methods("POST")

	gr.HandleFunc("/oauth/v2/token", handler.OnOAuth).Methods("POST")

	// and now group the
	api := gr.PathPrefix("/api").Subrouter()

	api.HandleFunc("/entries{ext:(?:.json)?}", handler.OnRetrieveEntries).Methods("GET")
	//api.HandleFunc("/entries.json", handler.OnRetrieveEntries).Methods("GET")
	api.HandleFunc("/entries{ext:(?:.json)?}", handler.OnCreateEntry).Methods("POST")
	//api.HandleFunc("/entries.json", handler.OnCreateEntry).Methods("POST")
	api.HandleFunc("/entries/{entry:[0-9]+}{ext:(?:.json)?}", handler.OnDeleteEntry).Methods("DELETE")
	api.HandleFunc("/entries/{entry:[0-9]+}{ext:(?:.json)?}", handler.OnGetEntry).Methods("GET")
	api.HandleFunc("/entries/{entry:[0-9]+}{ext:(?:.json)?}", handler.OnChangeEntry).Methods("PATCH")
	api.HandleFunc("/entries/{entry:[0-9]+}/export{ext:(?:.json)?}", handler.OnGetEntryFormatted).Methods("GET")
	api.HandleFunc("/entries/{entry:[0-9]+}/reload{ext:(?:.json)?}", handler.OnReloadEntry).Methods("PATCH")
	api.HandleFunc("/entries/{entry:[0-9]+}/tags{ext:(?:.json)?}", handler.OnGetTagsForEntry).Methods("GET")
	api.HandleFunc("/entries/{entry:[0-9]+}/tags{ext:(?:.json)?}", handler.OnAddTagsToEntry).Methods("POST")
	api.HandleFunc("/entries/{entry:[0-9]+}/tags/{tag:[0-9]+}{ext:(?:.json)?}", handler.OnDeleteTagOnEntry).Methods("DELETE")
	api.HandleFunc("/tag/label", handler.OnDeleteTagOnEntriesBySlug).Methods("DELETE")
	api.HandleFunc("/tags{ext:(?:.json|.txt|.xml)?}", handler.OnRetrieveAllTags).Methods("GET")
	//api.HandleFunc("/tags.json", handler.OnRetrieveAllTags).Methods("GET")
	// this one does not like the one below, sine both have the same path, left here for completeness of the API
	//api.HandleFunc("/tags/label", handler.OnRemoveTagsFromEveryEntry).Methods("DELETE")
	api.HandleFunc("/tags/:tag", handler.OnDeleteTagOnEntriesById).Methods("DELETE")
	api.HandleFunc("/version{ext:(?:.json|.txt|.xml|.html)?}", handler.OnRetrieveVersionNumber).Methods("GET")
	//api.HandleFunc("/version.html", handler.OnRetrieveVersionNumber).Methods("GET")
	//api.HandleFunc("/version.txt", handler.OnRetrieveVersionNumber).Methods("GET")

	api.HandleFunc("/entries/{entry:[0-9]+}{ext:(?:.json|.txt|.xml)?}", handler.OnGetEntry).Methods("GET")

	// TODO set a bypass function when no path triggers
	//r.GET('/', onHitRoot);
	//r.bypassed.add(onBypass);

	host := wombag.GetStringFromConfig("www.host")
	port := wombag.GetStringFromConfig("www.port")

	if env == "dev" {

		// give the user the possibility to trace and profile the app
		/*
		TODO RE ADD
		r.GET("/debug/pprof/block", pprofHandler(pprof.Index))
		r.GET("/debug/pprof/heap", pprofHandler(pprof.Index))
		r.GET("/debug/pprof/profile", pprofHandler(pprof.Profile))
		r.POST("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
		r.GET("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
		r.GET("/debug/pprof/trace", pprofHandler(pprof.Trace))
		*/

		// give the user some hints on what URLs she could test
		fmt.Printf("wombagd running on %v:%v\n", host, port)

		fmt.Printf("** get token  : curl -X POST 'http://%s:%s/oauth/v2/token' --data 'client_id=1&client_secret=secret&grant_type=password&password=pwd&username=uid' -H 'Content-Type:application/x-www-form-urlencoded'\n", host, port)
		fmt.Printf("** add entry  : curl -X POST 'http://%s:%s/api/entries/' --data 'url=http://test' -H 'Content-Type:application/x-www-form-urlencoded' -H 'Authorization: Bearer (access token)'\n", host, port)
		fmt.Printf("** get entries: curl -X GET 'http://%s:%s/api/entries/?page=1&perPage=20' -H 'Authorization: Bearer (access token)\n", host, port)
		fmt.Printf("** get entry  : curl -X GET 'http://%s:%s/api/entries/1' -H 'Authorization: Bearer (access token)\n", host, port)
		fmt.Printf("** patch entry: curl -X PATCH 'http://%s:%s/api/entries/1' --data 'archive=1&starred=1' -H 'Content-Type:application/x-www-form-urlencoded' -H 'Authorization: Bearer (access token)\n", host, port)

	}

	// have fun with wombagd
	http.ListenAndServe(host + ":" + port, n)
}

/*
func pprofHandler(h http.HandlerFunc) negroni.HandlerFunc {
handler := http.HandlerFunc(h)
return func(c *gin.Context) {
	handler.ServeHTTP(c.Writer, c.Request)
}

}
*/

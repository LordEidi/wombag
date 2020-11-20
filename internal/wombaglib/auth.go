package wombaglib

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
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
	"wombag/internal/wombaglib/tablemodule"
	"wombag/internal/wombaglib/util"
)

const AuthIsAuthenticated = "isauthenticated"
const AuthDevice = "authenticated_device"

// Middleware to check for Bearer Header
func OAuthMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	// this is a bit of a hack, since the mux did not behave as expected...
	if !strings.HasPrefix(req.URL.Path, "/api") {

		next(rw, req)
		return
	}

	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {

		log.Printf("There is neither an authorization header nor an access token in the URL.")
		http.Error(rw, "Access not authorised", 401)
		return
	}

	ahElements := strings.Split(authHeader, " ")
	if len(ahElements) != 2 {

		util.LogWarn("There is an authorization header but with wrong amount of elements.", nil)

		if util.IsDebuggingEnabled() {
			fields := logrus.Fields{"authHeader": authHeader}
			util.LogDebug("There is an authorization header but with wrong amount of elements.", fields)
		}
		http.Error(rw, "Access not authorised", 401)
		return
	}

	if ahElements[0] == "Bearer" {

		validateOAuthToken(ahElements[1], rw, req, next)

	} else {

		util.LogWarn("There is an authorization header but with wrong format.", nil)

		if util.IsDebuggingEnabled() {
			fields := logrus.Fields{"authHeader": authHeader}
			util.LogDebug("There is an authorization header but with wrong format.", fields)
		}
		http.Error(rw, "Access not authorised", 401)
		return
	}
}

func validateOAuthToken(accToken string, rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	device, err := tablemodule.ValidateAccessTokenInDB(accToken)

	if err == nil {

		ctx := req.Context()
		ctx = context.WithValue(ctx, AuthIsAuthenticated, true)
		ctx = context.WithValue(ctx, AuthDevice, device)

		// we need to forward the context
		next(rw, req.WithContext(ctx))
	} else {

		fields := logrus.Fields{"rip": req.RemoteAddr, "url": req.URL}
		util.LogWarn("Wrong AccessToken. Access denied.", fields)
		http.Error(rw, "Access not authorised", 401)
	}
}

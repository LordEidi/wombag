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
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"swordlord.com/wombag/tablemodule"
)

const AuthIsAuthenticated = "isauthenticated"
const AuthUser = "username"

// Middleware to check for Bearer Header
func ServiceOAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {

			// test the URL instead, this is a Hack for apps which think they should
			// put the access token into the URL...
			accessTokenUrl := c.Request.URL.Query().Get("access_token")
			if accessTokenUrl != "" {
				validateAccessToken(accessTokenUrl, c)
				return
			}

			log.Printf("There is neither an authorization header nor an access token in the URL.\n" )
			c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
			return
		}

		ahElements := strings.Split(authHeader, " ")
		if len(ahElements) != 2 {
			log.Printf("There is an authorization header but with wrong format: %s.\n", authHeader )
			c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
			return
		}

		if ahElements[0] == "Bearer" {

			validateAccessToken(ahElements[1], c)

		} else {

			log.Printf("There is an authorization header but with wrong format: %s.\n", authHeader )
			c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
		}
	}
}

func validateAccessToken(accToken string, c *gin.Context) {

	device, err := tablemodule.ValidateAccessTokenInDB(accToken)

	if err == nil {

		c.Set(AuthIsAuthenticated, true)
		c.Set(AuthUser, device.User)
		c.Next()
	} else {

		log.Printf("Wrong AccessToken. Access denied.\n" )
		c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
	}
}

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
			log.Printf("There is no authorization header.\n" )
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

			device, err := tablemodule.ValidateAccessTokenInDB(ahElements[1])

			if err != nil {

				c.Set(AuthIsAuthenticated, true)
				c.Set(AuthUser, device.User)
				c.Next()
			} else {

				log.Printf("Wrong AccessToken. Access denied.\n" )
				c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
			}
		} else {

			log.Printf("There is an authorization header but with wrong format: %s.\n", authHeader )
			c.AbortWithStatusJSON(401, gin.H{ "message": "Access not authorised"})
		}
	}
}

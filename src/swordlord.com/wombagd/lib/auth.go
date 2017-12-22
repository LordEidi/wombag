package lib

// Based on code Copyright 2014 Manu Martinez-Almeida.

import (
	//"crypto/subtle"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"swordlord.com/wombag/tablemodule"
)

const AuthUserName = "user"
const AuthIsAuthenticated = "isauthenticated"
const AuthPermissions = "permissions"

// BasicAuth returns a Basic HTTP Authorization middleware.
// (see http://tools.ietf.org/html/rfc2617#section-1.2)
func BasicAuth() gin.HandlerFunc{

	realm := "Basic realm=" + strconv.Quote("Oh Jasmin")

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		auth := strings.SplitN(authHeader, " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			c.AbortWithStatus(500)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO clean username and password before use
		username := pair[0]
		password := pair[1]

		// TODO Check user and Password against the database
		isAuthenticated := tablemodule.ValidateUserInDB(username, password)
		if isAuthenticated {

			// The user credentials was found, set user's id to key AuthUserKey in this context, the userId can be read later using
			// c.MustGet(gin.AuthUserKey)
			c.Set(AuthUserName, username)
			c.Set(AuthIsAuthenticated, true)
			return
		} else {

			c.Set(AuthIsAuthenticated, false)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "not authorised",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
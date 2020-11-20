package tablemodule

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
	"github.com/stretchr/testify/assert"
	"testing"
	"wombag/internal/wombaglib/util"
)

func TestAddOK(t *testing.T) {

	//
	util.InitConfig()

	//
	util.InitDatabase()
	defer util.CloseDB()

	// _, err := AddDevice("MyTestDevice", "Password", "TestUser")
	// ssert.NotNil(t, err, "There is an error creating a device for a not existing user")

	user, err := AddUser("TestUser", "Password")
	assert.Nil(t, err, "There is an error creating a device for a not existing user")
	assert.NotNil(t, user, "User should contain a new record")
	assert.Equal(t, user.Name, "TestUser", "Expecting new user to be named TestUser")

	d1, err := AddDevice("MyTestDevice", "Password", "User")
	assert.Nil(t, err, "There is an error creating a device for a not existing user")
	assert.NotNil(t, d1, "Expecting a new device")

	d2, err := ValidateDeviceInDB("MyTestDevice", "Password")
	assert.Nil(t, err, "There is an error authenticating the device")
	assert.NotNil(t, d2.AccessToken, "Expecting an AccessToken")
	assert.NotEmpty(t, d2.AccessToken, "Expecting an AccessToken")

	_, err1 := ValidateAccessTokenInDB(d2.AccessToken)
	assert.Nil(t, err1)

	_, err2 := ValidateAccessTokenInDB("complete_rubbish")
	assert.NotNil(t, err2)

	// cleaning up
	DeleteUser("TestUser")
	DeleteDevice("MyTestDevice")
}

package tablemodule

/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017-19 by SwordLord - the coding crew - http://www.swordlord.com
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
	uuid "github.com/satori/go.uuid"
	"log"
	"wombag/internal/wombaglib/model"
	"wombag/internal/wombaglib/util"
)

func ValidateDeviceInDB(deviceId, deviceToken string) (model.Device, error) {

	db := util.GetDB()

	device := model.Device{}

	token, err := hashPassword(deviceToken)
	if err != nil {
		log.Printf("Error with hashing password %q: %s\n", token, err)
		return model.Device{}, err
	}

	retDB := db.Where("id = ?", deviceId).First(&device)

	if retDB.Error != nil {
		log.Printf("Login of device failed %q: %s\n", deviceId, retDB.Error)
		return model.Device{}, retDB.Error
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("Login of device failed. Device not found: %s\n", deviceId)
		return model.Device{}, retDB.Error
	}

	err = checkHashedPassword(device.Token, deviceToken)
	if err != nil {
		log.Printf("Passwordhash missmatch for device: %s: %s\n", deviceId, err)
		return model.Device{}, err
	}

	u1 := uuid.NewV4()
	device.AccessToken = u1.String()

	u2 := uuid.NewV4()
	device.RefreshToken = u2.String()

	// TODO: set sensible validation time and validate it when authenticating
	device.ValidUntil = 3600

	updDB := db.Save(&device)
	if updDB.Error != nil {
		log.Printf("Updating of AccessToken failed %q: %s\n", deviceId, retDB.Error)
		return model.Device{}, retDB.Error
	}

	return device, nil
}

func ValidateAccessTokenInDB(accessToken string) (model.Device, error) {

	db := util.GetDB()

	device := model.Device{}

	retDB := db.Where("access_token = ?", accessToken).First(&device)

	if retDB.Error != nil {
		log.Printf("Login of device failed %s\n", retDB.Error)
		return model.Device{}, retDB.Error
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("Login of device failed.\n")
		return model.Device{}, retDB.Error
	}

	return device, nil
}

func ListDevice() {

	db := util.GetDB()

	var rows []*model.Device

	db.Find(&rows)

	var devices [][]string

	for _, device := range rows {

		devices = append(devices, []string{device.Id, device.Token, device.UserName, device.CrtDat.Format("2006-01-02 15:04:05"), device.UpdDat.Format("2006-01-02 15:04:05")})
	}

	util.WriteTable([]string{"Id", "Token", "User", "CrtDat", "UpdDat"}, devices)
}

func AddDevice(name string, password string, user string) (model.Device, error) {

	db := util.GetDB()

	pwd, err := hashPassword(password)
	if err != nil {
		log.Printf("Error with hashing password %q: %s\n", password, err)
		return model.Device{}, err
	}

	device := model.Device{Id: name, Token: pwd, UserName: user}
	retDB := db.Create(&device)

	if retDB.Error != nil {
		log.Printf("Error with Device %q: %s\n", name, retDB.Error)
		log.Fatal(retDB.Error)
		return model.Device{}, retDB.Error
	}

	fmt.Printf("Device %s for user %s added.\n", name, user)

	return device, nil
}

func UpdateDevice(name string, password string) error {

	db := util.GetDB()

	pwd, err := hashPassword(password)
	if err != nil {
		log.Printf("Error with hashing password %q: %s\n", password, err)
		return err
	}

	retDB := db.Model(&model.Device{}).Where("Id=?", name).Update("Token", pwd)

	if retDB.Error != nil {
		log.Printf("Error with Device %q: %s\n", name, retDB.Error)
		return retDB.Error
	}

	fmt.Printf("Device %s updated.\n", name)

	return nil
}

func DeleteDevice(name string) {

	db := util.GetDB()

	device := &model.Device{}

	retDB := db.Where("id = ?", name).First(&device)

	if retDB.Error != nil {
		log.Printf("Error with Device %q: %s\n", name, retDB.Error)
		log.Fatal(retDB.Error)
		return
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("Device not found: %s\n", name)
		log.Fatal("Device not found: " + name + "\n")
		return
	}

	log.Printf("Deleting Device: %s", &device.Id)

	db.Delete(&device)

	fmt.Printf("Device %s deleted.\n", name)
}

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
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"log"
	"wombag/internal/wombaglib/model"
	"wombag/internal/wombaglib/util"
)

// TODO return permission, not true/false. when empty permission, no access...
func ValidateUserInDB(name, password string) bool {

	permissions := []string{}
	// TODO get permissions from db
	permissions = append(permissions, "permission_from_DB")

	db := util.GetDB()

	user := &model.User{}

	retDB := db.Where("name = ? AND pwd = ?", name, password).First(&user)

	if retDB.Error != nil {
		fields := logrus.Fields{"name": name, "err": retDB.Error}
		util.LogFatal("Login of user failed.", fields)
		return false
	}

	if retDB.RowsAffected <= 0 {
		fields := logrus.Fields{"name": name}
		util.LogFatal("Login of user failed. User not found", fields)
		return false
	}

	hasAccess := true

	return hasAccess
}

func ListUser() {

	db := util.GetDB()

	var rows []*model.User

	db.Find(&rows)

	// Create
	//db.Create(&model.User{Name: "demo", Pwd: "demo"})

	//db.First(&user, "name = ?", "demo") // find product with id 1

	var users [][]string

	for _, user := range rows {

		users = append(users, []string{user.Name, user.Pwd, user.CrtDat.Format("2006-01-02 15:04:05"), user.UpdDat.Format("2006-01-02 15:04:05")})
	}

	util.WriteTable([]string{"Name", "Pwd", "CrtDat", "UpdDat"}, users)
}

func AddUser(name string, password string) (model.User, error) {

	db := util.GetDB()

	pwd, err := hashPassword(password)
	if err != nil {

		fields := logrus.Fields{"err": err}
		util.LogFatal("Error with hashing password.", fields)

		if util.IsDebuggingEnabled() {
			fields = logrus.Fields{"pwd": password}
			util.LogDebug("Password not hashable.", fields)
		}

		return model.User{}, err
	}

	user := model.User{Name: name, Pwd: pwd}
	retDB := db.Create(&user)

	if retDB.Error != nil {

		fields := logrus.Fields{"err": retDB.Error, "user": name}
		util.LogFatal("Error with user.", fields)
		return model.User{}, retDB.Error
	}

	fields := logrus.Fields{"user": name}
	util.LogInfo("user added.", fields)

	return user, nil
}

func UpdateUser(name string, password string) error {

	db := util.GetDB()

	pwd, err := hashPassword(password)
	if err != nil {
		log.Printf("Error with hashing password %q: %s\n", password, err)
		return err
	}

	retDB := db.Model(&model.User{}).Where("name=?", name).Update("pwd", pwd)

	if retDB.Error != nil {
		fields := logrus.Fields{"err": retDB.Error, "user": name}
		util.LogFatal("Error with user.", fields)
		return retDB.Error
	}

	fields := logrus.Fields{"user": name}
	util.LogInfo("user updated.", fields)

	return nil
}

func DeleteUser(name string) {

	db := util.GetDB()

	user := &model.User{}

	retDB := db.Where("name = ?", name).First(&user)

	if retDB.Error != nil {
		fields := logrus.Fields{"err": retDB.Error, "user": name}
		util.LogFatal("Error with user.", fields)
		return
	}

	if retDB.RowsAffected <= 0 {
		fields := logrus.Fields{"user": name}
		util.LogFatal("User not found.", fields)
		return
	}

	db.Delete(&user)

	fields := logrus.Fields{"user": name}
	util.LogInfo("user deleted.", fields)
}

func hashPassword(pwd string) (string, error) {

	password := []byte(pwd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

	// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match
}

func checkHashedPassword(hashedPassword string, password string) error {

	pwd := []byte(password)
	hashedPwd := []byte(hashedPassword)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)

	// nil means it is a match
	return err
}

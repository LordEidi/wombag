package tablemodule
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
	"log"
	"fmt"
	"swordlord.com/wombag/model"
	"swordlord.com/wombag"
)

// TODO return permission, not true/false. when empty permission, no access...
func ValidateUserInDB(name, password string) (bool) {

	permissions := []string{ }
	// TODO get permissions from db
	permissions = append(permissions, "permission_from_DB")

	db := wombag.GetDB()

	user := &model.User{}

	retDB := db.Where("name = ? AND pwd = ?", name, password).First(&user)

	if retDB.Error != nil {
		log.Printf("Login of user failed %q: %s\n", name, retDB.Error )
		log.Fatal(retDB.Error)
		return false
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("Login of user failed. User not found: %s\n", name)
		return false
	}

	hasAccess := true

	return hasAccess
}

func ensureUserTableExists(){

	/*
	CREATE TABLE user
	(
		name VARCHAR PRIMARY KEY,
		pwd VARCHAR NOT NULL,
		crt_dat DATETIME NOT NULL,
		upd_dat DATETIME
	);

	db, err := sql.Open("sqlite3", "./ohjasmin.db")
	if err != nil {
	log.Fatal(err)
	return false, permissions
	}
	defer db.Close()

	var count int

	// TODO get permissions from db
	permissions = append(permissions, "permission_from_DB")

	row := db.QueryRow("SELECT COUNT(*) FROM domain WHERE domain=? AND pwd=? ", user, password)
	err = row.Scan(&count)
	if err != nil {
	log.Fatal(err)
	return false, permissions
	}

	if count > 0 {
	return true, permissions
	}

	return false, permissions
	*/
}

func ListUser() {

	db := wombag.GetDB()

	var rows []*model.User

	db.Find(&rows)

	// Create
	//db.Create(&model.User{Name: "demo", Pwd: "demo"})

	//db.First(&user, "name = ?", "demo") // find product with id 1

	var users [][]string

	for _, user := range rows {

		users = append(users, []string{ user.Name, user.Pwd, user.CrtDat.Format("2006-01-02 15:04:05"), user.UpdDat.Format("2006-01-02 15:04:05")})
	}

	wombag.WriteTable([]string{"Name", "Pwd", "CrtDat", "UpdDat"}, users)
}

func AddUser(name string, pwd string) {

	db := wombag.GetDB()

	retDB := db.Create(&model.User{Name: name, Pwd: pwd})

	if retDB.Error != nil {
		log.Printf("Error with User %q: %s\n", name, retDB.Error )
		log.Fatal(retDB.Error)
		return
	}

	fmt.Printf("User %s added.\n", name)
}

func UpdateUser(name string, pwd string) {

	db := wombag.GetDB()

	retDB := db.Model(&model.User{}).Where("Name=?", name).Update("Pwd", pwd)

	if retDB.Error != nil {
		log.Printf("Error with User %q: %s\n", name, retDB.Error )
		log.Fatal(retDB.Error)
		return
	}

	fmt.Printf("User %s updated.\n", name)
}

func DeleteUser(name string) {

	db := wombag.GetDB()

	user := &model.User{}

	retDB := db.Where("name = ?", name).First(&user)

	if retDB.Error != nil {
		log.Printf("Error with User %q: %s\n", name, retDB.Error )
		log.Fatal(retDB.Error)
		return
	}

	if retDB.RowsAffected <= 0 {
		log.Printf("User not found: %s\n", name)
		log.Fatal("User not found: " + name + "\n")
		return
	}

	log.Printf("Deleting User: %s", &user.Name)

	db.Delete(&user)

	fmt.Printf("User %s deleted.\n", name)
}
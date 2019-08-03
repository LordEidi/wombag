package wombag

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
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

func InitConfig() {

	// Note: Viper does not require any initialization before using, unless we'll be dealing multiple different configurations.
	// check [working with multiple vipers](https://github.com/spf13/viper#working-with-multiple-vipers)

	// Set config file we want to read. 2 ways to do this.
	// 1. Set config file path including file name and extension
	//viper.SetConfigFile("./configs/wombag.config.json")

	// OR
	// 2. Register path to look for config files in. It can accept multiple paths.
	// It will search these paths in given order
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/config")
	// And then register config file name (no extension)
	viper.SetConfigName("wombag.config")
	// Optionally we can set specific config type
	viper.SetConfigType("json")

	// viper allows watching of config files for changes (and potential reloads)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {

		// TODO: don't just overwrite, check for existence first, then write a standard config file and move on...
		WriteStandardConfig()

		LogFatal("Error reading config file. Standard config file written.", logrus.Fields{"error": err, "file": "wombag.config.json"})
	}

	// Confirm which config file is used
	LogDebug("Loaded config.", logrus.Fields{"file": viper.ConfigFileUsed()})
}

func GetStringFromConfig(key string) string {

	return viper.GetString(key)
}

func GetLogLevel() string {

	return viper.GetString("env")
}

//
func WriteStandardConfig() error {

	err := ioutil.WriteFile("wombag.config.json", defaultConfig, 0700)

	return err
}

var defaultConfig = []byte(`{
	"env": "dev",
	"add_demo_users": "true",
	"www": {
		"host": "127.0.0.1",
		"port": "8081"
	},
	"db": {
		"dialect": "sqlite3",
		"args": "wombag.db"
	},
	"templates": {
		"dir": "./templates/",
		"auth": "auth.tmpl",
		"entries": "entries.tmpl",
		"entry": "entry.tmpl",
		"tags": "tags.tmpl"
	}
}
`)

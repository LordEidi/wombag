package wombag
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
	"github.com/spf13/viper"
	"log"
	"fmt"
)

var env string

func InitConfig() {

	// Note: Viper does not require any initialization before using, unless we'll be dealing multiple different configurations.
	// check [working with multiple vipers](https://github.com/spf13/viper#working-with-multiple-vipers)

	// Set config file we want to read. 2 ways to do this.
	// 1. Set config file path including file name and extension
	//viper.SetConfigFile("./configs/config.json")

	// OR
	// 2. Register path to look for config files in. It can accept multiple paths.
	// It will search these paths in given order
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/config")
	// And then register config file name (no extension)
	viper.SetConfigName("config")
	// Optionally we can set specific config type
	viper.SetConfigType("json")

	// viper allows watching of config files for changes (and potential reloads)
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	// })

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		// TODO: don't crash, write a standard config file and move on...
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	env = viper.GetString("env")

	// Confirm which config file is used
	fmt.Printf("Env set to: %s\n", env)
}

func GetStringFromConfig(key string) string {

	return viper.GetString(key)
}

func GetEnv() string {

	return env
}


package wombag
/*-----------------------------------------------------------------------------
 **
 ** - Wombag -
 **
 ** the alternative, native backend for your Wallabag apps
 **
 ** Copyright 2017-18 by SwordLord - the coding crew - http://www.swordlord.com
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
	"io/ioutil"
	"log"
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
	viper.SetConfigName("wombag.config")
	// Optionally we can set specific config type
	viper.SetConfigType("json")

	// viper allows watching of config files for changes (and potential reloads)
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	// })

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {

		// TODO: don't just overwrite, check for existence first, then write a standard config file and move on...
		WriteStandardConfig()

		if err := viper.ReadInConfig(); err != nil {
			// we tried it once, crash now
				log.Fatalf("Error reading config file, %s", err)
		}
	}

	// Confirm which config file is used
	// fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	env = viper.GetString("env")

	// Confirm which config file is used
	// fmt.Printf("Env set to: %s\n", env)

	EnsureTemplateFilesExist()
}

func GetStringFromConfig(key string) string {

	return viper.GetString(key)
}

func GetEnv() string {

	return env
}

//
func WriteStandardConfig() (error) {

	err := ioutil.WriteFile("wombag.config.json", defaultConfig, 0700)

	return err
}

var defaultConfig = []byte("{\n" +
	"\t\"env\": \"dev\",\n" +
	"\t\"add_demo_users\": \"true\",\n" +
	"\t\"www\": {\n" +
	"\t\t\"host\": \"0.0.0.0\",\n" +
	"\t\t\"port\": \"8081\"\n" +
	"\t},\n" +
	"\t\"db\": {\n" +
	"\t\t\"dialect\": \"sqlite3\",\n" +
	"\t\t\"args\": \"wombag.db\"\n" +
	"\t},\n" +
	"\t\"templates\": {\n" +
	"\t\t\"auth\": \"./templates/auth.tmpl\",\n" +
	"\t\t\"entries\": \"./templates/entries.tmpl\",\n" +
	"\t\t\"entry\": \"./templates/entry.tmpl\",\n" +
	"\t\t\"tags\": \"./templates/tags.tmpl\"\n" +
	"\t}\n" +
	"}")

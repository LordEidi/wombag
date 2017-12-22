package main

import (
	"fmt"
	"os"

	"swordlord.com/gohjasmincli/cmd"
	"swordlord.com/gohjasmin"
)

func main() {

	//
	gohjasmin.InitConfig()

	//
	gohjasmin.InitDatabase()
	defer gohjasmin.CloseDB()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		gohjasmin.CloseDB()
		os.Exit(1)
	}
}
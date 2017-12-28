package main

import (
	"fmt"
	"os"
	"swordlord.com/wombag"
	"swordlord.com/wombagcli/cmd"
)

func main() {

	//
	wombag.InitConfig()

	//
	wombag.InitDatabase()
	defer wombag.CloseDB()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		wombag.CloseDB()
		os.Exit(1)
	}
}
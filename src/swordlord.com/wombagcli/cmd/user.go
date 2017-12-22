// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"swordlord.com/gohjasmin/tablemodule"
)

// domainCmd represents the domain command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Add, change and manage your users.",
	Long: `Add, change and manage your users. Requires a subcommand.`,
	RunE: nil,
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users.",
	Long: `List all users.`,
	RunE: ListUser,
}

var userAddCmd = &cobra.Command{
	Use:   "add [user] [password]",
	Short: "Add new user to this instance of gohjasmin.",
	Long: `Add new user to this instance of gohjasmin.`,
	RunE: AddUser,
}

var userUpdateCmd = &cobra.Command{
	Use:   "update [user] [password]",
	Short: "Update the password of the user.",
	Long: `Update the password of the user.`,
	RunE: UpdateUser,
}

var userDeleteCmd = &cobra.Command{
	Use:   "delete [user]",
	Short: "Deletes a user.",
	Long: `Deletes a user.`,
	RunE: DeleteUser,
}

func ListUser(cmd *cobra.Command, args []string) error {

	tablemodule.ListUser()

	return nil
}

func AddUser(cmd *cobra.Command, args []string) error {

	if len(args) < 2 {
		er("command 'add' needs a user and a password")
	} else {
		tablemodule.AddUser(args[0], args[1])
	}

	return nil
}

func UpdateUser(cmd *cobra.Command, args []string) error {

	if len(args) < 2 {
		er("command 'update' needs a user and a new password")
	} else {
		tablemodule.UpdateUser(args[0], args[1])
	}

	return nil
}

func DeleteUser(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		er("command 'delete' needs a user")
	} else {
		tablemodule.DeleteUser(args[0])
	}

	return nil
}

func init() {
	RootCmd.AddCommand(userCmd)

	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userUpdateCmd)
	userCmd.AddCommand(userDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

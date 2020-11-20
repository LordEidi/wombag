package command

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
	"fmt"
	"strconv"
	"wombag/internal/wombaglib/model"
	"wombag/internal/wombaglib/tablemodule"

	"github.com/spf13/cobra"
)

// entryCmd represents the domain command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "With the entry command you can add, change and manage your entries.",
	Long:  `With the entry command you can add, change and manage your entries.`,
	RunE:  nil,
}

var entryListCmd = &cobra.Command{
	Use:   "list [user]",
	Short: "List all entries (without details, filter for optional user).",
	Long:  `List all entries (without details, filter for optional user).`,
	RunE:  ListEntry,
}

var entryAddCmd = &cobra.Command{
	Use:   "add [user] [url]",
	Short: "Add a new entry. Wombag will instantly get details from the given URL.",
	Long:  `Add a new entry. Wombag will instantly get details from the given URL and store the entry to the given users first device.`,
	Args:  cobra.ExactArgs(2),
	RunE:  AddEntry,
}

var entryDeleteCmd = &cobra.Command{
	Use:   "delete [entry]",
	Short: "Deletes an entry.",
	Long:  `Deletes an entry.`,
	Args:  cobra.ExactArgs(1),
	RunE:  DeleteEntry,
}

func ListEntry(cmd *cobra.Command, args []string) error {

	// todo if user is empty, run as admin and show all. if given, filter for that user

	adminDevice := model.GetAdminDevice()

	tablemodule.ListEntries(adminDevice)

	return nil
}

func AddEntry(cmd *cobra.Command, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("command 'add' needs an URL")
	}

	device := model.GetAdminDevice()

	// todo add entry for that user
	tablemodule.AddEntry(device, args[0])

	return nil
}

func DeleteEntry(cmd *cobra.Command, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("command 'delete' needs an ID")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("entry ID is not a number")
	}

	tablemodule.DeleteEntry(uint(id))

	return nil
}

func init() {
	RootCmd.AddCommand(entryCmd)

	entryCmd.AddCommand(entryAddCmd)
	entryCmd.AddCommand(entryListCmd)
	entryCmd.AddCommand(entryDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

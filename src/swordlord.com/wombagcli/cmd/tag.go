package cmd

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
	"github.com/spf13/cobra"
	"strconv"
	"swordlord.com/wombag/tablemodule"
)

// domainCmd represents the domain command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add, change and manage devices of your users.",
	Long: `Add, change and manage devices of your users. Requires a subcommand.`,
	RunE: nil,
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all devices.",
	Long: `List all devices.`,
	RunE: ListTags,
}

var tagAddCmd = &cobra.Command{
	Use:   "add [device] [password] [user]",
	Short: "Add new device to given user.",
	Long: `Add new device to given user.`,
	RunE: AddTag,
}

var tagUpdateCmd = &cobra.Command{
	Use:   "update [device] [password]",
	Short: "Update the password of the device.",
	Long: `Update the password of the device.`,
	RunE: UpdateTag,
}

var tagDeleteCmd = &cobra.Command{
	Use:   "delete [device]",
	Short: "Deletes a device.",
	Long: `Deletes a device.`,
	RunE: DeleteTag,
}

func ListTags(cmd *cobra.Command, args []string) error {

	tablemodule.ListTags()

	return nil
}

func AddTag(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		er("command 'add' needs a tag label and slug")
	} else {
		tablemodule.AddTag(args[0], args[1])
	}

	return nil
}

func UpdateTag(cmd *cobra.Command, args []string) error {

	if len(args) != 3 {
		er("command 'update' needs a tag id, a new label and new slug")
	} else {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			return err
		}

		tablemodule.UpdateTag(uint(id), args[1], args[2])
	}

	return nil
}

func DeleteTag(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		er("command 'delete' needs a tag Id")
	} else {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			return err
		}

		tablemodule.DeleteTag(uint(id))
	}

	return nil
}

func init() {
	RootCmd.AddCommand(tagCmd)

	deviceCmd.AddCommand(tagListCmd)
	deviceCmd.AddCommand(tagAddCmd)
	deviceCmd.AddCommand(tagUpdateCmd)
	deviceCmd.AddCommand(tagDeleteCmd)
}

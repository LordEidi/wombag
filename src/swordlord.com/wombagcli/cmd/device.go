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
	"fmt"
	"github.com/spf13/cobra"
	"swordlord.com/wombag/tablemodule"
)

// deviceCmd to manage devices
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Add, change and delete devices of your users.",
	Long: `Add, change and delete devices of your users. 
Requires a subcommand.`,
	RunE: nil,
}

var deviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all devices.",
	Long: `List all devices.`,
	RunE: ListDevices,
}

var deviceAddCmd = &cobra.Command{
	Use:   "add [device] [devicetoken] [user]",
	Short: "Adds new device for a user.",
	Long: `Adds new device for a user. Device can be used immediately.`,
	Args: cobra.ExactArgs(3),
	RunE: AddDevice,
}

var deviceUpdateCmd = &cobra.Command{
	Use:   "update [device] [devicetoken]",
	Short: "Updates the password of the device.",
	Long: `Updates the password of the device.`,
	Args: cobra.ExactArgs(2),
	RunE: UpdateDevice,
}

var deviceDeleteCmd = &cobra.Command{
	Use:   "delete [device]",
	Short: "Deletes a device.",
	Long: `Deletes a device.`,
	Args: cobra.ExactArgs(1),
	RunE: DeleteDevice,
}

func ListDevices(cmd *cobra.Command, args []string) error {

	tablemodule.ListDevice()

	return nil
}

func AddDevice(cmd *cobra.Command, args []string) error {

	if len(args) != 3 {
		return fmt.Errorf("command 'add' needs a device, devicetoken and user")
	}

	_, err := tablemodule.AddDevice(args[0], args[1], args[2])

	return err
}

func UpdateDevice(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		return fmt.Errorf("command 'update' needs a device and a new password")
	}

	err := tablemodule.UpdateDevice(args[0], args[1])

	return err
}

func DeleteDevice(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("command 'delete' needs a device Id")
	}

	tablemodule.DeleteDevice(args[0])

	return nil
}

func init() {
	RootCmd.AddCommand(deviceCmd)

	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceAddCmd)
	deviceCmd.AddCommand(deviceUpdateCmd)
	deviceCmd.AddCommand(deviceDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

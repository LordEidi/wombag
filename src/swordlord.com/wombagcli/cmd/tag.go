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
	"strconv"
	"swordlord.com/wombag/tablemodule"
)

// tagCmd, to manage tags
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add, change and manage tags either on entries or on their own.",
	Long: `Add, change and manage tags either on entries or on their own. 
Requires a subcommand`,
	RunE: nil,
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tags.",
	Long: `List all tags.`,
	RunE: ListTags,
}

var tagAddCmd = &cobra.Command{
	Use:   "add [tag] [slug]",
	Short: "Add new tag.",
	Long: `Add new tag.`,
	Args: cobra.ExactArgs(2),
	RunE: AddTag,
}

var tagAddToEntryCmd = &cobra.Command{
	Use:   "addtagtoentry [tag] [entry]",
	Short: "Add new tag to an entry.",
	Long: `Add new tag to an entry.`,
	Args: cobra.ExactArgs(2),
	RunE: AddTagToEntry,
}

var tagDeleteCmd = &cobra.Command{
	Use:   "delete [tag]",
	Short: "Deletes a tag from all entries.",
	Long: `Deletes a tag from all entries.`,
	Args: cobra.ExactArgs(1),
	RunE: DeleteTag,
}

func ListTags(cmd *cobra.Command, args []string) error {

	tablemodule.ListTags()

	return nil
}

func AddTag(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		return fmt.Errorf("command 'add' needs a tag label and slug")
	}

	tablemodule.AddTag(args[0], args[1])

	return nil
}

func AddTagToEntry(cmd *cobra.Command, args []string) error {

	if len(args) != 3 {
		return fmt.Errorf("command 'addtagtoentry' needs a tag and an entry")
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
		return fmt.Errorf("command 'delete' needs a tag Id")
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

	tagCmd.AddCommand(tagListCmd)
	tagCmd.AddCommand(tagAddCmd)
	tagCmd.AddCommand(tagAddToEntryCmd)
	tagCmd.AddCommand(tagDeleteCmd)
}

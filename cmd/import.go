/*
Copyright © 2022 Paul Norman <osm@paulnorman.ca>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "load into an empty database",
	Long:  `Loads data into an empty changeset database`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		fmt.Printf("import called on %s\n", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		createTables()
		importData(file)
	},
}

func createTables() {
	userSQL := fmt.Sprintf(`CREATE TABLE %s_users AS (
  id bigint PRIMARY KEY,
  name text NOT NULL);`, tablePrefix)

	changesetSQL := fmt.Sprintf(`CREATE TABLE %s_changesets AS (
  id bigint NOT NULL, -- defer index creation until post-import
  user_id bigint NOT NULL,
  created_at timestamptz NOT NULL,
  closed_at timestamptz,
  num_changes integer NOT NULL);`, tablePrefix)

	discussionSQL := fmt.Sprintf(`CREATE TABLE %s_discussion (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  created_at timestamptz NOT NULL,
  discussion text NOT NULL);`, tablePrefix)

	stateSQL := fmt.Sprintf(`CREATE TABLE %s_state (
  sequence bigint PRIMARY KEY
);`, tablePrefix)
	// Print the SQL instead of running it, for now
	fmt.Println(userSQL)
	fmt.Println(changesetSQL)
	fmt.Println(discussionSQL)
	fmt.Println(stateSQL)
}

func importData(file *os.File) {
	fmt.Println("Do stuff")
	parseOsm(file)
}

func init() {
	rootCmd.AddCommand(importCmd)

	// replication URI
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global variables related to the DB
	tablePrefix  string
	databaseName string
	databaseUser string
	databaseHost string
	// TODO: Change to integer type
	databasePort string

	rootCmd = &cobra.Command{
		Use:   "gochange",
		Short: "Loads OpenStreetMap changesets",
		Long:  `Loads OpenStreetChangeset metadata into a PostgreSQL database`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Global database options
	rootCmd.PersistentFlags().StringVarP(&databaseName, "database", "d", "", "database name")
	rootCmd.PersistentFlags().StringVarP(&databaseUser, "user", "U", "", "database user")
	rootCmd.PersistentFlags().StringVarP(&databaseHost, "host", "H", "", "database host")
	rootCmd.PersistentFlags().StringVarP(&databasePort, "port", "P", "", "database port")
	rootCmd.PersistentFlags().StringVarP(&tablePrefix, "prefix", "p", "osm_changes", "prefix for tables")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

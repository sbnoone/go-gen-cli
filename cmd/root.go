package cmd

import (
	"fmt"
	"gen/entities"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

var newCmd = &cobra.Command{
	Use:   "new [command]",
	Short: "Creates template files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run new...")
	},
}

var config bool
var types bool
var cmdComponent = &cobra.Command{
	Use:     "component [file name]",
	Short:   "Creates a component",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		component := entities.Component{
			FileName: args[0],
			Config:   config,
			Types:    types,
		}
		component.Create()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// var cmdRemove = &cobra.Command{
	// 	Use:   "remove [file name]",
	// 	Short: "Removes a file",
	// 	Long:  `Removes a file with specified name`,
	// 	Args:  cobra.MinimumNArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		removeFile(args[0])
	// 	},
	// }

	componentFlags := cmdComponent.Flags()
	componentFlags.BoolVarP(&config, "config", "c", false, "add a config file")
	componentFlags.BoolVarP(&types, "types", "t", false, "add a types file")

	newCmd.AddCommand(cmdComponent)
	rootCmd.AddCommand(newCmd)
}

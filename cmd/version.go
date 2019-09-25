package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gopherType string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get verison information",
	Long: `Get information about 
	cli tools verison.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		fmt.Println(args)
		fmt.Println(gopherType)
	},
}

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	versionCmd.Flags().StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	versionCmd.Flags().StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")

	rootCmd.AddCommand(versionCmd)
}

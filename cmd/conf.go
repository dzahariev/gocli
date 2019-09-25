package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Read conf",
	Long:  "Read the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("conf called")
		initConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}

func setDefaults() {
	viper.SetDefault("namevalue", "default")
}

func initConfiguration() {
	setDefaults()
	viper.SetConfigName("configuration")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found.")
		} else {
			panic(fmt.Errorf("fatal error config file: %s ", err))
		}
	}
	name := viper.Get("name")
	fmt.Printf("name: %s \n", name)
}

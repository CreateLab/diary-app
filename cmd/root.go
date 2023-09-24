package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diary",
	Short: "description of an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// настроить автозаполнение

func Execute() {
	err := rootCmd.Execute() // корневая команда кобры
	if err != nil {
		os.Exit(1) //если ошибка
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.diary-app.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

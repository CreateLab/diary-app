package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func viewNote() { // объявление функции для просмотра заметок
	if len(list) == 0 { //если заметок 0, показать сообщение (а их 0, их только через файл или бд сохранять)
		fmt.Println("empty:(")
		return
	}

	for _, note := range list {
		fmt.Printf("Title: %s", note.Title) //сделать показ по названиям
		fmt.Println("Date and time: ", note.CreatedAt)
		fmt.Println("Text:", note.Text)
		fmt.Println() // для разделения
	}
}

var viewCmd = &cobra.Command{ // аналогично create
	Use:   "view",
	Short: "View notes",
	Long:  "View all notes in the daily planner",
	Run: func(cmd *cobra.Command, args []string) {
		viewNote()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

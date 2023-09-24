package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// создание структуры Note с 3-мя полями: создано в (при помощи пакета тайм), заголовок и текст заметки
type Note struct {
	CreatedAt time.Time
	Title     string
	Text      string
}

// объявление переменной list типа срез, куда потом с помощью append будут добавляться заметки Note
var list []Note

// объявление функции, которая принимает из консоли значения заголовок и текст от пользователя
// при помощи пакета bufio, stdin для инпута
func addNote() {
	reader := bufio.NewReader(os.Stdin) // стандартный ввод

	fmt.Print("Enter a Title: ")        // в консоли просим ввести заголовок
	title, _ := reader.ReadString('\n') // инпут из консоли, считываем до enter (\n), возвращается строка + \n
	title = title[:len(title)-1]        // удаляем из строки \n

	fmt.Print("Enter a Text: ") // то же самое с текстом
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	note := Note{ // переменная, куда добавляются данные из инпута
		CreatedAt: time.Now(), // текущее время
		Title:     title,
		Text:      text,
	}

	list = append(list, note) // добавляем к срезу list заметку note
	fmt.Println("added!")     // выводим на экран сообщение
}

var createCmd = &cobra.Command{ // переменная (команда), описание и функция, которая ее вызывает
	Use:   "create",
	Short: "create a note",
	Run: func(cmd *cobra.Command, args []string) {
		addNote()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

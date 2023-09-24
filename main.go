package main

import (
	"fmt"

	"github.com/zalynskaya/diary-app.git/cmd"
)

func main() {
	defer func() {
	if err := recover(); err != nil {
	fmt.Println("recovered from panic:", err) //обработка ошибки
	}
	}()
	
	cmd.Execute()
	}
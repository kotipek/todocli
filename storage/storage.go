package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"todocli/models"
)

func WriteTasks(tasks []models.Task) {
	fmt.Println("11111111111111111111111111111111111")
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	taskJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при конвертации в JSON:", err)
		return
	}
	defer file.Close()

	file.WriteString(string(taskJSON))
	fmt.Println(string(taskJSON))

}

func LoadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}

package storage

import (
	"encoding/json"
	"fmt"
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

// func LoadTasks(tasks []models.Task) {
// 	file, err := os.Open("tasks.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// }

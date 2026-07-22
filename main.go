package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todocli/models"
	"todocli/storage"
)

var tasks []models.Task

func main() {

	var actionType string
	if len(os.Args) > 1 {
		actionType = os.Args[1]
		switch actionType {
		case "Создать":
			tasks = append(tasks, createTask())
			task := createTask()
			fmt.Println("таска - ", task)
		default:
			fmt.Println("Вы не выбрали действие")
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Введите действие - ")
			text, _ := reader.ReadString('\n')
			actionType = text
			actionType = strings.TrimSpace(text)
			fmt.Println(actionType)
			switch actionType {
			case "help":
				fmt.Println("Список действий:")
				fmt.Println("Создать задачу - Создать")
				fmt.Println("Вывести задачи - Вывести")
				fmt.Println("Изменить задачу - Изменить")
				fmt.Println("Удалить задачу - Удалить")
				fmt.Print("Введите действие - ")
			case "Создать":
				tasks = append(tasks, createTask())
				storage.WriteTasks(tasks)
			// case "Вывести":
			// 	getTasks()
			default:
				fmt.Println("Введите help для подсказки")
			}
		}
	}
}

func createTask() models.Task {
	return models.Task{
		Name:        "name1",
		Status:      models.StatusCompleted,
		Date:        time.Now(),
		Description: "gjrbnrnbgjr"}
}

// func getTasks() {
// 	for i := range tasks {
// 		fmt.Println("Название - ", tasks[i].Name,
// 			"Статус - ", tasks[i].Status,
// 			"Дата - ", tasks[i].Date,
// 			"Описание - ", tasks[i].Description,
// 			"Номер задачи - ", i+1)
// 	}
// }

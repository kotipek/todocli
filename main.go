package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
			// tasks = append(tasks, createTask())
			// task := createTask()
			// fmt.Println("таска - ", task)
		default:
			fmt.Println("Вы не выбрали действие")
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Введите действие - ")
			text, _ := reader.ReadString('\n')
			parts := strings.SplitN(text, " ", 3)
			// actionType = text
			actionType = strings.TrimSpace(parts[0])
			fmt.Println(actionType)
			switch actionType {
			case "help":
				fmt.Println("Список действий:")
				fmt.Println("Создать задачу - Создать <статус> <описание>")
				fmt.Println("Вывести задачи - Вывести")
				fmt.Println("Изменить задачу - Изменить")
				fmt.Println("Удалить задачу - Удалить")
			case "Создать":
				if len(parts) < 3 {
					fmt.Println("Создать <статус> <описание>")
					continue
				}
				task, err := createTask(parts[1], parts[2])
				if err != nil {
					fmt.Println(err)
					continue
				}
				tasks = append(tasks, task)
				storage.WriteTasks(tasks)
			case "Вывести":
				storage.LoadTasks()
			default:
				fmt.Println("Введите help для подсказки")
			}
		}
	}
}

func createTask(statusInput, description string) (models.Task, error) {
	status, err := models.ParseStatus(statusInput)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:          rand.Int(),
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
	}, nil
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

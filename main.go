package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todocli/models"
	"todocli/storage"

	"github.com/google/uuid"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Файла тасок нет")
	}

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
			parts := strings.SplitN(strings.TrimSpace(text), " ", 3)
			actionType := parts[0]
			if len(parts) > 1 {
				parts[1] = strings.TrimSpace(parts[1])
			}
			if len(parts) > 2 {
				parts[2] = strings.TrimSpace(parts[2])
			}
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
				task, err := CreateTask(parts[1], parts[2])
				if err != nil {
					fmt.Println(err)
					continue
				}
				tasks = append(tasks, task)
				err = storage.WriteTasks(tasks)
				if err != nil {
					fmt.Println("Ошибка сохранения:", err)
				}
			case "Вывести":
				if len(parts) > 1 {
					fmt.Println("тут могал быть ваша реклама")
					PrintTasksByStatus(tasks, parts[1])
				} else {
					PrintAllTasks(tasks)
				}
			default:
				fmt.Println("Введите help для подсказки")
			}
		}
	}
}

func CreateTask(statusInput, description string) (models.Task, error) {
	status, err := models.ParseStatus(statusInput)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:          uuid.New().String(),
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
	}, nil
}

func PrintAllTasks(tasks []models.Task) {
	for _, task := range tasks {
		PrintTask(task)
	}
}

func PrintTasksByStatus(tasks []models.Task, statusInput string) {
	status, err := models.ParseStatus(statusInput)
	if err != nil {
		return
	}
	for _, task := range tasks {
		if task.Status == status {
			PrintTask(task)
		}
	}
}

func PrintTask(task models.Task) {
	fmt.Println("--------------------------------------------")
	fmt.Println("ID - ", task.ID)
	fmt.Println("Статус - ", task.Status)
	fmt.Println("Дата создания - ", task.CreatedAt.Format("2020-01-14 15:05"))
	fmt.Println("Дата изменения - ", task.UpdatedAt.Format("2020-01-14 15:05"))
	fmt.Println("Описание - ", task.Description)
}

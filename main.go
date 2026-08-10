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
		var text string
		for i := 1; i < len(os.Args); i++ {
			text += string(os.Args[i]) + " "
		}
		actionType = os.Args[1]
		tasks = ExecuteCommand(actionType, text, tasks, err)
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("--------------------------------------------")
			fmt.Print("Введите действие - ")
			text, _ := reader.ReadString('\n')
			parts := strings.SplitN(strings.TrimSpace(text), " ", 2)
			actionType := parts[0]
			tasks = ExecuteCommand(actionType, text, tasks, err)
		}
	}
}

func ExecuteCommand(actionType string, text string, tasks []models.Task, err error) []models.Task {
	switch actionType {
	case "help":
		fmt.Println("Список действий:")
		fmt.Println("Создать задачу - Создать <статус> <описание>")
		fmt.Println("Вывести задачи - Вывести <id (если нужно конкретное)>")
		fmt.Println("Изменить задачу - Изменить <id> <поле> <значение>")
		fmt.Println("Удалить задачу - Удалить <id>")
	case "Создать":
		parts := SplitCommand(text, 3)
		if len(parts) < 3 {
			fmt.Println("Создать <статус> <описание>")
			break
		}
		task, err := CreateTask(parts[1], parts[2])
		if err != nil {
			fmt.Println(err)
			break
		}
		tasks = append(tasks, task)
		err = storage.WriteTasks(tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения:", err)
		}
	case "Вывести":
		parts := SplitCommand(text, 2)
		if len(parts) > 1 {
			PrintTasksByStatus(tasks, parts[1])
		} else {
			PrintAllTasks(tasks)
		}
	case "Удалить":
		parts := SplitCommand(text, 2)
		if len(parts) > 1 {
			tasks, err = storage.DeleteTask(tasks, parts[1])
			if err != nil {
				fmt.Println("Нету такой задачи")
			}
		}
		err = storage.WriteTasks(tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения:", err)
		}
	case "Изменить":
		parts := SplitCommand(text, 4)
		if len(parts) > 2 {
			tasks, err = storage.ChangeTask(tasks, parts[1], parts[2], parts[3])
			if err != nil {
				fmt.Println("Ошибка изменения:", err)
				break
			}
			err = storage.WriteTasks(tasks)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
		}
	default:
		fmt.Println("Введите help для подсказки")
	}
	return tasks
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
	fmt.Println("Дата создания - ", task.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Println("Дата изменения - ", task.UpdatedAt.Format("2006-01-02 15:04"))
	fmt.Println("Описание - ", task.Description)
}

func SplitCommand(command string, countParts int) []string {
	commands := strings.SplitN(strings.TrimSpace(command), " ", countParts)
	for i, val := range commands {
		commands[i] = strings.TrimSpace(val)
	}
	return commands
}

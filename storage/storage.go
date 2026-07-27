package storage

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"todocli/models"
)

func WriteTasks(tasks []models.Task) error {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	taskJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	file.Write(taskJSON)
	return nil

}

func LoadTasks() ([]models.Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(tasks []models.Task, id string) ([]models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	return tasks, errors.New("Нету такой задачи")
}

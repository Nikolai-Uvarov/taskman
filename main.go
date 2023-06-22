package main

import (
	"fmt"
	"taskman/pkg/storage"
	"time"
)

func main() {

	fmt.Println("Все задачи:")
	tasks, err := storage.GetTasks()
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}
	fmt.Println("")

	fmt.Println("Задачи по id автора = 1:")
	tasks, err = storage.GetTasksByAuthor(1)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}
	fmt.Println("")

	fmt.Println("Задачи по метке = Продукт:")
	tasks, err = storage.GetTasksByTag("Продукт")
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}
	fmt.Println("")

	fmt.Println("Обновление задачи с id = 1:")
	t, err := storage.UpdateTask(1, time.Now(), time.Unix(0, 0), 3, 1, `updated task 1`, `updated content`)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	fmt.Println(t.Id, t.Title, t.Author_id)
	fmt.Println("")

	fmt.Println("Создание новой задачи")
	_, err = storage.NewTask(1, 2, "New task", "test new task from golang code")
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	fmt.Println("Все задачи:")
	tasks, err = storage.GetTasks()
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}
	fmt.Println("")

	fmt.Println("Удаление задачи с id = 1:")
	err = storage.DeleteTask(1)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	fmt.Println("Все задачи:")
	tasks, err = storage.GetTasks()
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}
	fmt.Println("")

	storage.DB.Close()
}

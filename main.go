package main

import (
	"fmt"
	"taskman/pkg/storage"
)

func main() {

	//tasks, _ := storage.GetTasksByAuthor(1)
	//tasks, _ := storage.GetTasks()
	//fmt.Println((tasks[0]).Opened)
	tasks, _ := storage.GetTasksByTag("Продукт")

	//_, err := storage.NewTask(1, 2, "New task", "test new task from golang code")

	for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}

	storage.DB.Close()
}

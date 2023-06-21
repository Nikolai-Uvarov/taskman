package main

import (
	"fmt"
	"taskman/pkg/storage"
	"time"
)

func main() {

	//tasks, _ := storage.GetTasksByAuthor(1)
	//tasks, _ := storage.GetTasks()
	//fmt.Println((tasks[0]).Opened)
	//tasks, _ := storage.GetTasksByTag("Продукт")
	task, _ := storage.UpdateTask(1, time.Now(), time.Unix(0, 0), 3, 1, `updated task 1`, `updated content`)

	//_, err := storage.NewTask(1, 2, "New task", "test new task from golang code")

	fmt.Println(task)

	/*for _, t := range tasks {
		fmt.Println(t.Id, t.Title, t.Author_id)
	}*/

	storage.DB.Close()
}

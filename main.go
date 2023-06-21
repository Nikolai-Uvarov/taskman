package main

import (
	"fmt"
	"taskman/pkg/storage"
)

func main() {

	//tasks, _ := storage.GetTasks()

	//fmt.Println((tasks[0]).Opened)

	_, err := storage.NewTask(1, 2, "New task", "test new task from golang code")

	fmt.Println(err)

	storage.DB.Close()
}

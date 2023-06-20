package main

import (
	"fmt"
	"taskman/pkg/storage"
)

func main() {

	tasks, _ := storage.GetTasks()

	fmt.Println((tasks[0]).Opened)

	storage.DB.Close()
}

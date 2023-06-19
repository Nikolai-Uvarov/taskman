package main

import (
	"fmt"
	"taskman/pkg/storage"
)

func main() {
	fmt.Println(storage.GetTask())
	fmt.Println("OK")
	storage.DB.Close()
}

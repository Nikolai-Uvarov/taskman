/*
Пакет, который предоставляет необходимые методы для работы с БД.

API пакета storage  позволяет:

Создавать новые задачи,
Получать список всех задач,
Получать список задач по автору,
Получать список задач по метке,
Обновлять задачу по id,
Удалять задачу по id.*/

package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
	id          int64
	opened      time.Time
	closed      time.Time
	author_id   int64
	assigned_id int64
	title       string
	content     string
}

var DB *pgxpool.Pool
var ctx context.Context

func init() {
	ctx = context.Background()
	// Подключение к БД. Функция возвращает объект БД.
	var err error
	DB, err = pgxpool.Connect(ctx, "postgres://postgres:Stasi123@192.168.1.35:5432/tasks")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected succesfully")
}

func NewTask(id int64, opened time.Time, closed time.Time,
	author_id int64, assigned_id int64, title string, content string) *Task {

	return &Task{id, opened, closed, author_id, assigned_id, title, content}
}

func GetTask() string {
	rows, _ := DB.Query(ctx, `SELECT title FROM tasks;`)

	var title string
	for rows.Next() {
		rows.Scan(&title)
	}

	return title

}

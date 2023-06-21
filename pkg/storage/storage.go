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
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
	Id          int64
	Opened      time.Time
	Closed      time.Time
	Author_id   int64
	Assigned_id int64
	Title       string
	Content     string
}

var DB *pgxpool.Pool
var ctx context.Context

func init() {
	ctx = context.Background()
	// Подключение к БД
	dbpass := os.Getenv("dbpass")
	var err error
	DB, err = pgxpool.Connect(ctx, "postgres://postgres:"+dbpass+"@192.168.1.35:5432/tasks")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func NewTask(author_id int64, assigned_id int64, title string, content string) (*Task, error) {

	o := time.Now()

	rows, err := DB.Query(ctx,
		`INSERT INTO tasks(opened,author_id, assigned_id, title,content) 
		VALUES (($1), ($2), ($3), ($4),($5)) 
		RETURNING id;`,
		o.Unix(), author_id, assigned_id, title, content)

	if err != nil {
		return nil, err
	}

	//получаем из БД id созданной задачи
	var id []int64
	for rows.Next() {

		var ci int64
		err = rows.Scan(&ci)

		if err != nil {
			return nil, err
		}
		id = append(id, ci)
	}

	return &Task{id[0], o, time.Unix(0, 0), author_id, assigned_id, title, content}, rows.Err()
}

func GetTasks() ([]Task, error) {

	rows, err := DB.Query(ctx, `SELECT * FROM tasks ORDER BY id;`)

	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		var o, c int64
		err = rows.Scan(
			&t.Id,
			&o,
			&c,
			&t.Author_id,
			&t.Assigned_id,
			&t.Title,
			&t.Content)

		if err != nil {
			return nil, err
		}

		t.Opened = time.Unix(o, 0)
		t.Closed = time.Unix(c, 0)

		tasks = append(tasks, t)
	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

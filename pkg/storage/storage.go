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

	return parseTasks(rows)
}

func GetTasksByAuthor(id int64) ([]Task, error) {

	rows, err := DB.Query(ctx, `SELECT * FROM tasks WHERE author_id=($1) ORDER BY id;`, id)

	if err != nil {
		return nil, err
	}

	return parseTasks(rows)
}

func GetTasksByTag(tag string) ([]Task, error) {

	rows, err := DB.Query(ctx,
		`SELECT t.* FROM tasks t
		JOIN tasks_labels tl ON t.id = tl.task_id
		JOIN labels l ON l.id = tl.label_id
		WHERE l.name = ($1);`, tag)

	if err != nil {
		return nil, err
	}

	return parseTasks(rows)
}

func parseTasks(rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
}) (tasks []Task, err error) {

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

func UpdateTask(id int64, opened time.Time, closed time.Time,
	author_id int64, assigned_id int64, title string, content string) (*Task, error) {

	_, err := DB.Exec(ctx,
		`UPDATE tasks
		SET opened = ($1), closed = ($2), 
		author_id = ($3), assigned_id=($4), 
		title = ($5),content = ($6)
		WHERE id = ($7);`,
		opened.Unix(), closed.Unix(),
		author_id, assigned_id, title, content, id)

	if err != nil {
		return nil, err
	}

	return &Task{id, opened, closed, author_id, assigned_id, title, content}, nil
}

func DeleteTask(id int64) error {

	tx, err := DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM tasks_labels WHERE task_id = ($1); `, id)

	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `DELETE FROM tasks WHERE id = ($1); `, id)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

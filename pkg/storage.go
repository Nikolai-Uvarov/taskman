/*Задача: создать пакет, который бы позволял выполнять все требуемые приложению операции с БД.
Получение списка задач,
Получения информации о конкретной задаче по ее номеру,
Создание задачи,
Обновление задачи,
Удаление задачи,
Создание массива задач.*/

package storage

import "time"

type Task struct {
	id          int64
	opened      time.Time
	closed      time.Time
	author_id   int64
	assigned_id int64
	title       string
	content     string
}

func NewTask(id int64, opened time.Time, closed time.Time, author_id int64, assigned_id int64, title string, content string) *Task {
	return &Task{id, opened, closed, author_id, assigned_id, title, content}
}

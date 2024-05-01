package database

import (
	"database/sql"
)

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func getTaskByRows(rows *sql.Rows) ([]Task, error) {
	var res []Task
	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		res = append(res, task)
	}
	return res, nil
}

func (td *TaskData) Insert(task Task) (int, error) {
	query := insertTable

	res, err := td.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	return int(lastId), err
}

func (td *TaskData) Update(task Task) (bool, error) {
	query := updateTable

	res, err := td.db.Exec(query,
		sql.Named("id", task.Id),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return false, err
	}

	updated, err := res.RowsAffected()
	return updated == 1, err
}

func (td *TaskData) Delete(id int) (bool, error) {
	query := deleteWithCond

	res, err := td.db.Exec(query, sql.Named("id", id))
	if err != nil {
		return false, err
	}
	deleted, err := res.RowsAffected()
	return deleted == 1, err
}

func (td *TaskData) GetTask(id int) (Task, error) {
	query := getTaskWithCond

	row := td.db.QueryRow(query, sql.Named("id", id))
	task := Task{}
	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	return task, err
}

func (td *TaskData) GetTaskList(limit int) ([]Task, error) {
	query := getTaskList

	rows, err := td.db.Query(query, sql.Named("limit", limit))
	if err != nil {
		return nil, err
	}

	return getTaskByRows(rows)
}

func (td *TaskData) GetTaskByDate(date string, limit int) ([]Task, error) {
	query := getTaskByDate

	rows, err := td.db.Query(query,
		sql.Named("date", date),
		sql.Named("limit", limit),
	)
	if err != nil {
		return nil, err
	}

	return getTaskByRows(rows)
}

func (td *TaskData) GetTaskBySearchString(search string, limit int) ([]Task, error) {
	query := getTaskBySearchString

	rows, err := td.db.Query(query,
		sql.Named("search", "%"+search+"%"),
		sql.Named("limit", limit),
	)
	if err != nil {
		return nil, err
	}

	return getTaskByRows(rows)
}

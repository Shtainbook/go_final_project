package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type taskDataConnection interface {
	Insert(task Task) (int, error)
	Update(task Task) (bool, error)
	Delete(id int) (bool, error)
	GetTask(id int) (Task, error)
	GetTaskList(limit int) ([]Task, error)
	GetTaskByDate(date string, limit int) ([]Task, error)
	GetTaskBySearchString(search string, limit int) ([]Task, error)
	CloseDB()
}

type TaskData struct {
	db *sql.DB
}

func InitTaskData(dataSourceName string) (*TaskData, error) {

	db, err := openDb(dataSourceName)
	if err != nil {
		return nil, err
	}

	return &TaskData{db: db}, nil

}

func (td *TaskData) CloseDb() {
	_ = td.db.Close()
}

func openDb(dataSourceName string) (*sql.DB, error) {

	appPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true

	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if install {
		_, err = db.Exec(createTable)
	}
	return db, err
}

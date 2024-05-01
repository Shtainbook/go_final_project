package database

const (
	// createTable - создание таблицы
	createTable = `
	CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	date VARCHAR(8) NOT NULL,
	title TEXT NOT NULL,
	comment TEXT,
	repeat VARCHAR(128)
	);

	CREATE INDEX IF NOT EXISTS index_date ON scheduler (date);
	`

	updateTable = "UPDATE scheduler SET date=:date,title=:title,comment=:comment,repeat=:repeat WHERE id=:id"

	insertTable = "INSERT INTO scheduler(date,title,comment,repeat) VALUES (:date,:title,:comment,:repeat)"

	deleteWithCond = "DELETE FROM scheduler WHERE id=:id"

	getTaskWithCond = "SELECT * FROM scheduler WHERE id=:id"

	getTaskBySearchString = "SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search  ORDER BY date LIMIT :limit"

	getTaskByDate = "SELECT * FROM scheduler WHERE date=:date ORDER BY date LIMIT :limit"

	getTaskList = "SELECT * FROM scheduler ORDER BY date LIMIT :limit"
)

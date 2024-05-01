package config

import (
	"os"
)

const (
	todoPort  = "TODO_PORT"
	webPath   = "WEB_PATH"
	hostName  = "HOSTNAME"
	dbFile    = "TODO_DBFILE"
	todoPass  = "TODO_PASSWORD"
	secretKey = "SECRET_KEY"
)

var Manager struct {
	TodoPort  string
	WebPath   string
	HostName  string
	DBFile    string
	TODOPass  string
	SecretKey string
}

func InitEnv() {
	setTODOPort()
	setWEBPath()
	setHost()
	setDBFile()
	setTODOPass()
	setSecretKey()

	return
}

func setTODOPort() {

	if value := os.Getenv(todoPort); value != "" {
		Manager.TodoPort = value
		return
	}

	_ = os.Setenv(todoPort, "7540")
	Manager.TodoPort = os.Getenv(todoPort)
	return
}

func setWEBPath() {

	if value := os.Getenv(webPath); value != "" {
		Manager.WebPath = value
		return
	}

	_ = os.Setenv(webPath, "./web/")

	Manager.WebPath = os.Getenv(webPath)
	return
}

func setHost() {

	if value := os.Getenv(hostName); value != "" {
		Manager.HostName = value
		return
	}

	_ = os.Setenv(hostName, "localhost")
	Manager.HostName = os.Getenv(hostName)
	return
}

func setDBFile() {

	if value := os.Getenv(dbFile); value != "" {
		Manager.DBFile = value
		return
	}

	_ = os.Setenv(dbFile, "./scheduler.db")
	Manager.DBFile = os.Getenv(dbFile)
	return
}

func setTODOPass() {

	if value := os.Getenv(todoPass); value != "" {
		Manager.TODOPass = value
		return
	}

	_ = os.Setenv(todoPass, "")
	Manager.TODOPass = os.Getenv(todoPass)
	return
}

func setSecretKey() {

	if value := os.Getenv(secretKey); value != "" {
		Manager.SecretKey = value
		return
	}

	_ = os.Setenv(secretKey, "")
	Manager.SecretKey = os.Getenv(secretKey)
	return
}

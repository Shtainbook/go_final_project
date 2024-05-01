package main

import (
	"go_final_project/internal/config"
	"go_final_project/internal/database"
	"go_final_project/internal/handler/date"
	"go_final_project/internal/handler/file"
	"go_final_project/internal/handler/sign"
	"go_final_project/internal/handler/task"
	"go_final_project/internal/middleware"
	signS "go_final_project/internal/service/sign"
	taskS "go_final_project/internal/service/task"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	config.InitEnv()

	taskData, err := database.InitTaskData(config.Manager.DBFile)
	defer taskData.CloseDb()

	if err != nil {
		log.Fatalf("cann't create db: %v", err)
	}

	r := chi.NewRouter()

	r.Get("/*", file.Server)

	taskS.Service = taskS.InitTaskService(taskData)
	signS.Service = signS.InitSignService(config.Manager.TODOPass, []byte(config.Manager.SecretKey))

	r.Post("/api/signin", sign.PostPass)
	r.Get("/api/nextdate", date.GetNextDate)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Sign)

		r.Post("/api/task", task.PostTask)
		r.Put("/api/task", task.PutTask)
		r.Delete("/api/task", task.DeleteTask)
		r.Get("/api/task", task.GetTask)
		r.Post("/api/task/done", task.DonePostTask)
		r.Get("/api/tasks", task.GetTasks)
	})

	port := config.Manager.TodoPort
	host := config.Manager.HostName

	log.Printf("app is starting at %v:%v \n", host, port)
	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatalf("cann't run app: %v", err)
	}

}

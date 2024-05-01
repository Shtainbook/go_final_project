package task

import (
	"bytes"
	"encoding/json"
	"go_final_project/internal/database"
	taskS "go_final_project/internal/service/task"
	"go_final_project/internal/util"
	"net/http"
)

func taskFromBody(r *http.Request) (database.Task, error) {
	var task database.Task

	buff := bytes.Buffer{}

	_, err := buff.ReadFrom(r.Body)
	if err != nil {
		return database.Task{}, err
	}

	err = json.Unmarshal(buff.Bytes(), &task)
	return task, err
}

func DonePostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.URL.Query().Get("id")
	err := taskS.Service.DoneTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	_, _ = w.Write([]byte("{}"))
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	task, err := taskFromBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(util.MarshalError(err))
		return
	}

	id, err := taskS.Service.Create(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	ansBody, err := json.Marshal(
		struct {
			Id int `json:"id"`
		}{Id: id},
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	_, err = w.Write(ansBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	task, err := taskFromBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	err = taskS.Service.Update(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	w.Write([]byte("{}"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.URL.Query().Get("id")
	err := taskS.Service.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	_, _ = w.Write([]byte("{}"))
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.URL.Query().Get("id")
	task, err := taskS.Service.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	response, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}
	w.Write(response)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	search := r.URL.Query().Get("search")
	var tasks *database.TaskList
	var err error
	if len(search) == 0 {
		tasks, err = taskS.Service.GetTasks()
	} else {
		tasks, err = taskS.Service.SearchTasks(search)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	response, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(util.MarshalError(err))
		return
	}

	_, _ = w.Write(response)
}

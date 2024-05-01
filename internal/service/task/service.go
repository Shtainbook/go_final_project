package task

import (
	"errors"
	"go_final_project/internal/database"
	dateS "go_final_project/internal/service/date"
	"go_final_project/internal/util"
	"strconv"
	"time"
)

const (
	requireTitle = "require task title"
	notFoundTask = "not found task"
)

type Action interface {
}

type TaskService struct {
	taskData *database.TaskData
}

var Service *TaskService

func sliceToTasks(list []database.Task) *database.TaskList {
	if list == nil {
		return &database.TaskList{Tasks: []database.Task{}}
	}

	return &database.TaskList{Tasks: list}
}

func convert(task *database.Task) error {
	if len(task.Title) == 0 {
		return errors.New(requireTitle)
	}
	now := time.Now().Format(util.DateFormat)
	if len(task.Date) == 0 {
		task.Date = now
	}
	_, err := time.Parse(util.DateFormat, task.Date)
	if err != nil {
		return err
	}
	nextDate, err := dateS.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return err
	}

	if task.Date < now {
		if len(nextDate) == 0 {
			task.Date = now
		} else {
			task.Date = nextDate
		}
	}
	return nil
}

func InitTaskService(taskData *database.TaskData) *TaskService {
	return &TaskService{taskData: taskData}
}

func (service *TaskService) Create(task database.Task) (int, error) {
	err := convert(&task)
	if err != nil {
		return 0, err
	}
	id, err := service.taskData.Insert(task)
	return id, err
}

func (service *TaskService) Update(task database.Task) error {
	err := convert(&task)
	if err != nil {
		return err
	}

	updated, err := service.taskData.Update(task)
	if err != nil {
		return err
	}
	if !updated {
		return errors.New(notFoundTask)
	}
	return nil
}

func (service *TaskService) DeleteTask(id string) error {
	convId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	deleted, err := service.taskData.Delete(convId)
	if err != nil {
		return err
	}
	if !deleted {
		return errors.New(notFoundTask)
	}
	return nil
}

func (service *TaskService) GetTasks() (*database.TaskList, error) {
	list, err := service.taskData.GetTaskList(util.TaskListRowsLimit)
	if err != nil {
		return nil, err
	}
	return sliceToTasks(list), err
}

func (service *TaskService) SearchTasks(search string) (*database.TaskList, error) {
	date, err := time.Parse(util.SearchDateFormat, search)
	if err == nil {
		list, err := service.taskData.GetTaskByDate(date.Format(util.DateFormat), util.TaskListRowsLimit)
		if err != nil {
			return nil, err
		}
		return sliceToTasks(list), nil
	}
	list, err := service.taskData.GetTaskBySearchString(search, util.TaskListRowsLimit)
	return sliceToTasks(list), err
}

func (service *TaskService) GetTask(id string) (*database.Task, error) {
	convId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	task, err := service.taskData.GetTask(convId)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (service *TaskService) DoneTask(id string) error {
	convId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	task, err := service.taskData.GetTask(convId)
	if err != nil {
		return err
	}

	if len(task.Repeat) == 0 {
		deleted, err := service.taskData.Delete(convId)
		if err != nil {
			return err
		}
		if !deleted {
			return errors.New(notFoundTask)
		}
		return nil
	}

	task.Date, err = dateS.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return err
	}

	updated, err := service.taskData.Update(task)
	if err != nil {
		return err
	}
	if !updated {
		return errors.New(notFoundTask)
	}
	return nil
}

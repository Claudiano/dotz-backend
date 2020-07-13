package repositories

import (
	"log"

	"github.com/claudiano/dotz-backend/domain/model"
	"github.com/claudiano/dotz-backend/infra/database"
)

type TaskRepository struct{}

func (t *TaskRepository) FindOne(IDTask int32) (model.Task, error) {
	var task model.Task
	db := database.ConnectDB()

	err := db.Where("IDTask = ?", IDTask).Find(&task).Error

	if err != nil {
		log.Printf("Erro a buscar a task: %v\n", err)
	}

	return task, err
}

func (t *TaskRepository) FindAllByIDUser(IDUser int32) ([]model.Task, error) {
	db := database.ConnectDB()
	var tasks []model.Task

	err := db.Where("IDUsuario = ?", IDUser).Find(&tasks).Error

	if err != nil {
		log.Printf("Erro ao buscar todas as tasks: %v\n", err)
	}

	return tasks, err
}

func (t *TaskRepository) FindByIDUserAndStatus(IDUser int32, status string) ([]model.Task, error) {
	db := database.ConnectDB()
	var tasks []model.Task

	err := db.Where("IDUsuario = ? and Status = ?", IDUser, status).Find(&tasks).Error

	if err != nil {
		log.Printf("Erro ao buscar as tasks: %v\n", err)
	}

	return tasks, err
}

func (t *TaskRepository) Save(task model.Task) (model.Task, error) {
	db := database.ConnectDB()
	var err error

	if task.IDTask != 0 {
		err = db.Save(&task).Error
	} else {
		err = db.Create(&task).Error
	}
	return task, err
}

package repositories

import "github.com/claudiano/dotz-backend/domain/model"

type TaskRepository struct{}

func (t *TaskRepository) FindOne(idTask int32) (model.Task, error){
	return model.Task{}, nil
}

func (t *TaskRepository) FindAll() ([]model.Task, error){
	return nil, nil
}

func (t *TaskRepository) Save( task model.Task) (model.Task, error){
	return model.Task{}, nil
}

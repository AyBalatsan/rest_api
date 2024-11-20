package service

import (
	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
)

type TodoListServe struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListServe {
	return &TodoListServe{repo: repo}
}

func (s *TodoListServe) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListServe) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListServe) GetById(idUser, id int) (todo.TodoList, error) {
	return s.repo.GetById(idUser, id)
}
func (s *TodoListServe) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListServe) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}

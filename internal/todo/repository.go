package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(todo *Todo) error
	GetAll() ([]Todo, error)
	GetByID(id uint) (*Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(todo *Todo) error {
	return r.db.Create(todo).Error
}

func (r *repository) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *repository) GetByID(id uint) (*Todo, error) {
	var todo Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *repository) Update(todo *Todo) error {
	data := map[string]interface{}{
		"title":       todo.Title,
		"description": todo.Description,
		"completed":   todo.Completed,
	}
	return r.db.Model(&Todo{}).Where("id = ?", todo.ID).Updates(data).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Todo{}, id).Error
}

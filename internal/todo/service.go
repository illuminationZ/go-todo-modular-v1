package todo

type Service interface {
	CreateTodo(todo *Todo) error
	GetAllTodos() ([]Todo, error)
	GetTodoByID(id uint) (*Todo, error)
	UpdateTodoByID(todo *Todo) error
	DeleteTodoByID(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateTodo(todo *Todo) error {
	return s.repo.Create(todo)
}

func (s *service) GetAllTodos() ([]Todo, error) {
	return s.repo.GetAll()
}

func (s *service) GetTodoByID(id uint) (*Todo, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateTodoByID(todo *Todo) error {
	return s.repo.Update(todo)
}

func (s *service) DeleteTodoByID(id uint) error {
	return s.repo.Delete(id)
}
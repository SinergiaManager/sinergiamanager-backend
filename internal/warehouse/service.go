package warehouse

type Service interface {
	GetAll()
	GetById()
	Create()
	Update()
	Delete()
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() {}

func (s *service) GetById() {}

func (s *service) Create() {}

func (s *service) Update() {}

func (s *service) Delete() {}

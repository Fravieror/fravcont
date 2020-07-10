package domain

type Repository interface {
	Save(Movement) error
	Get(id string) (*Movement, error)
	GetList(client string) ([]Movement, error)
}

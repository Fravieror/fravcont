package infrastructure

import (
	"api-test/movement/domain"
	"fmt"
)

type MySQLRepository struct{}

func NewMySQLRepository(redisURL string) (domain.Repository, error) {
	repo := &MySQLRepository{}
	// Establish connection with external DB
	return repo, nil
}

func (m *MySQLRepository) Save(movement *domain.Movement) error {
	fmt.Println("Movement save")
	return nil
}
func (m *MySQLRepository) GetList() ([]*domain.Movement, error) {
	fmt.Println("Get all movements")
	list := []*domain.Movement{}
	return list, nil
}
func (m *MySQLRepository) Get(id string) (*domain.Movement, error) {
	fmt.Println("Get movement")
	return &domain.Movement{}, nil
}

package application

import (
	"api-test/movement/domain"
	"fmt"
)

type Bussines struct {
	repository domain.Repository
}

func NewBussines(rep domain.Repository) domain.Service {
	return &Bussines{
		repository: rep,
	}
}

func (b *Bussines) Get(id string) *domain.Movement {
	mov, err := b.repository.Get(id)
	if err != nil {
		fmt.Println("Error Get bussines")
	}
	return mov
}

func (b *Bussines) Save(movement *domain.Movement) {
	b.repository.Save(movement)
}

func (b *Bussines) GetList() []*domain.Movement {
	movs, err := b.repository.GetList()
	if err != nil {
		fmt.Println("Error Get bussines")
	}
	return movs
}

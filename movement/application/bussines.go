package application

import "api-test/movement/domain"

type bussines struct {
	repository domain.Repository
}

func (b *bussines) saveMovement(movement *domain.Movement) {
	b.saveMovement(movement)
}
func (b *bussines) getList() *[]domain.Movement {
	return b.getList()
}
func (b *bussines) getMovement(id string) *domain.Movement {
	return b.getMovement(id)
}

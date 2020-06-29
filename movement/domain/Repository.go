package domain

type Repository interface {
	saveMovement(Movement)
	getList() []Movement
	getMovement(string) Movement
}

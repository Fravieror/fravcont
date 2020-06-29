package domain

type Service interface {
	save(Movement)
	getList() []Movement
	get(string) Movement
}

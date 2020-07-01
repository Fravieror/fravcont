package domain

type Service interface {
	Save(*Movement)
	GetList() []*Movement
	Get(id string) *Movement
}

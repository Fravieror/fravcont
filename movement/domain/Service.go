package domain

type Service interface {
	Save(*Movement) ResponseMovement
	GetList(client string) ResponseMovement
	Get(id string) ResponseMovement
}

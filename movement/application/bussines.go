package application

import (
	"api-test/movement/domain"
	"encoding/json"
	"log"

	"github.com/teris-io/shortid"
)

type Bussines struct {
	repository domain.Repository
}

func NewBussines(rep domain.Repository) domain.Service {
	return &Bussines{
		repository: rep,
	}
}

func (b *Bussines) Get(id string) domain.ResponseMovement {
	var response = domain.NewResponseMovement()
	mov, err := b.repository.Get(id)
	if err != nil {
		log.Println(err.Error())
		response.Message = "no se pudo"
		return *response
	}
	if mov == nil {
		response.Message = "no hay registros"
		return *response
	}
	rawMsg, err := json.Marshal(mov)
	if err != nil {
		log.Println(err.Error())
		response.Message = "no se pudo"
		return *response
	}
	response.Data = rawMsg
	return *response

}

func (b *Bussines) Save(movement *domain.Movement) domain.ResponseMovement {
	var response = domain.NewResponseMovement()
	movement.Id = shortid.MustGenerate()
	err := b.repository.Save(*movement)
	if err != nil {
		log.Println(err.Error())
		response.Message = "no se pudo"
		return *response
	}
	return *response
}

func (b *Bussines) GetList(client string) domain.ResponseMovement {
	var response = domain.NewResponseMovement()
	movs, err := b.repository.GetList(client)
	if err != nil {
		log.Println(err.Error())
		response.Message = "no se pudo"
		return *response
	}
	if movs == nil {
		response.Message = "no hay registros"
		return *response
	}
	rawMsg, err := json.Marshal(movs)
	if err != nil {
		log.Println(err.Error())
		response.Message = "no se pudo"
		return *response
	}
	response.Data = rawMsg
	return *response
}

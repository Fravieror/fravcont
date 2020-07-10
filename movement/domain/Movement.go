package domain

//Object base af all.
type Movement struct {
	Id          string
	Client      string
	Description string
	Value       float32
	Date        string
}

type ResponseMovement struct {
	Message string
	Data    []byte
}

func NewResponseMovement() *ResponseMovement {
	return &ResponseMovement{
		Message: "ok",
		Data:    nil,
	}
}

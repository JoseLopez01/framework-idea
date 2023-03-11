package services

type AppService struct {
}

func (service AppService) GetMessage() string {
	return "message from service"
}

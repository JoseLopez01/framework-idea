package services

type AppService struct {
}

func (service AppService) GetMessage() string {
	return "MESSAGE FROM SERVICE"
}

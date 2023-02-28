package service

type Order interface {
	Get()
}

type orderService struct {
}

func (s *orderService) Get() {

}

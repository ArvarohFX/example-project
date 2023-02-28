package usecase

type Order interface {
	Get()
}

type orderService struct {
}

func (s *orderService) Get() {

}

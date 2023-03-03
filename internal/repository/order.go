package repository

type OrderRepository interface {
	Create() (string, int, error)
}

type orderRepository struct {
}

func (r *orderRepository) Create() (string, int, error) {
	return "param", 10, nil
}

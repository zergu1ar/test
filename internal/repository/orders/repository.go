package orders

import "applicationDesignTest/internal/entity"

type Repository interface {
	Create(order *entity.Order) error
	List() ([]entity.Order, error)
}

func NewRepository(cfg *Config) Repository {
	return &orders{
		cfg:  cfg,
		data: []entity.Order{},
	}
}

type orders struct {
	cfg  *Config
	data []entity.Order
}

func (r *orders) Create(order *entity.Order) error {
	r.data = append(r.data, *order)
	return nil
}

func (r *orders) List() ([]entity.Order, error) {
	return r.data, nil
}

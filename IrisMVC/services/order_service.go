package services

import (
	"app/datamodels"
	"app/repositories"
)

type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order, error)
	GetAllOrder() ([]*datamodels.Order, error)
	DeleteOrderByID(int64) (bool, error)
	InsertOrder(*datamodels.Order) (int64, error)
	UpdateOrder(*datamodels.Order) error
	GetAllOrderInfo() (map[int]map[string]string, error)
}
type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func NewOrderService(o repositories.IOrderRepository) IOrderService {
	return &OrderService{o}
}

func (o *OrderService) GetOrderByID(orderID int64) (*datamodels.Order, error) {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) DeleteOrderByID(orderID int64) (bool, error) {
	return o.OrderRepository.Delete(orderID)
}

func (o *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.OrderRepository.Update(order)
}
func (o *OrderService) GetAllOrderInfo() (map[int]map[string]string, error) {
	return o.OrderRepository.SelectAllWithInfo()
}

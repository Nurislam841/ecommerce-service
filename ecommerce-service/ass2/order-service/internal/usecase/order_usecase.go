package usecase

import (
	"order-service/internal/entity"
	"order-service/internal/repository"
	"proto/orderpb"
)

type OrderUsecase struct {
	repo *repository.OrderRepository
}

func NewOrderUsecase(repo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}
func (u *OrderUsecase) CreateOrder(order *entity.Order) error {
	return u.repo.CreateOrder(order)
}

func (u *OrderUsecase) CreateOrderFromRequest(req *orderpb.CreateOrderRequest) (*entity.Order, error) {
	order := &entity.Order{
		Status: "pending",
	}

	for _, item := range req.Items {
		order.Items = append(order.Items, entity.OrderItem{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	err := u.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *OrderUsecase) GetOrderByIDPB(id string) (*orderpb.Order, error) {
	order, err := u.repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	return convertToPBOrder(order), nil
}

func (u *OrderUsecase) UpdateOrderStatus(id string, status string) error {
	return u.repo.UpdateOrderStatus(id, status)
}

func convertToPBOrder(order *entity.Order) *orderpb.Order {
	pbItems := make([]*orderpb.OrderItem, len(order.Items))
	for i, item := range order.Items {
		pbItems[i] = &orderpb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	return &orderpb.Order{
		Id: order.ID,
		//UserId: order.UserID,
		Status: order.Status,
		Items:  pbItems,
	}
}

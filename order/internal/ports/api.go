package ports

import "github.com/KrongkaewPon/grpc-go/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}

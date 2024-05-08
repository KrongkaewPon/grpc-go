package ports

import "github.com/KrongkaewPon/grpc-go/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}

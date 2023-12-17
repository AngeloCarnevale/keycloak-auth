package productuc

import "demo-backend/domain/entities"

type ProducstDataStore interface {
	Create(product *entities.Product) error
}
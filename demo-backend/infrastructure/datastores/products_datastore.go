package datastores

import (
	"demo-backend/domain/entities"
	"sync"

	"github.com/google/uuid"
)

type producstDataStore struct {
	products map[uuid.UUID]entities.Product
	sync.Mutex
}

func NewProductsDatastore() *producstDataStore {
	return &producstDataStore{
		products: make(map[uuid.UUID]entities.Product),
	}
}

func (ds *producstDataStore) Create(product *entities.Product) error {
	ds.Lock()
	ds.products[product.Id] = *product
	ds.Unlock()
	return nil
}

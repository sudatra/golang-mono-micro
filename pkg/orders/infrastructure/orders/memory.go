package orders

import "golang-mono-micro/pkg/orders/domain/orders"

type MemoryRepository struct {
	orders []orders.Order
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]orders.Order{}};
}

func (m *MemoryRepository) Save(orderToSave *orders.Order) error {
	for i, o := range m.orders{
		if o.ID() == orderToSave.ID() {
			m.orders[i] = *orderToSave;
			return nil
		}
	}

	m.orders = append(m.orders, *orderToSave)
	return nil;
}

func (m *MemoryRepository) ByID(id orders.ID) (*orders.Order, error) {
	for _, o := range m.orders{
		if o.ID() == id {
			return &o, nil
		}
	}

	return &orders.Order{}, orders.ErrorNotFound;
}
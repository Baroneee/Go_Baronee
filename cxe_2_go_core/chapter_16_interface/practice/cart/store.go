package cart

type CartStore interface {
	GetById(ID string) (*Cart, error)
	Put(cart *Cart) (*Cart, error)
}
type CartStorePostgres struct{}

func (c *CartStorePostgres) GetById(ID string) (*Cart, error) { return nil, nil }
func (c *CartStorePostgres) Put(cart *Cart) (*Cart, error) {
	return cart, nil
}
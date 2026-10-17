package cart

import (
	"errors"
	"os/user"
	"time"

	"github.com/Baroneee/Go_Baronee/cxe_2_go_core/chapter_16_interface/practice/product"
	"github.com/Rhymond/go-money"
)

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
	isLocked  bool
	User         user.User
	Items        []Item
	CurrencyCode string

}

type Item struct {
    product.Product
    Quantity uint8
}

func (c *Cart) TotalPrice() (*money.Money, error) {
    total := money.New(0, c.CurrencyCode)
	var err error
	for _, i := range c.Items {
		itemTotal := i.Price.Multiply(int64(i.Quantity))
		total, err = total.Add(itemTotal)
		if err != nil {
			return nil, err
		}
	}
    return total, nil
}

func (c *Cart) Lock() error {
    if c.isLocked {
        return errors.New("cart is already locked")
    }
	c.isLocked = true
	c.lockedAt = time.Now()
    return nil
}

func (c *Cart) delete() error {
    // to implement
    return nil
}


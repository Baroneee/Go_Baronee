package cart

import (
	"errors"
	"os/user"
	"testing"
	"time"

	"github.com/Baroneee/Go_Baronee/cxe_2_go_core/chapter_14_methods/practice/product"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
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

func TestTotalPrice(t *testing.T) {
    items := []Item{
        {
            Product: product.Product{
                ID:    "p-1254",
                Name:  "Product test",
                Price: money.New(1000, "EUR"),
            },
            Quantity: 2,
        },
        {
            Product: product.Product{
                ID:    "p-1255",
                Name:  "Product test 2",
                Price: money.New(2000, "EUR"),
            },
            Quantity: 1,
        },
    }
    c := Cart{
        ID:           "1254",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        User:         user.User{},
        Items:        items,
        CurrencyCode: "EUR",
    }
    actual, err := c.TotalPrice()
    assert.NoError(t, err)
    assert.Equal(t, money.New(4000, "EUR"), actual)
}

func TestLock(t *testing.T) {
    c := Cart{
        ID: "1254",
    }
    err := c.Lock()
    assert.NoError(t, err)
    assert.True(t, c.isLocked)
    assert.True(t, c.lockedAt.Unix() > 0)
}

func TestLockAlreadyLocked(t *testing.T) {
    c := Cart{
        ID:       "1254",
        isLocked: true,
    }
    err := c.Lock()
    assert.Error(t, err)
}

package main

import (
	"log"

	"github.com/Baroneee/Go_Baronee/cxe_2_go_core/chapter_16_interface/practice/cart"
)

func main() {
    newCart := cart.Cart{}

    totalPrice, err := newCart.TotalPrice()
    if err != nil {
        log.Printf("impossible to compute price of the cart: %s", err)
        return
    }
    log.Println("Total Price", totalPrice.Display())

    var store cart.CartStore = &cart.CartStorePostgres{}

	savedCart, err := store.Put(&newCart)
	if err != nil {
		log.Printf("impossible to save cart: %s", err)
		return
	}
	log.Printf("Cart %s saved successfully to Postgres!\n", savedCart.ID)

    err = newCart.Lock()
    if err != nil {
        log.Printf("impossible to lock the cart: %s", err)
        return
    }

}
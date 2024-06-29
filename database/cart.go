package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't fint the product")
	ErrCantDecodeProduct  = errors.New("Can't decode the product")
	ErrUserIdIsNotValid   = errors.New("User ID is not valid")
	ErrCantUpdateUser     = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItemCart    = errors.New("Unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("Cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}

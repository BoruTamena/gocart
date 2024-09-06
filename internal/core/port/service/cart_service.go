package service

type CartService interface {
	AddItem()
	RemoveItem()
	UpdateQuantity()
	ViewCartItem()
	Checkout()
}

package handler

import (
	"github.com/BoruTamena/internal/core/port/service"
	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	Router  *gin.Engine
	Service service.CartService
}

func NewCartHandler(engine *gin.Engine, service service.CartService) *cartHandler {
	return &cartHandler{
		Router:  engine,
		Service: service,
	}
}

func InitHandler() {

}

func (ch cartHandler) AddItemToCart(c *gin.Context) {

}

func (ch cartHandler) RemoveItemFromCart(c *gin.Context) {

}

func (ch cartHandler) UpdateCartItem(c *gin.Context) {

}

func (ch cartHandler) ViewCartItems(c *gin.Context) {

}

func (ch cartHandler) CheckoutCartItems(c *gin.Context) {

}

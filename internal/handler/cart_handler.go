package handler

import (
	"net/http"

	"github.com/BoruTamena/internal/core/models"
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

	var Item models.Item

	if err := c.ShouldBind(&Item); err != nil {

		// setting error to gin context
		c.Error(err)
		return

	}
	// user don't have shopping session
	if Item.SessionId == 0 {

		session_id, err := ch.Service.CreateShoppingSession()

		if err != nil {

			// setting error
			c.Error(err)
			return
		}

		Item.SessionId = session_id

	}

	quantity, err := ch.Service.AddItem(Item)

	if err != nil {

		// setting error
		c.Error(err)
		return

	}

	if quantity == 10 {

		c.JSON(http.StatusConflict, gin.H{"message": "item exceed max limit "})

	}

	c.JSON(http.StatusCreated, gin.H{"message": "item added successfully", "quantity": quantity})

}

func (ch cartHandler) RemoveItemFromCart(c *gin.Context) {

}

func (ch cartHandler) UpdateCartItem(c *gin.Context) {

}

func (ch cartHandler) ViewCartItems(c *gin.Context) {

}

func (ch cartHandler) CheckoutCartItems(c *gin.Context) {

}

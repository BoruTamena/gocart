package handler

import (
	"net/http"
	"strconv"

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
	// user don't have shopping session before
	if Item.SessionId == 0 {
		// creating new shopping session for user
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

	var Deleted_item models.DeletedItem

	if err := c.ShouldBindQuery(&Deleted_item); err != nil {
		// setting error
		c.Error(err)

		return
	}

	// removing item
	affected_row, err := ch.Service.RemoveItem(Deleted_item)

	if err != nil {
		// setting error
		c.Error(err)
		return
	}
	if affected_row > 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "item removed from cart "})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": " something went wrong "})

	}
}

func (ch cartHandler) ViewCartItems(c *gin.Context) {

	user_id := c.Query("user_id")

	userID, err := strconv.Atoi(user_id)

	if err != nil {
		// setting error
		c.Error(err)
		return

	}

	items, err := ch.Service.ViewCartItem(userID)

	if err != nil {
		// setting error
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success", "data": items})

}

func (ch cartHandler) UpdateCartItem(c *gin.Context) {

}

func (ch cartHandler) CheckoutCartItems(c *gin.Context) {

}

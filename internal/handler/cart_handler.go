package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BoruTamena/internal/core/models"
	"github.com/BoruTamena/internal/core/port/service"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Router  *gin.Engine
	Service service.CartService
}

func NewCartHandler(engine *gin.Engine, service service.CartService) *CartHandler {
	return &CartHandler{
		Router:  engine,
		Service: service,
	}
}

func (ch CartHandler) InitHandler() {
	api := ch.Router.Group("cart")

	api.GET("", ch.ViewCartItems)
	api.POST("/item", ch.AddItemToCart)
	api.POST("/increment", ch.AddItemQuantity)
	api.POST("/decrement", ch.SubtractItemQuantity)
	api.DELETE("/remove", ch.RemoveItemFromCart)

}

func (ch CartHandler) AddItemToCart(c *gin.Context) {

	var Item models.Item

	if err := c.ShouldBind(&Item); err != nil {

		// setting error to gin context
		c.Error(err)
		return

	}
	// user don't have shopping session before
	if Item.SessionId == 0 {
		// creating new shopping session for user

		session_mdl := models.Session{
			ID:     1,
			UserID: 2,
		}
		session_id, err := ch.Service.CreateShoppingSession(c.Request.Context(), session_mdl)

		if err != nil {

			// setting error
			c.Error(err)
			return
		}

		Item.SessionId = session_id

	}

	log.Print(Item)

	quantity, err := ch.Service.AddItem(c.Request.Context(), Item)

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

func (ch CartHandler) RemoveItemFromCart(c *gin.Context) {

	var Deleted_item models.DeletedItem

	if err := c.ShouldBindQuery(&Deleted_item); err != nil {
		// setting error
		c.Error(err)

		return
	}

	// removing item
	affected_row, err := ch.Service.RemoveItem(c.Request.Context(), Deleted_item)

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

func (ch CartHandler) ViewCartItems(c *gin.Context) {

	session_id := c.Query("session_id")

	sessionID, err := strconv.Atoi(session_id)

	if err != nil {
		// setting error
		c.Error(err)
		return

	}

	items, err := ch.Service.ViewCartItem(c.Request.Context(), sessionID)

	if err != nil {
		// setting error
		c.Error(err)

		log.Fatal(err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success", "data": items})

}

func (ch CartHandler) AddItemQuantity(c *gin.Context) {

	var item models.Item

	if err := c.ShouldBind(&item); err != nil {
		// setting error
		c.Error(err)
		return

	}

	affected_row, err := ch.Service.IncreaseItemQuantity(c.Request.Context(), item)

	if err != nil {
		// setting error
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "item quantity increase", "row": affected_row})
}

func (ch CartHandler) SubtractItemQuantity(c *gin.Context) {

	var item models.Item

	if err := c.ShouldBind(&item); err != nil {
		// setting error
		c.Error(err)
		return

	}

	affected_row, err := ch.Service.DecreaseItemQuantity(c.Request.Context(), item)

	if err != nil {
		// setting error
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "item quantity decrease", "row": affected_row})

}

func (ch CartHandler) CheckoutCartItems(c *gin.Context) {

}

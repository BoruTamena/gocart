package steps

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/cucumber/godog"
)

type Item struct {
	Session_id int `json:"session_id,omitempty"`
	Product_id int `json:"product_id,omitempty"`
	Quantity   int `json:"quantity,omitempty"`
}

type Cart struct {
	item   Item
	server *httptest.Server
	resp   *http.Response
}

func (c *Cart) JsonMarshaller(item Item) []byte {

	js_item, _ := json.Marshal(item)

	return js_item

}

func (c *Cart) RegisterItem(product_id, session_id, quantity int) error {

	if product_id != 0 && session_id != 0 && quantity != 0 {

		c.item = Item{
			Product_id: product_id,
			Session_id: session_id,
			Quantity:   session_id,
		}

		return nil

	}

	return errors.New("invalid product ,session and quantity")
}

func (c *Cart) RegisterDeleteItem(P_id, S_id int) error {
	if P_id != 0 && S_id != 0 {

		c.item = Item{
			Product_id: P_id,
			Session_id: S_id,
		}

		return nil

	}

	return errors.New("can't register item ")
}

func (c *Cart) ViewCartItem(session_id int) error {

	Url, err := url.Parse(c.server.URL + "/cart")

	if err != nil {
		return err
	}

	params := url.Values{}

	params.Add("session_id", "1")

	Url.RawQuery = params.Encode()
	// c.server = httptest.NewServer(handler)
	resp, err := c.server.Client().Post(Url.String(), "application/json",
		bytes.NewBuffer(c.JsonMarshaller(c.item)))

	if err != nil {
		return err
	}
	c.resp = resp
	return nil

}

func (c *Cart) AddCartItem(want string) error {

	// adding new item to cart

	resp, err := c.server.Client().Post(c.server.URL+"cart/item",
		"application/json", bytes.NewBuffer(c.JsonMarshaller(c.item)))

	if err != nil {
		return err
	}
	c.resp = resp
	return nil

}

func (c *Cart) RemoveCartItem(S_id, P_id int) error {
	// removing cart item

	resp, err := c.server.Client().Post(c.server.URL+"cart/remove",
		"application/json", bytes.NewBuffer(c.JsonMarshaller(c.item)))

	if err != nil {
		return err
	}
	c.resp = resp
	return nil

}

func (c *Cart) RegisterSteps(ctx *godog.ScenarioContext, s *httptest.Server) {

	c.server = s
	ctx.Step(`^I add item a new product with a product_id,session_id and quantity (\d+),(\d+),(\d+),$`, c.RegisterItem)
	ctx.Step(`^the system should add a new product into cart and return "([^"]*)"\.$`, c.AddCartItem)
}

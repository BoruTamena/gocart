package steps

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/BoruTamena/utils"
)

type item struct {
	Session_id int
	Product_id int
}

type cart struct {
	item
	server *httptest.Server
	resp   *http.Response
}

func (c *cart) RegisterItem(S_id, P_id int) error {

	c.item = item{
		Session_id: S_id,
		Product_id: P_id,
	}

	return nil
}
func (c *cart) IncreaseQuantity() error {

	resp, err := c.server.Client().Post(c.server.URL+"/cart/increment", "application/json", bytes.NewBuffer(utils.JsonMarshaller(c.item)))

	if err != nil {
		return err
	}
	c.resp = resp
	return nil
}

func (c *cart) DecreaseQuantity() error {

	resp, err := c.server.Client().Post(c.server.URL+"/cart/decrement", "application/json", bytes.NewBuffer(utils.JsonMarshaller(c.item)))

	if err != nil {
		return err
	}
	c.resp = resp
	return nil

}

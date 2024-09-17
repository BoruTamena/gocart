package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/BoruTamena/infra/repository"
	"github.com/BoruTamena/internal/core/service"
	"github.com/BoruTamena/internal/handler"
	"github.com/DATA-DOG/go-txdb"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
)

var ( // data base connection variables
	host     = "localhost"
	port     = 5432
	dbtest   = "cart_go_db"
	user     = "postgres"
	password = "root"
	dns      string
)

type Item struct {
	Session_id int `json:"session_id,omitempty"`
	Product_id int `json:"product_id,omitempty"`
	Quantity   int `json:"quantity,omitempty"`
}

type featureTest struct {
	item   Item
	server *httptest.Server
	resp   *http.Response
}

func Init() {

	// we register an sql driver txdb
	dns = fmt.Sprintf("host=%v port=%d user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbtest)

	txdb.Register("txdb", "postgres", dns)

}

func (ft *featureTest) resetResponse(*godog.Scenario) {

	handler := ft.GetHandler()
	ft.server = httptest.NewServer(handler)
	ft.resp = nil

}

func (ft *featureTest) JsonMarshaller(item Item) ([]byte, error) {

	js_item, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return js_item, nil

}

func (ft *featureTest) JsonUnMarshaller() (map[string]any, error) {

	var ResponseBody map[string]any

	if ft.resp == nil {
		log.Println("response is nil, cannot unmarshal", ft.resp)
		return nil, errors.New("response is nil, cannot unmarshal")
	}

	body, err := io.ReadAll(ft.resp.Body)

	if err != nil {

		return nil, err
	}

	defer ft.resp.Body.Close()

	log.Printf("Raw response body: %s", string(body)) // Log the raw JSON

	err = json.Unmarshal(body, &ResponseBody)

	if err != nil {

		return nil, err
	}

	return ResponseBody, nil

}

func (ft *featureTest) GetHandler() *gin.Engine {
	db, err := repository.NewDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	cart_rep := repository.NewCartRepository(db)

	cart_service := service.NewCartService(cart_rep)

	Router := gin.Default()
	cart_handler := handler.NewCartHandler(Router, cart_service)

	cart_handler.InitHandler()

	return Router

}

// add ,view and remove item step definition

func (ft *featureTest) RegisterItem(product_id, session_id, quantity int) error {

	if product_id != 0 && session_id != 0 && quantity != 0 {

		ft.item = Item{
			Product_id: product_id,
			Session_id: session_id,
			Quantity:   session_id,
		}

		return nil

	}

	return errors.New("invalid product ,session and quantity")
}

func (ft *featureTest) RegisterDeleteItem(P_id, S_id int) error {

	if P_id != 0 && S_id != 0 {

		ft.item = Item{
			Product_id: P_id,
			Session_id: S_id,
		}

		return nil
	}

	return errors.New("can't register item ")
}

func (ft *featureTest) ViewCartItem(session_id int) error {

	Url, err := url.Parse(ft.server.URL + "/cart")

	if err != nil {
		return err
	}

	params := url.Values{}

	params.Add("session_id", "1")

	Url.RawQuery = params.Encode()

	b, err := ft.JsonMarshaller(ft.item)

	if err != nil {
		return err
	}

	resp, err := ft.server.Client().Post(Url.String(), "application/json",
		bytes.NewBuffer(b))

	if err != nil {
		return err
	}
	ft.resp = resp
	return nil

}

// add new item step definition

func (ft *featureTest) AddCartItem() error {

	// adding new item to cart
	js, err := ft.JsonMarshaller(ft.item)
	if err != nil {
		return err
	}

	log.Print(ft.item)
	resp, err := ft.server.Client().Post(ft.server.URL+"/cart/item",
		"application/json", bytes.NewBuffer(js))

	if err != nil {
		log.Print(err)
		return err

	}
	ft.resp = resp

	return nil

}

func (ft *featureTest) AddItemResponse(want string) error {

	body, err := ft.JsonUnMarshaller()

	if err != nil {
		log.Println(err.Error())

		return err
	}

	got := body["message"]

	if got != want {

		err := fmt.Sprintf("The system expected %v , but got %v", want, got)
		return errors.New(err)
	}

	return nil

}

func (ft *featureTest) RemoveCartItem(S_id, P_id int) error {

	b, err := ft.JsonMarshaller(ft.item)

	if err != nil {
		return err
	}
	// removing cart item
	resp, err := ft.server.Client().Post(ft.server.URL+"/cart/remove",
		"application/json", bytes.NewBuffer(b))

	if err != nil {
		return err
	}
	ft.resp = resp

	return nil

}

// update item step definition

func (ft *featureTest) RegisterUpdateItem(S_id, P_id int) error {

	ft.item = Item{
		Session_id: S_id,
		Product_id: P_id,
	}

	return nil

}

func (ft *featureTest) DecreaseQuantity() error {

	b, err := ft.JsonMarshaller(ft.item)

	if err != nil {
		return err
	}

	resp, err := ft.server.Client().Post(ft.server.URL+"/cart/decrement", "application/json",
		bytes.NewBuffer(b))

	if err != nil {
		return err
	}
	ft.resp = resp

	return nil

}

func (ft *featureTest) IncreaseQuantity() error {

	b, err := ft.JsonMarshaller(ft.item)

	if err != nil {
		return err
	}

	resp, err := ft.server.Client().Post(ft.server.URL+"/cart/increment", "application/json",
		bytes.NewBuffer(b))

	if err != nil {

		log.Println("request is failing ", err)
		return err
	}

	if resp == nil {
		log.Println("can't update the quantity of the item  ", resp)
		return errors.New("can't update the quantity of the item ")
	}
	ft.resp = resp
	log.Println("Response successfully set:", ft.resp)
	return nil

}

func (ft *featureTest) SystemUpdateQuantity(want int) error {

	// log.Println("bd------------------", ft.resp.Body)

	body, err := ft.JsonUnMarshaller()

	log.Print("marshaller", body, err)

	if err != nil {
		log.Println("response errors", err)
		return err
	}

	log.Print("marshaller", body)
	got := body["row"]

	if floatVal, ok := got.(float64); ok {
		// Convert float64 to int, assuming no fractional part is expected
		intvalue := int(floatVal)

		if intvalue != want {
			errMsg := fmt.Sprintf("The system expected row to be %v, but got %v", want, intvalue)
			log.Println(errMsg)
			return errors.New(errMsg)
		}
	} else {
		// Handle the case where 'got' is not a number (int/float)
		return errors.New("can't convert body['row'] to integer")
	}

	return nil

}

func TestFeature(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}

}

func InitializeScenario(c *godog.ScenarioContext) {

	ft := featureTest{}
	// cart := steps.Cart{}

	c.Before(

		func(c context.Context, sc *godog.Scenario) (context.Context, error) {
			ft.resetResponse(sc)
			return c, nil
		})

	c.Step(`^I add item a new product with a (\d+),(\d+) and (\d+),$`, ft.RegisterItem)
	c.Step(`^the system should add  product item into cart$`, ft.AddCartItem)
	c.Step(`^the system should return "([^"]*)"\.$`, ft.AddItemResponse)
	// update step
	c.Step(`^I update item in cart with (\d+),  (\d+)  by increasing quantity,$`, ft.RegisterUpdateItem)
	c.Step(`^The system should increase the item quantity$`, ft.IncreaseQuantity)
	c.Step(`^the response should  return affected row (\d+)\.$`, ft.SystemUpdateQuantity)

	// decrease
	c.Step(`^I update item in cart with (\d+)  and (\d+) by decreasing  quantity$`, ft.RegisterUpdateItem)
	c.Step(`^The system should decrease the item quantity$`, ft.DecreaseQuantity)
	c.Step(`^the response should  return affected row (\d+)$`, ft.SystemUpdateQuantity)

}

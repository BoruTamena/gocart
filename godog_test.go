package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func (ft *featureTest) JsonMarshaller(item Item) []byte {

	js_item, _ := json.Marshal(item)

	return js_item

}

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
	// c.server = httptest.NewServer(handler)
	resp, err := ft.server.Client().Post(Url.String(), "application/json",
		bytes.NewBuffer(ft.JsonMarshaller(ft.item)))

	if err != nil {
		return err
	}
	ft.resp = resp
	return nil

}

func (ft *featureTest) AddCartItem(want string) error {

	// adding new item to cart

	response, err := ft.server.Client().Post(ft.server.URL+"/cart/item",
		"application/json", bytes.NewBuffer(ft.JsonMarshaller(ft.item)))

	if err != nil {
		return err
	}
	ft.resp = response
	return nil

}

func (ft *featureTest) RemoveCartItem(S_id, P_id int) error {
	// removing cart item

	resp, err := ft.server.Client().Post(ft.server.URL+"/cart/remove",
		"application/json", bytes.NewBuffer(ft.JsonMarshaller(ft.item)))

	if err != nil {
		return err
	}
	ft.resp = resp
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

	c.Step(`^I add item a new product with a product_id,session_id and quantity (\d+),(\d+),(\d+),$`, ft.RegisterItem)
	c.Step(`^the system should add a new product into cart and return "([^"]*)"\.$`, ft.AddCartItem)

}

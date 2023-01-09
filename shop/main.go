package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmdavril/pubsub"
	catalogapi "github.com/jmdavril/template/shop/context/catalog/api"
	catalogdata "github.com/jmdavril/template/shop/context/catalog/data"
	catalogdomain "github.com/jmdavril/template/shop/context/catalog/domain"
	shopapi "github.com/jmdavril/template/shop/context/shop/api"
	shopdata "github.com/jmdavril/template/shop/context/shop/data"
	shopdomain "github.com/jmdavril/template/shop/context/shop/domain"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

func main() {
	run()
}

func run() {

	////////
	// data
	////////
	var db = connectDb()

	var customerRepo = shopdata.NewCustomerRepo(db)
	var orderRepo = shopdata.NewOrderRepo(db)
	var productSalesRepo = shopdata.NewProductSalesRepo(db)

	var productRepo = catalogdata.NewProductRepo(db)

	//////////
	// domain
	//////////
	var publisher = pubsub.NewPublisher()

	var customerService = shopdomain.NewCustomerService(customerRepo, orderRepo, productSalesRepo)
	var productSalesService = shopdomain.NewProductSalesService(productSalesRepo)
	shopdomain.StartSubscriptions(publisher, productSalesService)
	var productService = catalogdomain.NewProductService(productRepo, publisher)

	///////
	// api
	///////
	engine := gin.Default()

	var shopController = shopapi.NewShopController(engine, customerService)
	var catalogController = catalogapi.NewCatalogController(engine, productService)

	shopController.Run()
	catalogController.Run()

	return
}

func connectDb() *sqlx.DB {
	conn := os.Getenv("PG_CONN")
	db, err := sqlx.Open("postgres", conn)

	var retryErr error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		retryErr = db.Ping()
		if retryErr == nil {
			break
		}
		log.Println(retryErr)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if retryErr != nil {
		log.Fatalln("Error connecting to DB", err)
	}

	return db
}

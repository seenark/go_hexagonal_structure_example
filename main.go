package main

import (
	"bank/config"
	"bank/handler"
	"bank/repository"
	"bank/service"
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config := config.GetConfig()
	initTimeZone()
	mongoPath := fmt.Sprintf("mongodb+srv://%v:%v@hdgcluster.xmgsx.mongodb.net/golearn?retryWrites=true&w=majority", config.Mongo.Username, config.Mongo.Password)
	connection := options.Client().ApplyURI(mongoPath)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}

	db := client.Database("hexagonal")
	collection := db.Collection("customers")

	customerRepository := repository.NewCustomerRepositoryDB(collection, ctx)
	_ = customerRepository
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	e := echo.New()
	e.GET("/customers", customerHandler.GetAllCustomer)
	e.GET("/customer/:id", customerHandler.GetCustomerById)

	e.Start(fmt.Sprintf(":%v", config.App.Port))

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

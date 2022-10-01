package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stefanlester/ecommerce-go/controllers"
	"github.com/stefanlester/ecommerce-go/database"
	"github.com/stefanlester/ecommerce-go/middleware"
	"github.com/stefanlester/ecommerce-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())

	router.POST("/addaddress", controllers.AddAddress())

	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())

	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	
	log.Fatal(router.Run(":" + port))
}
 
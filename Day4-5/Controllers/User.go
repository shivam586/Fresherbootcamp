//Controllers/User.go

package Controllers

import (
	"HypotheticalRetailer/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var a = 0

func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"customer_id":customer.CustomerId, "name":customer.Name, "phone":customer.Phone})
	}

}
//Get all available products data'


func GetAllAvailableProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllAvailableProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
//Create an order request for a customer
func CreateOrderRequest(c *gin.Context) {
	var productordered Models.Order
	done := make(chan bool,1)
	c.BindJSON(&productordered)

	err := Models.CreateOrderRequest(&productordered)
    if a>0 {
    	done<-true
    	fmt.Println("wait")
    	time.Sleep(6*time.Second)
    	<-done
	}
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, productordered)

	}
    a = a+1
}
//Get a product with given id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("product_id")
	var productofgivenid Models.Product
	err := Models.GetProductByID(&productofgivenid, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, productofgivenid)
	}
}
//get customer of given id orders history
func GetCustomersOrdersHistory(c *gin.Context) {
	var customerorder []Models.Order

	id := c.Params.ByName("customer_id")
	err := Models.GetCustomersOrdersHistory(&customerorder, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customerorder)
	}
}

func GetAllTransactions(c *gin.Context) {
	var transactions []Models.Order
	err := Models.GetAllTransactions(&transactions)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transactions)
	}
}

func AddProduct(c *gin.Context) {
	var producttobeadded Models.Product
	c.BindJSON(&producttobeadded)
	err := Models.AddProduct(&producttobeadded)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"product_id":producttobeadded.ProductId,"name":producttobeadded.Name,"quantity":producttobeadded.Quantity,
		                            "price":producttobeadded.Price,"message":"Product successfully added"})
	}
}

//Update product partially

func PartialUpdate(c *gin.Context) {
	var modify Models.Product
	id := c.Params.ByName("product_id")
	err := Models.GetProductByID(&modify, id)
	if err != nil {
		c.JSON(http.StatusNotFound, modify)
	}
	  c.BindJSON(&modify)
	err = Models.PartialUpdate(&modify, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, modify)
	}
}
func DeleteProduct(c *gin.Context) {
	var producttobedeleted Models.Product
	id := c.Params.ByName("product_id")
	err :=Models.DeleteProduct(&producttobedeleted, id)
	if err !=nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"product_id" + id: "is deleted"})
	}

}
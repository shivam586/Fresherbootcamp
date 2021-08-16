//Routes/Routes.go

package Routes

import (
	"HypotheticalRetailer/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes

func SetupRouter() *gin.Engine {
	r := gin.Default()
	customergroup := r.Group("/customer-retailerapi")
	{
		customergroup.POST("createcustomer",Controllers.CreateCustomer)
		customergroup.GET("showallproducts", Controllers.GetAllAvailableProducts)
		customergroup.POST("orderrequest/:customer_id", Controllers.CreateOrderRequest)
		customergroup.GET("product/:product_id", Controllers.GetProductByID)
		customergroup.GET("orderhistory/:customer_id", Controllers.GetCustomersOrdersHistory)

	}
	retailergroup := r.Group("/retailer-specific")
	{
		retailergroup.GET("alltransaction", Controllers.GetAllTransactions)
		retailergroup.POST("retailer/addproduct", Controllers.AddProduct)
		retailergroup.PUT("retailer/partialupdate/:product_id", Controllers.PartialUpdate)
		retailergroup.DELETE("retailer/deleteprodcut/:product_id", Controllers.DeleteProduct)

	}
	return r
}
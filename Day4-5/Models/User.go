//Models/User.go

package Models

import (
	"HypotheticalRetailer/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
//Get all available products

func GetAllAvailableProducts(products *[]Product) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}
//Create order request

func CreateOrderRequest(order *Order) (err error) {

	if err = Config.DB.Model(&order).Create(order).Error; err != nil {
		return err
	}
	return nil

}


func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("prodcut_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomersOrdersHistory(customerorder *[]Order,id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).Find(&customerorder).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTransactions(order *[]Order) (err error) {
	if err = Config.DB.Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func PartialUpdate(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("product_id = ?", id).Delete(product)
	return nil
}


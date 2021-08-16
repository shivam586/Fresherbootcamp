//Models/Order.go

package Models

type Order struct {
	OrderId        uint      `json:"order_id" gorm:"primary_key"`
	CustomerId     uint      `json:"customer_id"`
	ProductId      uint      `json:"product_id"`
	Quantity       int       `json:"quantity"`
	Price          float32   `json:"price"`
}
func (o *Order) TableName() string {
	return "ordered"
}
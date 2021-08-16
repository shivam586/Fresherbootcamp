//Models/Product.go

package Models

type Product struct {
	ProductId             uint      `json:"product_id" gorm:"primary_key"`
	Name           string    `json:"name"`
	Quantity       int       `json:"quantity"`
	Price          float32   `json:"price"`

}
func (c *Product) TableName() string {
	return "product"
}
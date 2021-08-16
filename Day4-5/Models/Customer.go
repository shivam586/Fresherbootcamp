//Models/Customer.go

package Models

type Customer struct {
	CustomerId             uint    `json:"customer_id" gorm:"primary_key"`
	Name                   string  `json:"name"`
	Phone                  string  `json:"phone"`
}
func (b *Customer) TableName() string {
	return "customer"
}
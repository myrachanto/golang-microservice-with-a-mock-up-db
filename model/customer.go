package model

import(
	"time"
	"github.com/myrachanto/amicroservice/customermicroservice/httperors"
)

type Customer struct {
	ID int `json:"id" bson:"id"`
	Name string `gorm:"not null" bson:"name"`
	Company string `gorm:"not null" bson:"company"`
	Phone string `gorm:"not null" bson:"phone"`
	Address string `gorm:"not null" bson:"address"`
	Email string `gorm:"not null;unique" bson:"email"`
	Invoice []Invoice `gorm:"foreignkey:UserRefer"`//has many invoices
	Base
}
type Invoice struct {
	ID int `json:"id"`
	CustomerID uint64 `gorm:"not null" bson:"customer_id"`
	Customer Customer `gorm:"foreignKey:CustomerID; not null"`
	Title string `gorm:"not null" bson:"title"`
	Dated time.Time `gorm:"not null" bson:"date"`
	Due_date time.Time `gorm:"not null" bson:"due_date"`
	Discount float64 `gorm:"not null" bson:"discount"`
	Sub_total float64 `gorm:"not null" bson:"sub_total"`
	Total float64 `gorm:"not null" bson:"total"`
	InvoiceItem []InvoiceItem `gorm:"foreignkey:UserRefer"`//has many invoiceitems
	Base
}
type InvoiceItem struct {
	ID int `json:"id"`
	InvoiceID uint64 `gorm:"not null" bson:"invoice_id"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID; not null"`
	Description string `gorm:"not null" bson:"description"`
	Qty uint64 `gorm:"not null" bson:"qty"`
	Unit_price float64 `gorm:"not null" bson:"unit_price"`
	Base
}
type Base struct{
	Created_At time.Time `gorm:"created_at" bson:"created_at"`
	Updated_At time.Time `gorm:"updated_at" bson:"updated_at"`
	Delete_At *time.Time `gorm:"deleted_at" bson:"deleted_at"`

}

/*

        'id','company', 'email', 'name', 'phone', 'address', 'created_at'

        'company', 'email', 'name', 'phone', 'address'
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `gorm:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}*/
func (customer Customer) Validate() *httperors.HttpError{
	if customer.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if customer.Company == "" {
		return httperors.NewNotFoundError("Invalid Company")
	}
	if customer.Phone == "" {
		return httperors.NewNotFoundError("Invalid Phone")
	}
	if customer.Email == "" {
		return httperors.NewNotFoundError("Invalid Email")
	}
	if customer.Address == "" {
		return httperors.NewNotFoundError("Invalid Address")
	}
	return nil
}
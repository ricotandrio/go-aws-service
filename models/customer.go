package models

type Customers struct {
	CustomerID  			int    		`json:"customer_id" gorm:"primaryKey"`
	CustomerName  		string 		`json:"first_name"`
	CustomerEmail  		string 		`json:"email"`
	CustomerPhone  		string 		`json:"phone"`
	CustomerAddress 	string 		`json:"address"`
	CustomerCreatedAt string 		`json:"created"`
}
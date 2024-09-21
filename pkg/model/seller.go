package model

type Seller struct {
	ID       string  `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Email    string  `json:"email" gorm:"not null;unique"`
	Password string  `json:"password" gorm:"not null"`
	ShopName string  `json:"shop_name" gorm:"not null"`
	Balance  float64 `json:"balance" gorm:"default:0"`
	Items    []Item  `json:"items" gorm:"foreignKey:SellerID"`
}

type SellerOutput struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	ShopName string     `json:"shop_name"`
	Balance  float64    `json:"balance" gorm:"default:0"`
	Items    []ItemInfo `json:"items"`
}

type SellerSignInResponse struct {
	Token  string       `json:"token"`
	Seller SellerOutput `json:"seller"`
}

package models

type Payment struct {
	Id          int     `gorm:"primaryKey" json:"id"`
	Amount      float64 `gorm:"column:amount" json:"amount"`
	PaymentDate string  `gorm:"column:payment_date" json:"paymentDate"`
	PaymentMode string  `gorm:"column:payment_mode" json:"paymentMode"`
	IsPaid      bool    `gorm:"column:is_paid" json:"isPaid"`
	RPOrderId   string  `gorm:"column:rp_order_id" json:"rpOrderId"`
	Status      string  `gorm:"column:status" json:"status"`
}

func (Payment) TableName() string {
	return "payments"
}

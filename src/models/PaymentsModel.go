package models

type Payment struct {
	Id          int     `gorm:"primaryKey"`
	Amount      float64 `gorm:"column:amount"`
	PaymentDate string  `gorm:"column:payment_date"`
	PaymentMode string  `gorm:"column:payment_mode"`
	IsPaid      bool    `gorm:"column:is_paid"`
	RPOrderId   string  `gorm:"column:rp_order_id"`
	Status      string  `gorm:"column:status"`
}

func (Payment) TableName() string {
	return "payments"
}

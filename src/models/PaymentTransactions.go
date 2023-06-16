package models

type PaymentTransaction struct {
	Id          int    `gorm:"column:id;primary_key"`
	PaymentId   int    `gorm:"column:payment_id"`
	RPPaymentId string `gorm:"column:rp_payment_id"`
	Amount      int    `gorm:"column:amount"`
	IsSuccess   bool   `gorm:"column:is_success"`
	Status      string `gorm:"column:status"`
	MetaData    string `gorm:"column:meta_data"`
}

func (PaymentTransaction) TableName() string {
	return "payment_transactions"
}

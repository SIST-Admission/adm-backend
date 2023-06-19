package models

type PaymentTransaction struct {
	Id          int    `gorm:"column:id;primary_key" json:"id"`
	PaymentId   int    `gorm:"column:payment_id" json:"paymentId"`
	RPPaymentId string `gorm:"column:rp_payment_id" json:"rpPaymentId"`
	Amount      int    `gorm:"column:amount" json:"amount"`
	IsSuccess   bool   `gorm:"column:is_success" json:"isSuccess"`
	Status      string `gorm:"column:status" json:"status"`
	MetaData    string `gorm:"column:meta_data" json:"metaData"`
}

func (PaymentTransaction) TableName() string {
	return "payment_transactions"
}

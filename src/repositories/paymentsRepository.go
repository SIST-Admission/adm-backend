package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type PaymentsRepository struct{}

func (repo *PaymentsRepository) GetPayment(paymentId int) (*models.Payment, error) {
	logrus.Info("PaymentsRepository.GetPayment")
	logrus.Info("User: ", paymentId)
	db := db.GetInstance()
	var payment models.Payment
	if err := db.Where("id = ?", paymentId).First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (repo *PaymentsRepository) GetPaymentByUserId(userId int) (*models.Payment, error) {
	logrus.Info("PaymentsRepository.GetPaymentByUserId")
	logrus.Info("User: ", userId)
	db := db.GetInstance()
	var user models.User

	if err := db.Model(&models.User{}).Where("id = ?", userId).Preload("Application").Preload("Application.PaymentDetails").First(&user).Error; err != nil {
		return nil, err
	}

	return user.Application.PaymentDetails, nil
}

func (repo *PaymentsRepository) CreatePayment(applicationId int, payment *models.Payment) (*models.Payment, error) {
	logrus.Info("PaymentsRepository.CreatePayment")
	db := db.GetInstance()

	if err := db.Create(&payment).Error; err != nil {
		return nil, err
	}
	err := repo.AddPaymentToApplication(applicationId, payment.Id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (repo *PaymentsRepository) AddPaymentToApplication(appId, paymentId int) error {
	logrus.Info("PaymentsRepository.AddPaymentToApplication")
	logrus.Info("Application Id: ", appId)
	logrus.Info("Payment Id: ", paymentId)

	db := db.GetInstance()
	if err := db.Model(&models.Application{}).Where("id = ?", appId).Update("payment_id", paymentId).Error; err != nil {
		logrus.Error("Failed to update application with payment id: ", err)
		return err
	}

	return nil
}

func (repo *PaymentsRepository) UpdatePaymentStatusByOrderId(orderId, status string, isPaid bool) error {
	logrus.Info("PaymentsRepository.UpdatePayment")
	db := db.GetInstance()

	if err := db.Model(models.Payment{}).Where("rp_order_id = ?", orderId).Updates(models.Payment{
		IsPaid: isPaid,
		Status: status,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PaymentsRepository) AddTransaction(transaction *models.PaymentTransaction) (*models.PaymentTransaction, error) {
	logrus.Info("PaymentsRepository.AddTransaction")
	db := db.GetInstance()

	if err := db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

package service

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	razorpay "github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PaymentsService struct{}

var paymentsRepository repositories.PaymentsRepository = repositories.PaymentsRepository{}
var userRepository repositories.UserRepository = repositories.UserRepository{}

func (paymentsService *PaymentsService) GetOrder(userId int) (map[string]interface{}, *dto.Error) {
	logrus.Info("PaymentsService.GetOrder")
	logrus.Info("User: ", userId)

	user, err := userRepository.GetUserById(strconv.Itoa(userId))
	if err != nil {
		logrus.Error("Failed to get user: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get user",
		}
	}

	if user == nil {
		return nil, &dto.Error{
			Code:    404,
			Message: "User not found",
		}
	} else if user.ApplicationId == 0 {
		return nil, &dto.Error{
			Code:    400,
			Message: "Application not found",
		}
	}

	// Get Payment by user id
	paymentDetails, err := paymentsRepository.GetPaymentByUserId(userId)
	if err != nil {
		logrus.Error("Failed to get payment Details: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get payment Details",
		}
	}

	logrus.Info("Payment Details: ", paymentDetails)

	key := viper.GetString(viper.GetString("env") + ".razorpay.key")
	secret := viper.GetString(viper.GetString("env") + ".razorpay.secret")
	amount := viper.GetInt(viper.GetString("env") + ".application.fee")
	logrus.Info("Razorpay Key: ", key)
	logrus.Info("Razorpay Secret: ", secret)

	// Razorpay client
	client := razorpay.NewClient(key, secret)

	if paymentDetails != nil {
		logrus.Info("Fetching existing order details")
		orderDetails, err := client.Order.Fetch(paymentDetails.RPOrderId, nil, nil)
		if err != nil {
			logrus.Error("Failed to fetch existing order: ", err)
			return nil, &dto.Error{
				Code:    500,
				Message: "Failed to fetch order",
			}
		}
		orderDetails["payment_id"] = paymentDetails.Id
		orderDetails["application_id"] = user.ApplicationId
		orderDetails["user_id"] = userId
		return orderDetails, nil
	}

	// Order Details Doesn't exist, create new order
	data := map[string]interface{}{
		"amount":          amount * 100,
		"currency":        "INR",
		"receipt":         "app_fee_" + strconv.Itoa(userId) + "_" + strconv.Itoa(user.ApplicationId) + "_" + strconv.Itoa(int(time.Now().Unix())),
		"partial_payment": false,
		"notes": map[string]interface{}{
			"application_id": user.ApplicationId,
			"user_id":        userId,
		},
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		logrus.Error("Failed to create order: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to create order",
		}
	}

	payment, err := paymentsRepository.CreatePayment(user.ApplicationId, &models.Payment{
		Amount:      body["amount"].(float64),
		PaymentDate: time.Now().Format("2006-01-02 15:04:05"),
		PaymentMode: "online",
		IsPaid:      false,
		RPOrderId:   body["id"].(string),
	})

	if err != nil {
		logrus.Error("Failed to create payment: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to create payment",
		}
	}

	logrus.Info("Payment: ", payment)
	body["payment_id"] = payment.Id
	body["application_id"] = user.ApplicationId
	body["user_id"] = userId
	return body, nil
}

func (paymentsService *PaymentsService) VerifyPayment(payload, signature string) error {
	logrus.Info("PaymentsService.VerifyPayment")
	var paymentDetails map[string]interface{}

	secret := viper.GetString(viper.GetString("env") + ".razorpay.secret")

	// Verify Signature
	if utils.VerifyWebhookSignature(payload, signature, secret) {
		logrus.Info("Signature verified")
	} else {
		logrus.Error("Signature verification failed")
	}
	// Parse payload
	if err := json.Unmarshal([]byte(payload), &paymentDetails); err != nil {
		logrus.Error("Failed to parse payment details: ", err)
		return err
	}

	return nil
}

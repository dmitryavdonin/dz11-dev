package broker

import (
	"encoding/json"
	"fmt"
	"notification/internal/model"
	"notification/internal/service"
	"time"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type PaymentStatusEvent struct {
	Data struct {
		OrderId int    `json:"order_id"`
		UserId  int    `json:"user_id"`
		Status  string `json:"status"`
		Reason  string `json:"reason"`
	} `json:"data"`
}

type PaymentStatusHandler struct {
	services           *service.Services
	paymentStatusTopic string
}

func BuildPaymentStatusHandler(services *service.Services, paymentStatusTopic string) PaymentStatusHandler {
	return PaymentStatusHandler{
		services:           services,
		paymentStatusTopic: paymentStatusTopic,
	}
}

func (psh PaymentStatusHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	logrus.Printf("ConsumeClaim(): BEGIN consuming messages from topic = %s", psh.paymentStatusTopic)

	for msg := range claim.Messages() {
		pse := PaymentStatusEvent{}
		err := json.Unmarshal(msg.Value, &pse)
		if err != nil {
			logrus.Errorf("ConsumeClaim(): Event hasn't been handled, error =  %s", err.Error())
			session.MarkMessage(msg, "")
			continue
		}

		logrus.Printf("ConsumeClaim(): Message received: order_id = %d, user_id = %d, status = %s, reason = %s",
			pse.Data.OrderId, pse.Data.UserId, pse.Data.Status, pse.Data.Reason)

		// create notification with the results of payment status
		logrus.Printf("ConsumeClaim(): Try to create notification for order_id = %d, user_id = %d, status = %s, reason = %s",
			pse.Data.OrderId, pse.Data.UserId, pse.Data.Status, pse.Data.Reason)

		n := model.Notification{
			UserId:     pse.Data.UserId,
			OrderId:    pse.Data.OrderId,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}

		if pse.Data.Status == "success" {
			n.Message = fmt.Sprintf("Order status = %s", pse.Data.Status)
		} else {
			n.Message = fmt.Sprintf("Order status = %s, reason = %s", pse.Data.Status, pse.Data.Reason)
		}

		id, err := psh.services.Notification.Create(n)
		if err != nil {
			logrus.Errorf("ConsumeClaim(): Cannot create notification for order_id = %d, error =  %s", pse.Data.OrderId, err.Error())
			session.MarkMessage(msg, "")
			continue
		}

		logrus.Printf("ConsumeClaim(): END notification id = %d for order_id = %d created", id, pse.Data.OrderId)

		session.MarkMessage(msg, "")
	}

	return nil
}

func (PaymentStatusHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (PaymentStatusHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

package utils

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go_fiber/logger"
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetConfig(config string) []byte {

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(config), &data); err != nil {
		logger.Error("error formatting json :: ", zap.Error(err))
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("Error encoding JSON:", zap.Error(err))
	}

	return jsonBytes
}

func GenerateOTP() string {
	rand.NewSource(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		code += fmt.Sprintf("%d", digit)
	}

	return code
}

func SendEmail(recipient string) {
	m := gomail.NewMessage()

	Port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m.SetAddressHeader("From", os.Getenv("EMAIL"), os.Getenv("EMAIL_SENDER_NAME"))

	m.SetHeader("To", recipient)

	m.SetHeader("Subject", "Test Email")

	m.SetBody("text/html", "Hello from Golang! This is the email content.")

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), Port, os.Getenv("EMAIL"), os.Getenv("EMAIL_PWD"))

	if err := d.DialAndSend(m); err != nil {
		logger.Error("error sending email :: ", zap.Error(err))
	}
}

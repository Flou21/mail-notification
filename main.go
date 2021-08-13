package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		log.Println("continue without")
	}

	config := &Config{}
	readConfig(config)
	sendMail(config)
}

func sendMail(config *Config) {

	// Message.
	message := []byte(config.Message)

	// Authentication.
	auth := smtp.PlainAuth("", config.SenderName, config.SenderPassword, config.SmtpHost)

	// Sending email.
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.SenderName, config.Recipients, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func readConfig(config *Config) {
	if config.SmtpHost = os.Getenv("SMTP_HOST"); config.SmtpHost == "" {
		log.Fatalln("SMTP_HOST not set")
	}

	if config.SmtpPort = os.Getenv("SMTP_PORT"); config.SmtpPort == "" {
		log.Fatalln("SMTP_PORT not set")
	}

	if config.SenderName = os.Getenv("SENDER_NAME"); config.SenderName == "" {
		log.Fatalln("SENDER_NAME not set")
	}

	if config.SenderPassword = os.Getenv("SENDER_PASSWORD"); config.SenderPassword == "" {
		log.Fatalln("SENDER_PASSWORD not set")
	}

	if config.Message = os.Getenv("MESSAGE"); config.Message == "" {
		log.Fatalln("MESSAGE not set")
	}

	recipients := os.Getenv("RECIPIENTS")

	if config.Recipients = strings.Split(recipients, ","); len(config.Recipients) <= 0 {
		log.Fatalln("Recipients not set")
	}
}

type Config struct {
	SenderName     string
	SenderPassword string
	Recipients     []string
	SmtpHost       string
	SmtpPort       string
	Message        string
}

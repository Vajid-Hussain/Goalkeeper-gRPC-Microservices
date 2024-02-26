package utils

import (
	"log"

	"github.com/vajid-hussain/grpc-microservice-mail-svc/pkg/model"
	"github.com/wneessen/go-mail"
)

// func RemainderMain(credential *model.SendRemainder) {
// 	m := mail.NewMsg()
// 	if err := m.From(credential.From); err != nil {
// 		log.Fatalf("failed to set From address: %s", err)
// 	}

// 	if err := m.To(credential.To); err != nil {
// 		log.Fatalf("failed to set To address: %s", err)
// 	}
// 	m.Subject("Welcome to GoalKeeper!")

// 	// Customize the email body with a welcoming message
// 	body := "Dear " + credential.To + ",\n\n"
// 	body += "Thank you for signing up for GoalKeeper! We're excited to have you on board.\n"
// 	body += "With GoalKeeper, you can create categories, add data, set reminders, and stay organized.\n\n"
// 	body += "Get started now and make the most out of GoalKeeper!\n\n"
// 	body += "Best regards,\n"
// 	body += "The GoalKeeper Team"

// 	m.SetBodyString(mail.TypeTextPlain, body)

// 	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
// 		mail.WithUsername(credential.From), mail.WithPassword(credential.AppPasskey))
// 	if err != nil {
// 		log.Fatalf("failed to create mail client: %s", err)
// 	}

// 	if err := c.DialAndSend(m); err != nil {
// 		log.Fatalf("failed to send mail: %s", err)
// 	}
// }

func RemainderMain(credential *model.SendRemainder) {
	// Create a new message
	m := mail.NewMsg()

	// Set the sender address
	if err := m.From(credential.From); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}

	// Set the recipient address
	if err := m.To(credential.To); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

	// Set the email subject
	m.Subject("Reminder for Today!")

	// Customize the email body with the reminders
	body := "Dear " + credential.To + ",\n\n"
	body += "Here are your reminders for today:\n\n"

	// Append each reminder to the email body
	for _, reminder := range credential.Data {
		body += "- " + reminder + "\n"
	}

	// Add a closing message
	body += "\nBest regards,\nThe GoalKeeper Team"

	// Set the email body
	m.SetBodyString(mail.TypeTextPlain, body)

	// Create a mail client
	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(credential.From), mail.WithPassword(credential.AppPasskey))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	// Send the email
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}

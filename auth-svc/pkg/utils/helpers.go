package utils

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	reqestmodel "github.com/vajid-hussain/grpc-microservice-auth-svc/pkg/models/reqestModel"
	"github.com/wneessen/go-mail"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hastpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hastpassword), nil
}

func CompareHast(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}

func CreateJwt(secret string, userID string) string {
	claim := jwt.MapClaims{
		"userid": userID,
		"exp":    time.Now().Unix() + 3600000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, _ := token.SignedString([]byte(secret))

	return signedToken
}

func VerifyJwt(token, secret string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("token is not valid")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	id, _ := claims["userid"].(string)

	return id, nil
}

// func SendMail(credential reqestmodel.Email) {
// 	m := mail.NewMsg()
// 	if err := m.From(credential.From); err != nil {
// 		log.Fatalf("failed to set From address: %s", err)
// 	}

// 	if err := m.To(credential.To); err != nil {
// 		log.Fatalf("failed to set To address: %s", err)
// 	}
// 	m.Subject("This is my first mail with go-mail!")
// 	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")

// 	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
// 		mail.WithUsername(credential.From), mail.WithPassword(credential.AppPasskey))
// 	if err != nil {
// 		log.Fatalf("failed to create mail client: %s", err)
// 	}

// 	if err := c.DialAndSend(m); err != nil {
// 		log.Fatalf("failed to send mail: %s", err)
// 	}
// }

func GreetingMain(credential reqestmodel.Email) {
    m := mail.NewMsg()
    if err := m.From(credential.From); err != nil {
        log.Fatalf("failed to set From address: %s", err)
    }

    if err := m.To(credential.To); err != nil {
        log.Fatalf("failed to set To address: %s", err)
    }
    m.Subject("Welcome to GoalKeeper!")

    // Customize the email body with a welcoming message
    body := "Dear " + credential.To + ",\n\n"
    body += "Thank you for signing up for GoalKeeper! We're excited to have you on board.\n"
    body += "With GoalKeeper, you can create categories, add data, set reminders, and stay organized.\n\n"
    body += "Get started now and make the most out of GoalKeeper!\n\n"
    body += "Best regards,\n"
    body += "The GoalKeeper Team"

    m.SetBodyString(mail.TypeTextPlain, body)

    c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
        mail.WithUsername(credential.From), mail.WithPassword(credential.AppPasskey))
    if err != nil {
        log.Fatalf("failed to create mail client: %s", err)
    }

    if err := c.DialAndSend(m); err != nil {
        log.Fatalf("failed to send mail: %s", err)
    }
}

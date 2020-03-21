package main

import (
	"fmt"
	"log"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(`fd4682de54a16a42fba0c1778ff4baec`, `4d59681120a1e6f086722e24c63db5ff`)

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				//			Email: "rajesh.jain@innutech.com",
				Email: "rajeshkumjain@gmail.com",
				Name:  "Rajesh",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "rajeshkumjain@gmail.com",
					Name:  "Rajesh",
				},
			},
			Subject:  "Greetings from Mailjet.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

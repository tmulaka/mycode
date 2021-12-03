/* RZFeeser | Alta3 Research
   Sending an Email (SMTP) message   */
   
package main

import (
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "talk2tara@gmail.com"            // update this to reflect your value
    password := "test from GoLang Training"     // update this to reflect your value
    to := []string{"tara.mulakaluri@optum.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // update this to be your message body 
    message := []byte ("Subject: Testing email from go\r\n\n.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
}


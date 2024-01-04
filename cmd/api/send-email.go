package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
)

type EmailInfo struct {
	To   		string  `json:"to"`
	Subject     string  `json:"subject"`
	Data  		float64 `json:"data"`
}

const (
    Sender = "danjavia@gmail.com"

    Recipient = "danjavia@gmail.com"
    
    Subject = "Prueba de envio de summary"
    
    HtmlBody =  "<h1>Welcome to summary " +
                "<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
                "<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
    
    // This field is for emiails without html support.
    TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
    
    CharSet = "UTF-8"
)

func SendEmail(c *gin.Context) {
	var emailInfo EmailInfo

	err := json.NewDecoder(c.Request.Body).Decode(&emailInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error decoding JSON: " + err.Error(),
		})
		return
	}

    sess, err := session.NewSession(&aws.Config{
        Region:aws.String("us-east-1")},
    )
    
    // Create an SES session.
    svc := ses.New(sess)
    
    // Assemble the email.
    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{
            },
            ToAddresses: []*string{
                aws.String(Recipient),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Html: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(HtmlBody),
                },
                Text: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(TextBody),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(CharSet),
                Data:    aws.String(Subject),
            },
        },
        Source: aws.String(Sender),
    }

    // Send email
    result, err := svc.SendEmail(input)
    
    // Display error messages if they occur.
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case ses.ErrCodeMessageRejected:
                fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
            case ses.ErrCodeMailFromDomainNotVerifiedException:
                fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
            case ses.ErrCodeConfigurationSetDoesNotExistException:
                fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
    
        return
    }
    
    fmt.Println("Email Sent to address: " + Recipient)
    fmt.Println(result)
}
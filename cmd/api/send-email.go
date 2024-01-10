package api

import (
    "database/sql"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
    _ "github.com/lib/pq" 
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

type EmailInfo struct {
	To   		string  `json:"to"`
	Subject     string  `json:"subject"`
	Data  		float64 `json:"data"`
}

type EmailRequestBody struct {
    ReceiverEmail string            `json:"receiverEmail"`
    SenderEmail   string            `json:"senderEmail"`
    TemplateName  string            `json:"templateName"`
    Placeholders  map[string]string `json:"placeholders"`
}

func sendTemplatedEmail(client *ses.Client, input *ses.SendTemplatedEmailInput) (string, error) {
    output, err := client.SendTemplatedEmail(context.Background(), input)
    if err != nil {
        errorMessage := fmt.Sprintf(`{"error_message": "%s"}`, err.Error())
        return "", fmt.Errorf(errorMessage)
    }
    return *output.MessageId, nil
}

func SendEmail(db *sql.DB) gin.HandlerFunc {
    return func (c *gin.Context) {
        var emailInfo EmailInfo

        err := json.NewDecoder(c.Request.Body).Decode(&emailInfo)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  "error",
                "message": "Error decoding JSON: " + err.Error(),
            })
            return
        }

        sess, err := config.LoadDefaultConfig(context.TODO(),
            config.WithRegion(os.Getenv("AWS_REGION")),
        )
        
        // Create an SES session.
        client := ses.NewFromConfig(sess)


        templateData, err := json.Marshal(map[string]interface{}{
            "USER_EMAIL": "Danjavia@gmail.com",
            "TOTAL_AMOUNT": "+580",
            "AVG_CREDIT": "+678",
            "AVG_DEBIT": "-67",
        })
        
        // Assemble the email.
        input := &ses.SendTemplatedEmailInput{
            Source:       aws.String("danjavia@gmail.com"),
            Destination:  &types.Destination{ToAddresses: []string{"danjavia@gmail.com"}},
            Template:     aws.String("STORI_TMPL"),
            TemplateData: aws.String(string(templateData)),
        }

        // Send email
        messageId, err := sendTemplatedEmail(client, input)

        successMessage := fmt.Sprintf("Message successfully sent with Message ID: %s", messageId)
        
        // Display error messages if they occur.
        if err != nil {
            fmt.Println(err.Error())
        
            return
        }
        
        fmt.Println("Email Sent to address: " + "danjavia@gmail.com")
        fmt.Println(successMessage)
    }
}
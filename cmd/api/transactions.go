package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/danjavia/stori_csv/cmd/infraestructure/db"
	"github.com/danjavia/stori_csv/cmd/infraestructure/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Transaction struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Transaction float64 `json:"transaction"`
}


func CheckTransactions(client *dynamodb.Client) gin.HandlerFunc {
	return func (c *gin.Context) {
		var transactions []Transaction
	
		err := json.NewDecoder(c.Request.Body).Decode(&transactions)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Error decoding JSON: " + err.Error(),
			})
			return
		}
	
	
		// Calculate desired metrics
		var totalAmount float64
		monthCounts := make(map[string]int)
		var totalCredit, totalDebit float64
	
	
		for _, transaction := range transactions {
			totalAmount += transaction.Transaction
			date, _ := time.Parse("1/2", transaction.Date)
			month := date.Month().String()
			monthCounts[month]++ // Sum Transaction according to month
			if transaction.Transaction < 0 {
				totalCredit += transaction.Transaction
			} else {
				totalDebit += transaction.Transaction
			}
		}
	
		// Output Structure
		output := map[string]interface{}{
			"totalAmount": totalAmount,
			"monthCounts": monthCounts,
			"avgCredit": totalCredit / float64(len(transactions)), // Assuming non-zero transactions
			"avgDebit": totalDebit / float64(len(transactions)), // Assuming non-zero transactions
		}

		summary := &models.Summary{
			ID:          uuid.New().String(),
			UserId:      "",
			UserEmail:   "",
			Summary:     "",
			ArtifactUrl: "",
		}


		// save data on DB
		db.CreateSummary(c, client, summary)
	
	
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": output, // this value will change after grouped data correctly for sending email via sendgrid
		})
	}
  }


package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/danjavia/stori_csv/cmd/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var ginLambda *ginadapter.GinLambda


func init() {
	log.Printf("Gin cold start")

	// Initialize configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"), // Replace with your desired region
	)

	if err != nil {
		// Handle configuration loading errors
		panic(err) // Or use a more graceful error handling mechanism
	}

	// Create DynamoDB client
	client := dynamodb.NewFromConfig(cfg)


	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", api.Ping)
	r.POST("/transactions", api.CheckTransactions(client))
	r.POST("/send-email", api.SendEmail)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}


func main() {
	lambda.Start(Handler)
}

package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/danjavia/stori_csv/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var ginLambda *ginadapter.GinLambda


func init() {
	log.Printf("Gin cold start")
	r := gin.Default()

	r.GET("/ping", api.Ping)
	r.POST("/transactions", api.CheckTransactions)
	r.POST("/send-email", api.SendEmail)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}


func main() {

	lambda.Start(Handler)

	router := gin.Default()

	router.Use(cors.Default())

	router.Run(":8080")
}

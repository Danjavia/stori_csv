package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/danjavia/stori_csv/cmd/infraestructure/models"
)

// Helper function to handle attribute value types consistently
func getValueAsString(v types.AttributeValue) string {
  switch v := v.(type) {
  case *types.AttributeValueMemberS:
      return v.Value
  case *types.AttributeValueMemberN:
      return v.Value // Assuming numbers are stored as strings in DynamoDB
  default:
      // Log or handle unexpected attribute value types
      return "" // Return an empty string as a fallback
  }
}


func CreateSummary(ctx context.Context, client *dynamodb.Client, summary *models.Summary) error {
  // Create the request
  input := &dynamodb.PutItemInput{
    TableName: aws.String("summary"),
    Item: map[string]types.AttributeValue{
      "id": &types.AttributeValueMemberS{Value: summary.ID},
      "userId": &types.AttributeValueMemberS{Value: summary.UserId},
      "email": &types.AttributeValueMemberS{Value: summary.UserEmail},
      "summary": &types.AttributeValueMemberS{Value: summary.Summary},
      "artifactUrl": &types.AttributeValueMemberS{Value: summary.ArtifactUrl},
    },
  }

  // Send the request
  _, err := client.PutItem(ctx, input)
  return err
}

func GetSummaries(ctx context.Context, client *dynamodb.Client, userId string) ([]*models.Summary, error) {
  // Create the query input
  input := &dynamodb.QueryInput{
      TableName: aws.String("summary"),
      KeyConditionExpression: aws.String("userId = :userId"),
      ExpressionAttributeValues: map[string]types.AttributeValue{
          ":userId": &types.AttributeValueMemberS{Value: userId},
      },
  }

  // Send the query
  result, err := client.Query(ctx, input)
  if err != nil {
      return nil, err
  }

  // Parse the response and create summary objects
  summaries := []*models.Summary{}

  for _, item := range result.Items {
      summary := &models.Summary{}
      for k, v := range item {
          switch k {
          case "id":
              summary.ID = getValueAsString(v) // Handle both string and number types
          case "userId":
              summary.UserId = getValueAsString(v)
          case "email":
              summary.UserEmail = getValueAsString(v)
          case "summary":
              summary.Summary = getValueAsString(v)
          case "artifactUrl":
              summary.ArtifactUrl = getValueAsString(v)
          default:
              // Log or handle unexpected attributes
          }
      }
      summaries = append(summaries, summary)
  }

  return summaries, nil
}
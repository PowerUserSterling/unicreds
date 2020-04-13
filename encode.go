package unicreds

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// Encode return the value encoded as a map of dynamo attributes.
func Encode(rawVal interface{}) (map[string]dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(rawVal)
}

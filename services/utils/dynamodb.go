package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Hostname string `json:"hostname"`
	ID       string `json:"id"`
	IP       string `json:"ip"`
	OS       string `json:"os"`
	Port     string `json:"port"`
}

func GetAgents(dbname string) string {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Got error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Create DynamoDB client
	svcDynamo := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(dbname),
	}

	// Make the DynamoDB Query API call
	result, err := svcDynamo.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	// fmt.Print(result)

	var listOfAgents []string

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// fmt.Println(item.Hostname)
		// fmt.Println(item.IP)
		ipPort := item.IP + ":" + item.Port
		listOfAgents = append(listOfAgents, ipPort)

	}
	// fmt.Print(strings.Join(listOfAgents, ","))
	listOfAgentsJoined := strings.Join(listOfAgents, ",")
	return listOfAgentsJoined
}
package utils

import (
	"fmt"
	"os"

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

func DBGetAgents(dbname string, count int) map[string]string {
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

	// var agents [][]string
	agents := make(map[string]string, 0)

	for i, res := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(res, &item)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// fmt.Println(item.Hostname)
		// fmt.Println(item.IP)
		// hostname := item.Hostname
		ipPort := item.IP + ":" + item.Port
		// agents = append(agents, ipPort)
		agents[item.Hostname] = ipPort
		// To get only one entry
		// break
		if count != -1 {
			if i == (count - 1) {
				break
			}
		}
	}
	// fmt.Print(strings.Join(listOfAgents, ","))
	// listOfAgentsJoined := strings.Join(listOfAgents, ",")
	return agents
}

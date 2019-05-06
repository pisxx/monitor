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
	PublicIP string `json:"public_ip"`
	OS       string `json:"os"`
	Port     string `json:"port"`
	Env      string `json:"env"`
}

type agentDetailsP map[string]string
type agentDetailsMap map[string]*agentDetailsP

func DBGetAgents(dbname string, count int) []string {
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
	var agents []string

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
		// ipPort := item.IP + ":" + item.Port
		// agents = append(agents, ipPort)
		agents = append(agents, item.Hostname)
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

func DBGetAgentsDetailed(dbname string, count int) agentDetailsMap {
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
	// var agents []string

	// agentsList :=
	agentsDetailsMap := make(agentDetailsMap, 0)
	for i, res := range result.Items {
		agentDetails := make(agentDetailsP, 0)

		item := Item{}

		err = dynamodbattribute.UnmarshalMap(res, &item)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		agentDetails["hostname"] = item.Hostname
		agentDetails["port"] = item.Port
		agentDetails["ip"] = item.IP
		agentDetails["public_ip"] = item.PublicIP
		agentDetails["env"] = item.Env
		agentsDetailsMap[item.Hostname] = &agentDetails
		// log.Print(agentDetails)
		// log.Print("-------", agentsMap)
		// fmt.Println(item.Hostname)
		// fmt.Println(item.IP)
		// hostname := item.Hostname
		// ipPort := item.IP + ":" + item.Port
		// agents = append(agents, ipPort)

		// agents = append(agents, item.Hostname)
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
	return agentsDetailsMap
}

package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SendAgentsList - sends list of agents to provided queue on SQS
func SendAgentsList(agentsList string, queue string) string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svcSQS := sqs.New(sess)

	// URL to our queue
	qURL := queue

	resultSQS, err := svcSQS.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"ListOfAgents": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(agentsList),
				// StringValue: aws.String("The Whistler"),
			},
		},
		// MessageGroupId: aws.String("100"),
		MessageBody: aws.String("Poll from Agents"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		// fmt.Println("Error", err)
		return err.Error()
	}

	return *resultSQS.MessageId
}

// GetAgentsList - Get list of agents to poll metrics from
func GetAgentsList(queue string) []map[string]string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	qURL := queue

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		panic(err.Error())
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		// panic(err.Error())
		os.Exit(1)
	}
	listOfAgents := *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue
	// fmt.Printf("%v, %T", *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue, result.Messages[0])
	// fmt.Print(strings.Split(listOfAgents, ","))

	// Poll from agents
	// var metricsSlice []utils.MetricsStruct
	fmt.Println()
	// type M map[string]string
	metricsMapList := make([]map[string]string, 3)
	// var metricsMap M
	metricsMap := make(map[string]string)
	for _, i := range strings.Split(listOfAgents, ",") {
		var metrics MetricsStruct
		fmt.Printf("Polling metrics from %s\n", i)
		resp, err := http.Get("http://" + i)
		if err != nil {
			log.Fatal("Unable to poll metrics")
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&metrics)
		if err != nil {
			log.Fatal("issue")
		}
		for _, v := range metrics.Metrics {
			// fmt.Printf("%s: %s\t", v.Name, v.Value)
			metricsMap[v.Name] = v.Value
			metricsMapList = append(metricsMapList, metricsMap)
		}
		// fmt.Println()
		// fmt.Println()
	}
	// fmt.Println(message)
	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		panic(err.Error())
	}

	fmt.Println("Metrics polled, deleting poll request from SQS")
	return metricsMapList
}

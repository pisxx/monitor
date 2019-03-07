package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pisxx/monitor/services/utils"
)

const (
	chooserQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
	// metricsQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/metrics-queue"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	qURL := chooserQURL

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
		return
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return
	}
	listOfAgents := *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue
	// fmt.Printf("%v, %T", *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue, result.Messages[0])
	// fmt.Print(strings.Split(listOfAgents, ","))

	// Poll from agents
	// var metricsSlice []utils.MetricsStruct
	fmt.Println()
	for _, i := range strings.Split(listOfAgents, ",") {
		var metrics utils.MetricsStruct
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
			fmt.Printf("%s: %s\t", v.Name, v.Value)
		}
		fmt.Println()
		fmt.Println()
	}
}

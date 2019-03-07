package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
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
	// var message sqs.Message
	// sqs.
	// json.Unmarshal([]byte(result.Messages[0]), &message)
	// err = json.NewDecoder(result.Messages[0]).Decode(&message)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	// return nil, err
	// }
	listOfAgents := *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue
	// fmt.Printf("%v, %T", *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue, result.Messages[0])
	fmt.Print(strings.Split(listOfAgents, ","))

	// fmt.Println(message)
}

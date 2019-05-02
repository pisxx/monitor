package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SendAgentsList - sends list of agents to provided queue on SQS
func SQSSendAgentsList(agents map[string]string, queue string) string {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))
	// sess.Config['Region'] = "us-east-1"
	svcSQS := sqs.New(sess)

	// URL to our queue
	qURL := queue

	agentsHostnameIpPort := make([]string, 0)
	for k, v := range agents {
		hostnameIpPort := k + ";" + v
		fmt.Println(hostnameIpPort)
		agentsHostnameIpPort = append(agentsHostnameIpPort, hostnameIpPort)

	}
	// fmt.Println(agentsHostnameIpPort)
	agentsList := strings.Join(agentsHostnameIpPort, ",")
	// fmt.Println(agentList)
	// agentsList := strings.Join(agents, ",")
	resultSQS, err := svcSQS.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
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
func SQSGetAgentsList(queue string) ([]string, *string) {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

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
		WaitTimeSeconds:     aws.Int64(3),
	})

	if err != nil {
		fmt.Println("Error", err)
		panic(err.Error())
	}

	if len(result.Messages) == 0 {
		// fmt.Println("Received no messages")
		log.Println("Received no messages")
		// var message *string
		message := "Received no messages"
		return nil, &message
		// panic(err.Error())
		// os.Exit(1)
	}
	if len(result.Messages) != 0 {
		listOfAgents := *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue
		// fmt.Printf("%v, %T", *result.Messages[0].MessageAttributes["ListOfAgents"].StringValue, result.Messages[0])
		// fmt.Print(strings.Split(listOfAgents, ","))
		listOfAgentsSplitted := strings.Split(listOfAgents, ",")
		return listOfAgentsSplitted, result.Messages[0].ReceiptHandle
	}
	return nil, nil
}

func DeleteMessage(messageID *string, qURL string) error {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	))

	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: messageID,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return err
	}
	// fmt.Println("Deleted message from queue")
	return nil
}

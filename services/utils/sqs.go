package aws

import (
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
		MessageBody: aws.String("Agent to poll metrics from"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		// fmt.Println("Error", err)
		return err.Error()
	}

	return *resultSQS.MessageId
}

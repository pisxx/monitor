import boto3
import os
import uuid

not_enough_args_error = "Not enough args"

def lambda_handler(event, context):
    
    # recordId = str(uuid.uuid4())
    # voice = event["voice"]
    # text = event["text"]
    if len(event) != 3:
        return "Bad Request: Not enough args"
        # raise ValueError("Bad Request: {}".format)
    agentID = str(uuid.uuid4())
    Hostname = event["Hostname"]
    OS = event["OS"]
    IP = event["IP"]

    print('Registering new Agent, with ID: ' + agentID)
    # print('Input Text: ' + text)
    # print('Selected voice: ' + voice)
    
    #Creating new record in DynamoDB table
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table("agents")
    table.put_item(
        Item={
            'agentID' : agentID,
            'Hostaname': Hostname,
            'OS': OS,
            'IP': IP
        }
    )
    
    # #Sending notification about new post to SNS
    # client = boto3.client('sns')
    # client.publish(
    #     TopicArn = os.environ['SNS_TOPIC'],
    #     Message = recordId
    # )
    
    return "Agent {} registered".format(agentID)

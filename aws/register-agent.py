import boto3
import os
import uuid
from boto3.dynamodb.conditions import Key, Attr

not_enough_args_error = "Not enough args"

def lambda_handler(event, context):
    
    # recordId = str(uuid.uuid4())
    # voice = event["voice"]
    # text = event["text"]

    if len(event) != 3:
        return "Bad Request: Not enough args"
        # raise ValueError("Bad Request: {}".format)
    hostname = event["hostname"]
    id = check_hostname(hostname)
    if id:
        return "Agent already registered with ID: {}".format(id["Items"][0]["id"])
    id = str(uuid.uuid4())
    os = event["os"]
    ip = event["ip"]

    print('Registering new Agent, with ID: ' + id)
    # print('Input Text: ' + text)
    # print('Selected voice: ' + voice)
    
    #Creating new record in DynamoDB table
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table("agents_hostname")
    table.put_item(
        Item={
            'id' : id,
            'hostname': hostname,
            'os': os,
            'ip': ip
        }
    )
    
    # #Sending notification about new post to SNS
    # client = boto3.client('sns')
    # client.publish(
    #     TopicArn = os.environ['SNS_TOPIC'],
    #     Message = recordId
    # )
    
    return "Agent {} registered".format(id)


def check_hostname(hostname):
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table("agents_hostname")
    agentId = table.query(
        KeyConditionExpression=Key('hostname').eq(hostname)
    )
    return agentId
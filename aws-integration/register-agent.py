import boto3
import os
import uuid
import json
from boto3.dynamodb.conditions import Key, Attr

response = {}

class RegisterException(Exception):
    pass


def lambda_handler(event, context):

    if len(event) != 6:
        raise RegisterException("Bad Request: Not enough args")
    hostname = event["hostname"]
    query_db = check_hostname(hostname)
    if query_db["Items"]:
        response["message"] = "Agent already registered with ID: {}".format(query_db["Items"][0]["id"])
        # return "Agent already registered with ID: {}".format(query_db["Items"][0]["id"])
        return response
        # return json.dumps(response)
    id = str(uuid.uuid4())
    os = event["os"]
    ip = event["ip"]
    public_ip = event["public_ip"]
    port = event["port"]
    env = event["env"]
    if ip == "127.0.0.1":
        raise RegisterException("Bad Request: 127.0.0.1 is not allowed")
    print('Registering new Agent, with ID: ' + id)
    
    #Creating new record in DynamoDB table
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table("agents_hostname")
    table.put_item(
        Item={
            # 'id' : id,
            'hostname': hostname,
            'os': os,
            'ip': ip,
            'public_ip': public_ip,
            'port': port,
            'env' : env
        }
    )
    
    # #Sending notification about new post to SNS
    # client = boto3.client('sns')
    # client.publish(
    #     TopicArn = os.environ['SNS_TOPIC'],
    #     Message = recordId
    # )
    response["message"] = "Agent {} registered".format(id)
    # response["status_code"] = 200
    # json.dumps(response)
    return response
    # return "Agent {} registered".format(id)


def check_hostname(hostname):
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table("agents_hostname")
    hostname_db = table.query(
        KeyConditionExpression=Key('hostname').eq(hostname)
    )
    return hostname_db
#!/bin/bash
zip register.zip register-agent.py
aws lambda update-function-code --function-name monitor-register-agent --zip-file fileb://register.zip

## testing issue

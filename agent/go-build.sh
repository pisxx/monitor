#!/bin/bash

export GOOS=linux
go build -o agent-linux agent.go
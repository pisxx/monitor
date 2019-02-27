package register

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

const registerURL = "https://1iw0vyiwzc.execute-api.us-east-1.amazonaws.com/dev/register"

// Agent - register agent in DynamoDB in AWS
func Agent() (string, error) {

	message := make(map[string]string)

	agentID, err := os.Hostname()
	if err != nil {
		return "Failed", err
	}
	message["agentID"] = agentID
	message["agentOS"] = runtime.GOOS

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return "Failed", err
	}
	resp, err := http.Post(registerURL, "application/text", bytes.NewBuffer(messageBytes))
	defer resp.Body.Close()
	if err != nil {
		return "Failed", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Failed", err
	}
	bodyString := string(body)
	return string(bodyString), nil
	// resp.Body.Close()
	// json.NewDecoder(resp.Body).Decode(&result)

}

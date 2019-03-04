package register

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type registerError struct {
	Message string
}

const registerURL = "https://1iw0vyiwzc.execute-api.us-east-1.amazonaws.com/dev/register"

// RegisterAgent - register agent in DynamoDB in AWS
func RegisterAgent(agentIP string) (string, error) {

	message := make(map[string]string)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	// message["agentID"] = agentID
	message["os"] = runtime.GOOS
	message["ip"] = agentIP
	message["hostname"] = hostname

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return "Failed", err
	}
	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return "Failed", err
	}
	defer resp.Body.Close()

	if err != nil {
		return "error", err
	}

	// return string(body), nil
	if resp.StatusCode == 400 {
		var respErr registerError
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return err.Error(), nil
			// panic(err.Error())
		}
		errMsg := strings.Split(respErr.Message, ":")
		errMsgS := strings.Trim(errMsg[1], " ")
		// fmt.Printf(errMsgS)
		switch errMsgS {
		// switch os := runtime.GOOS; os {
		case "Not enough args":
			log.Print("Unable to register Agent, see response below")
			// log.Print("See response below")
			log.Printf("Response code: %d", resp.StatusCode)
			log.Fatalf("Response body: \"%s\"", respErr.Message)
			return respErr.Message, err
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "issue", err
	}
	return strings.Trim(string(body), "\""), nil
	// resp.Body.Close()
	// json.NewDecoder(resp.Body).Decode(&result)

}

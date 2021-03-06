package register

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type RegisterResponse struct {
	Message string
	// StatusCode int
}

const registerURL = "https://1iw0vyiwzc.execute-api.us-east-1.amazonaws.com/dev/register"

// RegisterAgent - register agent in DynamoDB in AWS
func RegisterAgent(agentIP string, agentPort string, agentEnv string) (string, error) {

	message := make(map[string]string)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	// message["agentID"] = agentID
	publicIP := getAWSPublicIP()
	if publicIP != "NOT AWS" {
		message["public_ip"] = publicIP
	} else {
		message["public_ip"] = "NA"
	}
	message["os"] = runtime.GOOS
	message["ip"] = agentIP
	message["port"] = agentPort
	message["hostname"] = hostname
	message["env"] = agentEnv

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return "Failed", err
	}
	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return err.Error(), err
	}
	defer resp.Body.Close()

	// return string(body), nil
	if resp.StatusCode == 400 {
		var respErr RegisterResponse
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return err.Error(), nil
			// panic(err.Error())
		}
		errMsg := strings.Split(respErr.Message, ":")
		errMsgS := strings.Trim(errMsg[1], " ")
		// fmt.Println(errMsg)
		switch errMsgS {
		// switch os := runtime.GOOS; os {
		case "Not enough args":
			// log.Print("Unable to register Agent, see response below")
			// log.Print("See response below")
			// log.Printf("Response code: %d", resp.StatusCode)
			// log.Fatalf("Response body: \"%s\"", respErr.Message)
			return respErr.Message, err
		case "127.0.0.1 is not allowed":
			// log.Print("Unable to register agent, please use accesible ip")
			// log.Printf("You may still check metrics at below ip:port")
			// return "Unable to register agent, please use accesible ip", err
			log.Printf("Unable to register")
			return respErr.Message, err
		}
	}

	var registerResp RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		return err.Error(), err
	}
	return registerResp.Message, nil
	// return json.Unmarshal(body, RegisterResponse) nil
	// resp.Body.Close()
	// json.NewDecoder(resp.Body).Decode(&result)

}

package register

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
			return respErr.Message, err
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), err
	}
	var registerResp RegisterResponse
	if err := json.NewDecoder(body).Decode(&registerResp); err != nil {
		return err.Error(), err
	}
	return registerResp.Message, nil
	// return json.Unmarshal(body, RegisterResponse) nil
	// resp.Body.Close()
	// json.NewDecoder(resp.Body).Decode(&result)

}

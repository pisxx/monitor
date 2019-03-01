package register

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

const kvPath = "monitor/agents"

// ConsulRegister - register Agent in Consul KV store
func ConsulRegister(agentIP string) {
	// Get a new client
	agentID := generateID(10)
	var config api.Config
	config.Address = "3.88.101.138:8500"
	// client, err := api.NewClient(api.DefaultConfig())
	client, err := api.NewClient(&config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Check if ID already exits

	// PUT a new KV pair
	// p := &api.KVPair{Key: "monitor/agents", Value: []byte("1000")}
	agentKey := kvPath + "/" + agentID
	p := &api.KVPair{Key: agentKey + "/ip", Value: []byte(agentIP)}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get(agentKey+"/ip", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

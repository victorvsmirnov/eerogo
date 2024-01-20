package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/victorvsmirnov/eerogo"
)

func main() {
	config := flag.String("c", ".env", "Configuration file")
	flag.Parse()
	configuration := Config{}
	err := eerogo.LoadViperConfiguration(*config, &configuration)
	if err != nil {
		panic(err)
	}
	eeroclient := eerogo.NewEeroClient(configuration.Eero)

	err = eeroclient.LoadCookie()
	if err != nil {
		err = eeroclient.Login()
		if err != nil {
			panic(err)
		}
		var verificatinoKey string
		fmt.Printf("Enter verification key: ")
		n, err := fmt.Scan(&verificatinoKey)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			panic(fmt.Errorf("null input"))
		}
		eeroclient.VerifyKey(verificatinoKey)
		if err != nil {
			panic(err)
		}
		err = eeroclient.SaveCookie()
		if err != nil {
			panic(err)
		}
	} else {
		err = eeroclient.LoginRefresh()
		if err != nil {
			panic(err)
		}
	}
	account, err := eeroclient.Account()
	if err != nil {
		panic(err)
	}
	fmt.Printf("account: %v\n", account)
	fmt.Printf("eeroclient: %v\n", eeroclient)

	for networkId := range eeroclient.Cache.Networks {
		network, err := eeroclient.Network(networkId)
		if err != nil {
			panic(err)
		}
		j, _ := json.Marshal(network)
		fmt.Printf("network: %s\n\n", string(j))

		clients, err := eeroclient.NetworkClients(networkId)
		if err != nil {
			panic(err)
		}
		j, _ = json.Marshal(clients)
		fmt.Printf("clients: %s\n", string(j))
	}

}

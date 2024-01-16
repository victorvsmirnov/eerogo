package main

import (
	"flag"
	"fmt"

	"github.com/victorvsmirnov/eerogo"
)

func main() {

	// sessionKey := flag.String("sessionKey", "", "Eero session key")
	// loginID := flag.String("loginID", "", "Eero loginId")
	config := flag.String("c", ".env", "Configuration file")
	// verificationKey := flag.String("verificationKey", "", "Eero verification key")
	// networkID := flag.String("networkID", "", "Network ID to monitor")
	// command := flag.String("command", "monitor", "Actions to perform")
	flag.Parse()
	configuration := Config{}
	err := eerogo.LoadViperConfiguration(*config, &configuration)
	if err != nil {
		panic(err)
	}
	eeroclient := eerogo.NewEeroClient(configuration.Eero)
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

}

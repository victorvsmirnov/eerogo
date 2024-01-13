package main

import (
	"flag"
	"fmt"
)

func main() {

	sessionKey := flag.String("sessionKey", "", "Eero session key")
	loginID := flag.String("loginID", "", "Eero loginId")
	verificationKey := flag.String("verificationKey", "", "Eero verification key")
	networkID := flag.String("networkID", "", "Network ID to monitor")
	command := flag.String("command", "monitor", "Actions to perform")
	flag.Parse()

	if *verificationKey != "" && *sessionKey != "" {
		verifyKey(verificationKey, sessionKey)
		fmt.Printf("Next monitor netowork with networkID (/2.2/networks/[ID]):\n")
		fmt.Printf("\t./eeroMonitor -sessionKey=\"%s\" --networkID=\n", *sessionKey)

	} else if *loginID != "" {
		sessionKey := login(loginID)
		//fmt.Printf("sessionKey=%s\n", sessionKey)
		fmt.Printf("Next verify session with verification code: \n")
		fmt.Printf("\t./eeroMonitor -sessionKey=\"%s\" -verificationKey=\n", sessionKey)

	} else if *sessionKey != "" && *networkID != "" {
		if *command == "" || *command == "monitor" {
			monitor(sessionKey, networkID)
		} else {
			fmt.Printf("Unknown command: %s\n", *command)

		}

	} else {
		fmt.Printf("Unknow set of arguments...\n")
		fmt.Printf("\t./eeroMonitor -loginID=[YOUR_LOGIN_ID]\n")
		return
	}
}

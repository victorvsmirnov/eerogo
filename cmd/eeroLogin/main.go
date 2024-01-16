package main

import (
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
	err = eeroclient.LoginRefresh()
	if err != nil {
		panic(err)
	}

}

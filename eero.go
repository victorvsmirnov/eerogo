package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	// "github.com/pkg/errors"
)

// LoginRequest is use to login to Eero. Set the account email address
type LoginRequest struct {
	Login string `json:"login"`
}

// LoginResponse response to the login request
// user_token will be used as `Cookie: s={{user_token}}`
// Code looks similar to the HTTP response code
type LoginResponse struct {
	Meta struct {
		Code       int       `json:"code"`
		ServerTime time.Time `json:"server_time"`
	} `json:"meta"`
	Data struct {
		UserToken string `json:"user_token"`
	} `json:"data"`
}

// LoginVerifyRequest sends the code from email
type LoginVerifyRequest struct {
	Code string `json:"code"`
}

type UserContact struct {
	Value    string `json:"value"`
	Verified bool   `json:"verified"`
}

type LoginVerifyData struct {
	Name     string      `json:"name"`
	Phone    UserContact `json:"phone"`
	Email    UserContact `json:"email"`
	LogID    string      `json:"log_id"`
	Networks struct {
		Count int `json:"count"`
		Data  []struct {
			URL     string    `json:"url"`
			Name    string    `json:"name"`
			Created time.Time `json:"created"`
		} `json:"data"`
	} `json:"networks"`
	Role          string `json:"role"`
	CanTransfer   bool   `json:"can_transfer"`
	IsProOwner    bool   `json:"is_pro_owner"`
	PremiumStatus string `json:"premium_status"`
	PushSettings  struct {
		NetworkOffline bool `json:"networkOffline"`
		NodeOffline    bool `json:"nodeOffline"`
	} `json:"push_settings"`
	TrustCertificatesEtag string `json:"trust_certificates_etag"`
}

// LoginVerifyResponse Returns details about your network
type LoginVerifyResponse struct {
	Meta struct {
		Code       int       `json:"code"`
		ServerTime time.Time `json:"server_time"`
	} `json:"meta"`
	Data LoginVerifyData `json:"data"`
}

// LogoutResponse session logout
type LogoutResponse struct {
	Meta struct {
		Code       int       `json:"code"`
		ServerTime time.Time `json:"server_time"`
	} `json:"meta"`
}

type Meta struct {
	Code       int       `json:"code"`
	ServerTime time.Time `json:"server_time"`
}

type NetworkConnectivity struct {
	RxBitrate string  `json:"rx_bitrate"`
	Signal    string  `json:"signal"`
	SignalAvg string  `json:"signal_avg"`
	Score     float64 `json:"score"`
	ScoreBars int     `json:"score_bars"`
}

type NetworkProfile struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Paused bool   `json:"paused"`
}

type Source struct {
	Location string `json:"location"`
}

type NetworkInterface struct {
	Frequency     string `json:"frequency"`
	FrequencyUnit string `json:"frequency_unit"`
}

type Usage struct {
	DownMbps float64 `json:"down_mbps"`
	UpMbps   float64 `json:"up_mbps"`
}
type NetworkData struct {
	URL            string              `json:"url"`
	Mac            string              `json:"mac"`
	Eui64          string              `json:"eui64"`
	Manufacturer   string              `json:"manufacturer"`
	IP             string              `json:"ip"`
	Ips            []string            `json:"ips"`
	Nickname       interface{}         `json:"nickname"`
	Hostname       string              `json:"hostname"`
	Connected      bool                `json:"connected"`
	Wireless       bool                `json:"wireless"`
	ConnectionType string              `json:"connection_type"`
	Source         Source              `json:"source"`
	LastActive     time.Time           `json:"last_active"`
	FirstActive    time.Time           `json:"first_active"`
	Connectivity   NetworkConnectivity `json:"connectivity"`
	Interface      NetworkInterface    `json:"interface"`
	Usage          Usage               `json:"usage"`
	Profile        NetworkProfile      `json:"profile"`
	DeviceType     string              `json:"device_type"`
}

// NetworkDeviceResponse Details about devices on network
type NetworkDeviceResponse struct {
	Meta Meta          `json:"meta"`
	Data []NetworkData `json:"data"`
}

type EeroConfiguration struct {
	Login          string
	CookieFileName string
	URL            string // https://api-user.e2ro.com / https://api-user.e2ro.com/2.2/
}

type EeroClient struct {
	httpClient *http.Client
	config     EeroConfiguration
	userToken  string
}

func (e *EeroClient) Login() (err error) {
	fmt.Printf("Login: %s\n", e.config.Login)
	// url := "https://api-user.e2ro.com/2.2/login?"

	loginRequest := LoginRequest{Login: e.config.Login}
	var loginResponse LoginResponse

	err = e.do("POST", "login", &loginRequest, &loginResponse)
	//r, err := http.Post("https://api-user.e2ro.com/2.2/login?", "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}
	// if loginResponse.Meta.Code?
	e.userToken = loginResponse.Data.UserToken
	return nil
}

func (e *EeroClient) verifyKey(verificationKey string) string {
	// fmt.Printf("Verify: %s, %s\n", *verificationKey, *sessionKey)

	verifyRequest := LoginVerifyRequest{Code: verificationKey}
	var verifyResponse LoginVerifyResponse
	// url := "https://api-user.e2ro.com/2.2/login/verify?"
	err := e.do("POST", "login/verify", &verifyRequest, &verifyResponse)

	if err != nil {
		panic(err)
	}

	//fmt.Println(verifyResponse)

	networks := verifyResponse.Data.Networks.Data
	for _, network := range networks {
		fmt.Printf("%s - %s\n", network.Name, network.URL)
	}
	return ""
}

// func monitor(sessionKey *string, networkID *string) {
// 	fmt.Printf("Monitoring Network: %s\n", *networkID)
// 	url := fmt.Sprintf("https://api-user.e2ro.com%s/devices?thread=true", *networkID)
// 	fmt.Printf("URL: %s\n", url)

// 	// TODO: Query /2.2/networks/[NETWORK_ID]/burst_reporters?
// 	// to schedule next time to query usage

// 	for {
// 		var networkDeviceResponse NetworkDeviceResponse
// 		err := doRequest(url, sessionKey, nil, &networkDeviceResponse)

// 		if err != nil {
// 			panic(err)
// 		}

// 		networks := networkDeviceResponse.Data
// 		foundResult := false
// 		for _, device := range networks {
// 			up := device.Usage.UpMbps
// 			down := device.Usage.DownMbps
// 			if up > 0 || down > 0 {
// 				foundResult = true
// 				fmt.Printf("%s - %s (%f Mbps, %f Mbps)\n", device.Hostname, device.DeviceType, device.Usage.DownMbps, device.Usage.UpMbps)
// 			}
// 		}

// 		if foundResult {
// 			fmt.Printf("\n\n\n\n")
// 		}
// 		time.Sleep(60 * time.Second)
// 	}
// }

func (e *EeroClient) do(method string, url string, reqObj interface{}, respObj interface{}) error {

	b := new(bytes.Buffer)
	if reqObj != nil {
		method = "POST"
		json.NewEncoder(b).Encode(reqObj)
	} else {
		b = nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	if e.userToken == "" {
		// sessionString := fmt.Sprintf("s=%s", *token)
		//fmt.Printf("Session Key: %s\n", sessionString)
		req.Header.Add("Cookie", fmt.Sprintf("s=%s", e.userToken))
	}
	if req != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("Request failed: (%d) - %s\nURL: %s %s\nRequest: %s", r.StatusCode, r.Status, method, url, reqObj)
	}
	if r.Body == nil && respObj == nil {
		return nil
	}
	err = json.NewDecoder(r.Body).Decode(respObj)
	if err != nil {
		return err
	}

	return nil
}

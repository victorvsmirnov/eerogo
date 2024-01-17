package eerogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type EeroURL string

type EeroCache struct {
	Networks map[EeroURL]AccountNetworkData
}

func NewEeroClient(c EeroConfiguration) *EeroClient {
	cache := EeroCache{
		Networks: map[EeroURL]AccountNetworkData{},
	}
	return &EeroClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		config:    c,
		userToken: "",
		Cache:     &cache,
	}
}

type EeroConfiguration struct {
	Login          string
	CookieFileName string
	URL            string // https://api-user.e2ro.com / https://api-user.e2ro.com/2.2/
}

type EeroClient struct {
	httpClient      *http.Client
	config          EeroConfiguration
	userToken       string
	LoginVerifyData LoginVerifyData
	Cache           *EeroCache
}

func (e *EeroClient) Login() (err error) {
	loginRequest := LoginRequest{Login: e.config.Login}
	var loginResponse LoginResponse

	err = e.do("POST", "2.2/login", &loginRequest, &loginResponse)
	if err != nil {
		return err
	}
	// if loginResponse.Meta.Code?
	e.userToken = loginResponse.Data.UserToken
	return nil
}

func (e *EeroClient) SaveCookie() error {
	if e.userToken == "" {
		return fmt.Errorf("client not authenticated")
	}
	return os.WriteFile(e.config.CookieFileName, []byte(e.userToken), 0600)
}
func (e *EeroClient) LoadCookie() error {
	b, err := os.ReadFile(e.config.CookieFileName)
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return fmt.Errorf("cookie file is empty")
	}
	e.userToken = string(b)
	return nil
}

func (e *EeroClient) LoginRefresh() error {
	var response LoginResponse
	err := e.do("POST", "2.2/login/refresh", nil, &response)
	if err != nil {
		return err
	}
	e.userToken = response.Data.UserToken
	return nil
}

func (e *EeroClient) Account() (*AccountResponse, error) {
	var response AccountResponse
	err := e.do("GET", "2.2/account", nil, &response)
	if err != nil {
		return nil, err
	}
	for _, network := range response.Data.Networks.Data {
		e.Cache.Networks[network.URL] = network
	}
	return &response, nil
}

func (e *EeroClient) Network(url EeroURL) (any, error) {
	var response any
	err := e.do("GET", url, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (e *EeroClient) VerifyKey(verificationKey string) error {
	verifyRequest := LoginVerifyRequest{Code: verificationKey}
	var verifyResponse LoginVerifyResponse
	err := e.do("POST", "login/verify", &verifyRequest, &verifyResponse)

	if err != nil {
		return err
	}
	e.LoginVerifyData = verifyResponse.Data
	//fmt.Println(verifyResponse)
	return nil
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

func (e *EeroClient) do(method string, url EeroURL, reqObj interface{}, respObj interface{}) error {

	b := new(bytes.Buffer)
	if reqObj != nil {
		json.NewEncoder(b).Encode(reqObj)
	}
	// else {
	// 	b = nil
	// }

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", e.config.URL, url), b)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	if e.userToken != "" {
		req.Header.Add("Cookie", fmt.Sprintf("s=%s", e.userToken))
	}
	if req != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	r, err := e.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 && r.StatusCode != 401 {
		return fmt.Errorf("request failed: (%d) - %s\nURL: %s %s\nRequest: %s", r.StatusCode, r.Status, method, url, reqObj)
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

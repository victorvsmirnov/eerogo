package eerogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"
)

const APIVersion = "2.2"

type ResourceID string

type EeroCache struct {
	Networks map[ResourceID]string
}

func (ec *EeroCache) Clear() {
	ec.Networks = make(map[ResourceID]string)
}

func NewEeroClient(c EeroConfiguration) *EeroClient {
	cache := EeroCache{
		Networks: map[ResourceID]string{},
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

type EeroClient struct {
	httpClient      *http.Client
	config          EeroConfiguration
	userToken       string
	LoginVerifyData LoginVerifyData
	Cache           *EeroCache
}

func (eeroclient *EeroClient) LoginSequence(verificationInputFunc func() (string, error)) (err error) {
	err = eeroclient.LoadCookie()
	if err != nil {
		err = eeroclient.Login()
		if err != nil {
			return err
		}
		verificationKey, err := verificationInputFunc()
		if err != nil {
			return err
		}

		err = eeroclient.VerifyKey(verificationKey)
		if err != nil {
			return err
		}
		err = eeroclient.SaveCookie()
		if err != nil {
			return err
		}
	} else {
		err = eeroclient.LoginRefresh()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EeroClient) Login() (err error) {
	loginRequest := LoginRequest{Login: e.config.Login}
	var loginResponse LoginResponse

	err = e.do("POST", fmt.Sprintf("%s/login", APIVersion), &loginRequest, &loginResponse)
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
	err := e.do("POST", fmt.Sprintf("%s/login/refresh", APIVersion), nil, &response)
	if err != nil {
		return err
	}
	e.userToken = response.Data.UserToken
	return nil
}

var urlReMap map[string]*regexp.Regexp = map[string]*regexp.Regexp{
	"networks": regexp.MustCompile(`\/[\d\.]+\/networks\/(\d+)`),
	"eeros":    regexp.MustCompile(`\/[\d\.]+\/eeros\/(\d+)`),
}

func extract_resource_id(url string) (ResourceID, error) {
	for _, r := range urlReMap {
		m := r.FindStringSubmatch(string(url))
		if m != nil {
			return ResourceID(m[1]), nil
		}
	}
	return "", fmt.Errorf("unbale to extract id from %s", url)
}

func (e *EeroClient) Account() (*AccountResponse, error) {
	var response AccountResponse
	err := e.do("GET", fmt.Sprintf("%s/account", APIVersion), nil, &response)
	if err != nil {
		return nil, err
	}
	for _, network := range response.Data.Networks.Data {
		id, err := extract_resource_id(network.URL)
		if err != nil {
			return nil, err
		}
		e.Cache.Networks[id] = network.URL
	}
	return &response, nil
}

func (e *EeroClient) Network(id ResourceID) (*NetworkData, error) {
	url, ok := e.Cache.Networks[id]
	if !ok {
		return nil, fmt.Errorf("network id is not fould")
	}
	var response NetworkData
	err := e.do("GET", url, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (e *EeroClient) NetworkClients(id ResourceID) (*NetworkClientsResponse, error) {
	url, ok := e.Cache.Networks[id]
	if !ok {
		return nil, fmt.Errorf("network id is not fould")
	}
	var response NetworkClientsResponse
	err := e.do("GET", fmt.Sprintf("%s/clients", url), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (e *EeroClient) VerifyKey(verificationKey string) error {
	verifyRequest := LoginVerifyRequest{Code: verificationKey}
	var verifyResponse LoginVerifyResponse
	err := e.do("POST", fmt.Sprintf("%s/login/verify", APIVersion), &verifyRequest, &verifyResponse)
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

func (e *EeroClient) do(method string, url string, reqObj interface{}, respObj interface{}) error {

	b := new(bytes.Buffer)
	if reqObj != nil {
		json.NewEncoder(b).Encode(reqObj)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", e.config.URL, url), b)
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

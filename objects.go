package eerogo

import "time"



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

type NetworkBriefData struct {
	URL     string    `json:"url"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type LoginVerifyData struct {
	Name     string      `json:"name"`
	Phone    UserContact `json:"phone"`
	Email    UserContact `json:"email"`
	LogID    string      `json:"log_id"`
	Networks struct {
		Count int                `json:"count"`
		Data  []NetworkBriefData `json:"data"`
	} `json:"networks"`
	Role                  string       `json:"role"`
	CanTransfer           bool         `json:"can_transfer"`
	IsProOwner            bool         `json:"is_pro_owner"`
	PremiumStatus         string       `json:"premium_status"`
	PushSettings          PushSettings `json:"push_settings"`
	TrustCertificatesEtag string       `json:"trust_certificates_etag"`
}

// LoginVerifyResponse Returns details about your network
type LoginVerifyResponse struct {
	Meta Meta            `json:"meta"`
	Data LoginVerifyData `json:"data"`
}

// LogoutResponse session logout
type LogoutResponse struct {
	Meta Meta `json:"meta"`
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

type NetworkUsage struct {
	DownMbps float64 `json:"down_mbps"`
	UpMbps   float64 `json:"up_mbps"`
}
type NetworkDeviceData struct {
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
	Usage          NetworkUsage        `json:"usage"`
	Profile        NetworkProfile      `json:"profile"`
	DeviceType     string              `json:"device_type"`
}

// NetworkDeviceResponse Details about devices on network
type NetworkDeviceResponse struct {
	Meta Meta                `json:"meta"`
	Data []NetworkDeviceData `json:"data"`
}

type AccountResponse struct {
	Data AccountData `json:"data"`
	Meta Meta        `json:"meta"`
}
type Auth struct {
	ProviderID any    `json:"provider_id"`
	ServiceID  any    `json:"service_id"`
	Type       string `json:"type"`
}
type MarketingEmails struct {
	Consented bool `json:"consented"`
}
type Consents struct {
	MarketingEmails MarketingEmails `json:"marketing_emails"`
}
type Email struct {
	Value    string `json:"value"`
	Verified bool   `json:"verified"`
}
type NetworkData struct {
	AccessExpiresOn  any       `json:"access_expires_on"`
	AmazonDirectedID any       `json:"amazon_directed_id"`
	Created          time.Time `json:"created"`
	Name             string    `json:"name"`
	NicknameLabel    any       `json:"nickname_label"`
	URL              string    `json:"url"`
}
type Networks struct {
	Count int           `json:"count"`
	Data  []NetworkData `json:"data"`
}
type Phone struct {
	CountryCode    string `json:"country_code"`
	NationalNumber string `json:"national_number"`
	Value          string `json:"value"`
	Verified       bool   `json:"verified"`
}
type PremiumDetails struct {
	HasPaymentInfo       bool   `json:"has_payment_info"`
	Interval             any    `json:"interval"`
	IsIapCustomer        bool   `json:"is_iap_customer"`
	NextBillingEventDate any    `json:"next_billing_event_date"`
	PaymentMethod        any    `json:"payment_method"`
	Tier                 string `json:"tier"`
	TrialEnds            any    `json:"trial_ends"`
}
type PushSettings struct {
	NetworkOffline bool `json:"networkOffline"`
	NodeOffline    bool `json:"nodeOffline"`
}

type AccountData struct {
	Auth                      Auth           `json:"auth"`
	BusinessDetails           any            `json:"business_details"`
	CanMigrateToAmazonLogin   bool           `json:"can_migrate_to_amazon_login"`
	CanTransfer               bool           `json:"can_transfer"`
	Consents                  Consents       `json:"consents"`
	EeroForBusiness           bool           `json:"eero_for_business"`
	Email                     Email          `json:"email"`
	ImageAssets               any            `json:"image_assets"`
	IsBetaBugReporterEligible bool           `json:"is_beta_bug_reporter_eligible"`
	IsPremiumCapable          bool           `json:"is_premium_capable"`
	LogID                     string         `json:"log_id"`
	MduProgram                bool           `json:"mdu_program"`
	Name                      string         `json:"name"`
	Networks                  Networks       `json:"networks"`
	OrganizationID            any            `json:"organization_id"`
	PaymentFailed             bool           `json:"payment_failed"`
	Phone                     Phone          `json:"phone"`
	PremiumDetails            PremiumDetails `json:"premium_details"`
	PremiumStatus             string         `json:"premium_status"`
	PushSettings              PushSettings   `json:"push_settings"`
	Role                      string         `json:"role"`
	TrustCertificatesEtag     string         `json:"trust_certificates_etag"`
}
type Meta struct {
	Code       int       `json:"code"`
	ServerTime time.Time `json:"server_time"`
}

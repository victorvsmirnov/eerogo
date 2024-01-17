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
	URL     EeroURL   `json:"url"`
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
	URL    Eero URL `json:"url"`
	Name   string  `json:"name"`
	Paused bool    `json:"paused"`
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
	URL            EeroURL             `json:"url"`
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
type AccountNetworkData struct {
	AccessExpiresOn  any       `json:"access_expires_on"`
	AmazonDirectedID any       `json:"amazon_directed_id"`
	Created          time.Time `json:"created"`
	Name             string    `json:"name"`
	NicknameLabel    any       `json:"nickname_label"`
	URL              EeroURL   `json:"url"`
}
type AccountNetworks struct {
	Count int           `json:"count"`
	Data  []AccountNetworkData `json:"data"`
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
	Networks                  AccountNetworks       `json:"networks"`
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


type NetworkData struct {
	Data struct {
		AccessExpiresOn       any  `json:"access_expires_on"`
		AlexaSkill            bool `json:"alexa_skill"`
		AmazonAccountLinked   bool `json:"amazon_account_linked"`
		AmazonDeviceNickname  bool `json:"amazon_device_nickname"`
		AmazonDirectedID      any  `json:"amazon_directed_id"`
		AmazonFullName        any  `json:"amazon_full_name"`
		BackupInternetEnabled bool `json:"backup_internet_enabled"`
		BandSteering          bool `json:"band_steering"`
		Capabilities          struct {
			AcCompat struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion  bool `json:"has_min_mobile_version"`
					HasSupportedHardware bool `json:"has_supported_hardware"`
				} `json:"requirements"`
			} `json:"ac_compat"`
			AccountLinking struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion   bool `json:"has_min_mobile_version"`
					HasRequiredAuthMethod bool `json:"has_required_auth_method"`
					InEnabledGroups       bool `json:"in_enabled_groups"`
					IsFeatureEnabled      bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"account_linking"`
			Acs struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"acs"`
			AdBlock struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAdblockFeatureFlag      bool `json:"has_adblock_feature_flag"`
					HasAdblockGroup            bool `json:"has_adblock_group"`
					HasAdblockMinMobileVersion bool `json:"has_adblock_min_mobile_version"`
					HasAdblockMinVersion       bool `json:"has_adblock_min_version"`
					HasAdblockThrottle         bool `json:"has_adblock_throttle"`
					HasPremium                 bool `json:"has_premium"`
					PremiumCapable             bool `json:"premium_capable"`
				} `json:"requirements"`
			} `json:"ad_block"`
			AdBlockViewable struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAdblockFeatureFlag      bool `json:"has_adblock_feature_flag"`
					HasAdblockGroup            bool `json:"has_adblock_group"`
					HasAdblockMinMobileVersion bool `json:"has_adblock_min_mobile_version"`
					HasAdblockThrottle         bool `json:"has_adblock_throttle"`
					PremiumPaymentFlowCapable  bool `json:"premium_payment_flow_capable"`
				} `json:"requirements"`
			} `json:"ad_block_viewable"`
			AdblockForProfiles struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"adblock_for_profiles"`
			AdblockProfileEntry struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"adblock_profile_entry"`
			AlexaSkill struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					InEnabledGroup          bool `json:"in_enabled_group"`
					IsAccountLinkingCapable bool `json:"is_account_linking_capable"`
					IsFeatureEnabled        bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"alexa_skill"`
			AllowBlockEdit struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"allow_block_edit"`
			AmazonDeviceNickname struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					IsAccountLinkingCapable bool `json:"is_account_linking_capable"`
					IsFeatureEnabled        bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"amazon_device_nickname"`
			BackupAccessPoint struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinFirmwareVersion bool `json:"has_min_firmware_version"`
				} `json:"requirements"`
			} `json:"backup_access_point"`
			BandSteering struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion bool `json:"has_min_version"`
				} `json:"requirements"`
			} `json:"band_steering"`
			Bifrost struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasBifrostAlgo      bool `json:"has_bifrost_algo"`
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
				} `json:"requirements"`
			} `json:"bifrost"`
			BlockApps struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"block_apps"`
			BlockAppsCategories struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"block_apps_categories"`
			CanAutoTrial struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					AutoTrialCapable bool `json:"auto_trial_capable"`
					Enabled          bool `json:"enabled"`
					IsRetail         bool `json:"is_retail"`
					IsTrialEligible  bool `json:"is_trial_eligible"`
					OwnerEligible    bool `json:"owner_eligible"`
					PremiumCapable   bool `json:"premium_capable"`
					SingleNetwork    bool `json:"single_network"`
				} `json:"requirements"`
			} `json:"can_auto_trial"`
			Cedar struct {
				Capable                     bool `json:"capable"`
				CapableWithUserRemediations bool `json:"capable_with_user_remediations"`
				Requirements                struct {
					FeatureFlagEnabled       bool `json:"feature_flag_enabled"`
					HasMinMobileVersion      bool `json:"has_min_mobile_version"`
					HasMinVersion            bool `json:"has_min_version"`
					InEnabledGroups          bool `json:"in_enabled_groups"`
					InRequiredConnectionMode bool `json:"in_required_connection_mode"`
					IsCertifiedGateway       bool `json:"is_certified_gateway"`
					IsFeatureEnabled         bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"cedar"`
			DdnsEnabled struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion       bool `json:"has_min_mobile_version"`
					HasMinUpsellMobileVersion bool `json:"has_min_upsell_mobile_version"`
					HasMinVersion             bool `json:"has_min_version"`
					HasSecurePlus             bool `json:"has_secure_plus"`
					IsEnabled                 bool `json:"is_enabled"`
					RequireSecurePlus         bool `json:"require_secure_plus"`
				} `json:"requirements"`
			} `json:"ddns_enabled"`
			Delorean struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
				} `json:"requirements"`
			} `json:"delorean"`
			DeviceBlacklist struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"device_blacklist"`
			DeviceManagement struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinVersion           bool `json:"has_min_version"`
					HasNodes                bool `json:"has_nodes"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"device_management"`
			DeviceUsage struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasFeatureFlag          bool `json:"has_feature_flag"`
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"device_usage"`
			Diagnostics struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
				} `json:"requirements"`
			} `json:"diagnostics"`
			DNSCaching struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"dns_caching"`
			DnsfilterAllowlist struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"dnsfilter_allowlist"`
			DnsfilterBlacklists struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"dnsfilter_blacklists"`
			DnsfilterBlocklist struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"dnsfilter_blocklist"`
			DnsfilterThreatCategories struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"dnsfilter_threat_categories"`
			DnsfilterWhitelists struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"dnsfilter_whitelists"`
			EeroBusinessReady struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					AllSupportedEeros bool `json:"all_supported_eeros"`
					NotInBridgeMode   bool `json:"not_in_bridge_mode"`
					NotInCustomMode   bool `json:"not_in_custom_mode"`
				} `json:"requirements"`
			} `json:"eero_business_ready"`
			EeroBusinessRetailUpsell struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					DoesNotHaveActiveBusinessLicense bool `json:"does_not_have_active_business_license"`
					HasRequiredNetworkCount          bool `json:"has_required_network_count"`
					IsNetworkOwner                   bool `json:"is_network_owner"`
					IsNotPremiumUser                 bool `json:"is_not_premium_user"`
					IsRetailUser                     bool `json:"is_retail_user"`
					IsSupportedLocation              bool `json:"is_supported_location"`
				} `json:"requirements"`
			} `json:"eero_business_retail_upsell"`
			EeroForBusinessCapable struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAvailableSeats   bool `json:"has_available_seats"`
					HasMinVersion       bool `json:"has_min_version"`
					IsSupportedLocation bool `json:"is_supported_location"`
				} `json:"requirements"`
			} `json:"eero_for_business_capable"`
			Ffs struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					IsAccountLinkingCapable bool `json:"is_account_linking_capable"`
					IsFeatureEnabled        bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"ffs"`
			HasAutoTrial struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					Enabled        bool `json:"enabled"`
					HasAutoTrial   bool `json:"has_auto_trial"`
					PremiumCapable bool `json:"premium_capable"`
				} `json:"requirements"`
			} `json:"has_auto_trial"`
			HistoricalInsights struct {
				Capable      bool `json:"capable"`
				Requirements struct {
				} `json:"requirements"`
			} `json:"historical_insights"`
			HistoricalUsage struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasFeatureFlag      bool `json:"has_feature_flag"`
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinNodeFirmware  bool `json:"has_min_node_firmware"`
					InEnabledGroups     bool `json:"in_enabled_groups"`
				} `json:"requirements"`
			} `json:"historical_usage"`
			HistoricalUsageNotifications struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasFeatureFlag         bool `json:"has_feature_flag"`
					HasMinMobileVersion    bool `json:"has_min_mobile_version"`
					HasMinNodeVersion      bool `json:"has_min_node_version"`
					HasRightConnectionMode bool `json:"has_right_connection_mode"`
					InEnabledGroups        bool `json:"in_enabled_groups"`
				} `json:"requirements"`
			} `json:"historical_usage_notifications"`
			Homekit struct {
				Capable                     bool `json:"capable"`
				CapableWithUserRemediations bool `json:"capable_with_user_remediations"`
				Requirements                struct {
					FeatureFlagEnabled       bool `json:"feature_flag_enabled"`
					HasMinMobileVersion      bool `json:"has_min_mobile_version"`
					HasMinVersion            bool `json:"has_min_version"`
					InEnabledGroups          bool `json:"in_enabled_groups"`
					InRequiredConnectionMode bool `json:"in_required_connection_mode"`
					IsCertifiedGateway       bool `json:"is_certified_gateway"`
					IsFeatureEnabled         bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"homekit"`
			ImprovedProfileCreationFlow struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"improved_profile_creation_flow"`
			ImprovedSwitchNetworks struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion         bool `json:"has_min_mobile_version"`
					HasProperOrganizationalRole bool `json:"has_proper_organizational_role"`
					IsEnabled                   bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"improved_switch_networks"`
			Ipv6 struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"ipv6"`
			Ipv6Editable struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasNoOrganizationOverride bool `json:"has_no_organization_override"`
					IsIpv6Capable             bool `json:"is_ipv6_capable"`
				} `json:"requirements"`
			} `json:"ipv6_editable"`
			IsAndroidWebPayments struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"is_android_web_payments"`
			LedAction struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
				} `json:"requirements"`
			} `json:"led_action"`
			NewPrivateDevicesNotifications struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					InEnabledGroups  bool `json:"in_enabled_groups"`
					IsFeatureEnabled bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"new_private_devices_notifications"`
			OrgAutoAssociatesNetworks struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"org_auto_associates_networks"`
			OwnershipTransfer struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion            bool `json:"has_min_mobile_version"`
					HasOwnershipTransferPermission bool `json:"has_ownership_transfer_permission"`
					IsEnabled                      bool `json:"is_enabled"`
					IsIspManaged                   bool `json:"is_isp_managed"`
				} `json:"requirements"`
			} `json:"ownership_transfer"`
			PerDeviceInsights struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"per_device_insights"`
			PortForwardRange struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"port_forward_range"`
			PostSetupWanTroubleshooting struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					GatewayHasMinFirmwareVersion bool `json:"gateway_has_min_firmware_version"`
				} `json:"requirements"`
			} `json:"post_setup_wan_troubleshooting"`
			PowerSaving struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					Firmware                            string `json:"firmware"`
					GatewayModel                        string `json:"gateway_model"`
					HasSupportedFirmware                bool   `json:"has_supported_firmware"`
					HasSupportedHardware                bool   `json:"has_supported_hardware"`
					IsCountryCapableOrPreReleaseNetwork bool   `json:"is_country_capable_or_pre_release_network"`
				} `json:"requirements"`
			} `json:"power_saving"`
			Pppoe struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasGateway         bool `json:"has_gateway"`
					HasMinimumFirmware bool `json:"has_minimum_firmware"`
				} `json:"requirements"`
			} `json:"pppoe"`
			Premium struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedCountryCode      bool `json:"has_allowed_country_code"`
					HasAllowedGroup            bool `json:"has_allowed_group"`
					HasMinMobileVersion        bool `json:"has_min_mobile_version"`
					HasMinVersion              bool `json:"has_min_version"`
					HasProperConnectionMode    bool `json:"has_proper_connection_mode"`
					HasProperOrgConfigs        bool `json:"has_proper_org_configs"`
					HasProperOwnedNetworkCount bool `json:"has_proper_owned_network_count"`
					HasProperOwnerRole         bool `json:"has_proper_owner_role"`
				} `json:"requirements"`
			} `json:"premium"`
			PremiumBranding struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg           bool `json:"has_allowed_org"`
					IsPremiumPaymentCapable bool `json:"is_premium_payment_capable"`
				} `json:"requirements"`
			} `json:"premium_branding"`
			PremiumCancelImmediately struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"premium_cancel_immediately"`
			PremiumIspPlanEnabled struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"premium_isp_plan_enabled"`
			PremiumIspSelfSignup struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"premium_isp_self_signup"`
			PremiumManagement struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"premium_management"`
			PremiumPartnerships struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg    bool `json:"has_allowed_org"`
					IsPremiumCapable bool `json:"is_premium_capable"`
				} `json:"requirements"`
			} `json:"premium_partnerships"`
			PremiumPaymentFlow struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedCountryCode      bool `json:"has_allowed_country_code"`
					HasAllowedGroup            bool `json:"has_allowed_group"`
					HasMinMobileVersion        bool `json:"has_min_mobile_version"`
					HasProperConnectionMode    bool `json:"has_proper_connection_mode"`
					HasProperOwnedNetworkCount bool `json:"has_proper_owned_network_count"`
					HasProperOwnerRole         bool `json:"has_proper_owner_role"`
				} `json:"requirements"`
			} `json:"premium_payment_flow"`
			PremiumUpsell struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasAllowedOrg            bool `json:"has_allowed_org"`
					IsEeroPlusUpgradeEnabled bool `json:"is_eero_plus_upgrade_enabled"`
					IsPremiumPaymentCapable  bool `json:"is_premium_payment_capable"`
				} `json:"requirements"`
			} `json:"premium_upsell"`
			ProxiedNodesBeta1 struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
					IsCountryAllowed        bool `json:"is_country_allowed"`
					IsFeatureEnabled        bool `json:"is_feature_enabled"`
					IsGatewayNodeEligible   bool `json:"is_gateway_node_eligible"`
					NetworkHasWifiNodes     bool `json:"network_has_wifi_nodes"`
				} `json:"requirements"`
			} `json:"proxied_nodes_beta_1"`
			ProxiedNodesBeta2 struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
					IsCountryAllowed        bool `json:"is_country_allowed"`
					IsFeatureEnabled        bool `json:"is_feature_enabled"`
					IsGatewayNodeEligible   bool `json:"is_gateway_node_eligible"`
					NetworkHasWifiNodes     bool `json:"network_has_wifi_nodes"`
				} `json:"requirements"`
			} `json:"proxied_nodes_beta_2"`
			PushNotificationSettingActivityReport struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
					IsPremiumUser       bool `json:"is_premium_user"`
				} `json:"requirements"`
			} `json:"push_notification_setting_activity_report"`
			RingLte struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					InRequiredConnectionMode bool `json:"in_required_connection_mode"`
					IsGatewayKilimanjaro     bool `json:"is_gateway_kilimanjaro"`
				} `json:"requirements"`
			} `json:"ring_lte"`
			SimpleSetup struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion            bool `json:"has_min_version"`
					InEnabledCountries       bool `json:"in_enabled_countries"`
					InEnabledGroups          bool `json:"in_enabled_groups"`
					IsFeatureEnabled         bool `json:"is_feature_enabled"`
					NotAccountLinkingCapable bool `json:"not_account_linking_capable"`
				} `json:"requirements"`
			} `json:"simple_setup"`
			SmartHomeHub struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					IsAccountLinkingCapable  bool `json:"is_account_linking_capable"`
					IsFeatureEnabled         bool `json:"is_feature_enabled"`
					MeetsHardwareRequirement bool `json:"meets_hardware_requirement"`
				} `json:"requirements"`
			} `json:"smart_home_hub"`
			SoftwareEndOfLife struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasFeatureFlag      bool `json:"has_feature_flag"`
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
				} `json:"requirements"`
			} `json:"software_end_of_life"`
			Sqm struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
				} `json:"requirements"`
			} `json:"sqm"`
			ThreadCommissioning struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasBorderAgent         bool `json:"has_border_agent"`
					HasMinMobileVersion    bool `json:"has_min_mobile_version"`
					IsActive               bool `json:"is_active"`
					IsEnabled              bool `json:"is_enabled"`
					IsThreadNetworkCapable bool `json:"is_thread_network_capable"`
				} `json:"requirements"`
			} `json:"thread_commissioning"`
			ThreadKeychainSharing struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					FeatureEnabled          bool `json:"feature_enabled"`
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinVersion           bool `json:"has_min_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
					HasSupportedHardware    bool `json:"has_supported_hardware"`
				} `json:"requirements"`
			} `json:"thread_keychain_sharing"`
			ThreadNetwork struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasFeatureFlag          bool `json:"has_feature_flag"`
					HasMinMobileVersion     bool `json:"has_min_mobile_version"`
					HasMinNodeVersion       bool `json:"has_min_node_version"`
					HasMinThreadVersion     bool `json:"has_min_thread_version"`
					HasProperConnectionMode bool `json:"has_proper_connection_mode"`
					HasSupportedHardware    bool `json:"has_supported_hardware"`
				} `json:"requirements"`
			} `json:"thread_network"`
			UDPSpeedTests struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					GatewayModel   string `json:"gateway_model"`
					GatewayVersion string `json:"gateway_version"`
					MinimumVersion any    `json:"minimum_version"`
				} `json:"requirements"`
			} `json:"udp_speed_tests"`
			UnifiedContentFilters struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinMobileVersion bool `json:"has_min_mobile_version"`
					HasMinVersion       bool `json:"has_min_version"`
					IsEnabled           bool `json:"is_enabled"`
				} `json:"requirements"`
			} `json:"unified_content_filters"`
			Vlan struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasGateway         bool `json:"has_gateway"`
					HasMinimumFirmware bool `json:"has_minimum_firmware"`
				} `json:"requirements"`
			} `json:"vlan"`
			Wpa3 struct {
				Capable      bool `json:"capable"`
				Requirements struct {
					HasMinVersion    bool `json:"has_min_version"`
					InEnabledGroups  bool `json:"in_enabled_groups"`
					IsFeatureEnabled bool `json:"is_feature_enabled"`
				} `json:"requirements"`
			} `json:"wpa3"`
		} `json:"capabilities"`
		Clients struct {
			Count int    `json:"count"`
			URL   string `json:"url"`
		} `json:"clients"`
		Connection struct {
			Mode string `json:"mode"`
		} `json:"connection"`
		Ddns struct {
			Enabled   bool   `json:"enabled"`
			Subdomain string `json:"subdomain"`
		} `json:"ddns"`
		Dhcp struct {
			Custom any    `json:"custom"`
			Mode   string `json:"mode"`
		} `json:"dhcp"`
		DisplayName string `json:"display_name"`
		DNS         struct {
			Caching bool   `json:"caching"`
			Custom  any    `json:"custom"`
			Mode    string `json:"mode"`
			Parent  struct {
				Ips []string `json:"ips"`
			} `json:"parent"`
		} `json:"dns"`
		Eeros struct {
			Count int `json:"count"`
			Data  []struct {
				BackupWan                     any      `json:"backup_wan"`
				Bands                         []string `json:"bands"`
				ConnectedClientsCount         int      `json:"connected_clients_count"`
				ConnectedWiredClientsCount    int      `json:"connected_wired_clients_count"`
				ConnectedWirelessClientsCount int      `json:"connected_wireless_clients_count"`
				ConnectionType                string   `json:"connection_type"`
				EthernetAddresses             []string `json:"ethernet_addresses"`
				EthernetStatus                struct {
					SegmentID string `json:"segmentId"`
					Statuses  []struct {
						DeratedReason         any    `json:"derated_reason"`
						HasCarrier            bool   `json:"hasCarrier"`
						InterfaceNumber       int    `json:"interfaceNumber"`
						IsLeafWiredToUpstream bool   `json:"isLeafWiredToUpstream"`
						IsLte                 bool   `json:"isLte"`
						IsWanPort             bool   `json:"isWanPort"`
						Neighbor              any    `json:"neighbor"`
						OriginalSpeed         any    `json:"original_speed"`
						PowerSaving           bool   `json:"power_saving"`
						Speed                 string `json:"speed"`
					} `json:"statuses"`
					WiredInternet bool `json:"wiredInternet"`
				} `json:"ethernet_status"`
				ExtendedBorderAgentAddress string `json:"extended_border_agent_address"`
				Gateway                    bool   `json:"gateway"`
				HeartbeatOk                bool   `json:"heartbeat_ok"`
				IPAddress                  string `json:"ip_address"`
				Ipv6Addresses              []struct {
					Address   string `json:"address"`
					Interface string `json:"interface"`
					Scope     string `json:"scope"`
				} `json:"ipv6_addresses"`
				IsPrimaryNode   bool      `json:"is_primary_node"`
				Joined          string    `json:"joined"`
				LastHeartbeat   time.Time `json:"last_heartbeat"`
				LastReboot      time.Time `json:"last_reboot"`
				LedBrightness   int       `json:"led_brightness"`
				LedOn           bool      `json:"led_on"`
				Location        string    `json:"location"`
				MacAddress      string    `json:"mac_address"`
				MeshQualityBars int       `json:"mesh_quality_bars"`
				Messages        []any     `json:"messages"`
				Model           string    `json:"model"`
				ModelNumber     string    `json:"model_number"`
				Network         struct {
					Created time.Time `json:"created"`
					Name    string    `json:"name"`
					URL     string    `json:"url"`
				} `json:"network"`
				Nightlight   any `json:"nightlight"`
				Organization struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"organization"`
				Os          string `json:"os"`
				OsVersion   string `json:"os_version"`
				PortDetails []struct {
					EthernetAddress string `json:"ethernet_address"`
					PortName        string `json:"port_name"`
					Position        int    `json:"position"`
				} `json:"port_details"`
				PowerInfo struct {
					PowerSource         string `json:"power_source"`
					PowerSourceMetadata struct {
					} `json:"power_source_metadata"`
				} `json:"power_info"`
				ProvideDevicePower              any  `json:"provide_device_power"`
				ProvidesWifi                    bool `json:"provides_wifi"`
				RequiresAmazonPreAuthorizedCode bool `json:"requires_amazon_pre_authorized_code"`
				Resources                       struct {
					Connections string `json:"connections"`
					LedAction   string `json:"led_action"`
					Reboot      string `json:"reboot"`
				} `json:"resources"`
				Serial          string `json:"serial"`
				Status          string `json:"status"`
				UpdateAvailable bool   `json:"update_available"`
				UpdateStatus    struct {
					SupportExpirationDate   any    `json:"support_expiration_date"`
					SupportExpirationString string `json:"support_expiration_string"`
					SupportExpired          bool   `json:"support_expired"`
				} `json:"update_status"`
				URL        string   `json:"url"`
				UsingWan   bool     `json:"using_wan"`
				WifiBssids []string `json:"wifi_bssids"`
				Wired      bool     `json:"wired"`
			} `json:"data"`
		} `json:"eeros"`
		Ffs   bool `json:"ffs"`
		Flags []struct {
			Flag  string `json:"flag"`
			Value bool   `json:"value"`
		} `json:"flags"`
		Gateway   string `json:"gateway"`
		GatewayIP string `json:"gateway_ip"`
		GeoIP     struct {
			AreaCode    any    `json:"areaCode"`
			Asn         int    `json:"asn"`
			City        string `json:"city"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			Isp         string `json:"isp"`
			MetroCode   any    `json:"metroCode"`
			Org         string `json:"org"`
			PostalCode  string `json:"postalCode"`
			Region      string `json:"region"`
			RegionName  string `json:"regionName"`
			Timezone    string `json:"timezone"`
		} `json:"geo_ip"`
		GuestNetwork struct {
			Enabled   bool   `json:"enabled"`
			Name      string `json:"name"`
			Password  string `json:"password"`
			Resources struct {
				Password string `json:"password"`
			} `json:"resources"`
			URL string `json:"url"`
		} `json:"guest_network"`
		Health struct {
			EeroNetwork struct {
				Status string `json:"status"`
			} `json:"eero_network"`
			Internet struct {
				IspUp  bool   `json:"isp_up"`
				Status string `json:"status"`
			} `json:"internet"`
		} `json:"health"`
		Homekit     any `json:"homekit"`
		ImageAssets struct {
			Description string `json:"description"`
			Expires     any    `json:"expires"`
			Hash        string `json:"hash"`
			ID          int    `json:"id"`
			URL         string `json:"url"`
		} `json:"image_assets"`
		IPSettings struct {
			DoubleNat bool   `json:"double_nat"`
			PublicIP  string `json:"public_ip"`
		} `json:"ip_settings"`
		Ipv6 struct {
			NameServers struct {
				Custom []any  `json:"custom"`
				Mode   string `json:"mode"`
			} `json:"name_servers"`
		} `json:"ipv6"`
		Ipv6Lease    any       `json:"ipv6_lease"`
		Ipv6Upstream bool      `json:"ipv6_upstream"`
		LastReboot   time.Time `json:"last_reboot"`
		Lease        struct {
			Dhcp struct {
				IP     string `json:"ip"`
				Mask   string `json:"mask"`
				Router string `json:"router"`
			} `json:"dhcp"`
			Mode   string `json:"mode"`
			Static any    `json:"static"`
		} `json:"lease"`
		Messages            []any  `json:"messages"`
		Name                string `json:"name"`
		NetworkCustomerType string `json:"network_customer_type"`
		NicknameLabel       any    `json:"nickname_label"`
		Organization        struct {
			Brand string `json:"brand"`
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Type  string `json:"type"`
		} `json:"organization"`
		Owner          string `json:"owner"`
		Password       string `json:"password"`
		PowerSaving    bool   `json:"power_saving"`
		PppoeEnabled   any    `json:"pppoe_enabled"`
		PppoeUsername  any    `json:"pppoe_username"`
		PremiumDetails struct {
			HasPaymentInfo       bool   `json:"has_payment_info"`
			Interval             string `json:"interval"`
			IsMySubscription     bool   `json:"is_my_subscription"`
			NextBillingEventDate any    `json:"next_billing_event_date"`
			PaymentMethod        any    `json:"payment_method"`
			Tier                 string `json:"tier"`
			TrialEnds            any    `json:"trial_ends"`
		} `json:"premium_details"`
		PremiumDNS struct {
			AdBlockSettings struct {
				BusinessSubnets []any `json:"business_subnets"`
				Enabled         bool  `json:"enabled"`
				Profiles        []any `json:"profiles"`
			} `json:"ad_block_settings"`
			AdvancedContentFilters       any  `json:"advanced_content_filters"`
			AnyPoliciesEnabledForNetwork bool `json:"any_policies_enabled_for_network"`
			DNSPolicies                  struct {
				AdBlock      bool `json:"ad_block"`
				BlockMalware bool `json:"block_malware"`
			} `json:"dns_policies"`
			DNSPoliciesEnabled     bool   `json:"dns_policies_enabled"`
			DNSProvider            string `json:"dns_provider"`
			ZscalerLocationEnabled bool   `json:"zscaler_location_enabled"`
		} `json:"premium_dns"`
		PremiumStatus string `json:"premium_status"`
		ProxiedNodes  any    `json:"proxied_nodes"`
		Rebooting     any    `json:"rebooting"`
		Resources     struct {
			AcCompat        string `json:"ac_compat"`
			BurstReporters  string `json:"burst_reporters"`
			DeviceBlacklist string `json:"device_blacklist"`
			Devices         string `json:"devices"`
			Diagnostics     string `json:"diagnostics"`
			Eeros           string `json:"eeros"`
			Forwards        string `json:"forwards"`
			Guestnetwork    string `json:"guestnetwork"`
			Insights        string `json:"insights"`
			Ouicheck        string `json:"ouicheck"`
			Password        string `json:"password"`
			Profiles        string `json:"profiles"`
			Reboot          string `json:"reboot"`
			Reservations    string `json:"reservations"`
			Routing         string `json:"routing"`
			Settings        string `json:"settings"`
			Speedtest       string `json:"speedtest"`
			Support         string `json:"support"`
			Thread          string `json:"thread"`
			Transfer        string `json:"transfer"`
			Updates         string `json:"updates"`
		} `json:"resources"`
		RingLte any `json:"ring_lte"`
		Speed   struct {
			Date time.Time `json:"date"`
			Down struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"down"`
			Status string `json:"status"`
			Up     struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"up"`
		} `json:"speed"`
		Sqm            bool   `json:"sqm"`
		Status         string `json:"status"`
		TemporaryFlags struct {
		} `json:"temporary_flags"`
		Thread   bool `json:"thread"`
		Timezone struct {
			GeoIP string `json:"geo_ip"`
			Value string `json:"value"`
		} `json:"timezone"`
		Updates struct {
			CanUpdateNow      bool      `json:"can_update_now"`
			HasUpdate         bool      `json:"has_update"`
			LastUpdateStarted time.Time `json:"last_update_started"`
			LastUserUpdate    struct {
				IncompleteEeros   []any     `json:"incomplete_eeros"`
				LastUpdateStarted time.Time `json:"last_update_started"`
				UnresponsiveEeros []any     `json:"unresponsive_eeros"`
			} `json:"last_user_update"`
			ManifestResource    string `json:"manifest_resource"`
			MinRequiredFirmware string `json:"min_required_firmware"`
			PreferredUpdateHour int    `json:"preferred_update_hour"`
			ScheduledUpdateTime any    `json:"scheduled_update_time"`
			TargetFirmware      string `json:"target_firmware"`
			UpdateRequired      bool   `json:"update_required"`
			UpdateStatus        any    `json:"update_status"`
			UpdateToFirmware    string `json:"update_to_firmware"`
		} `json:"updates"`
		Upnp     bool   `json:"upnp"`
		Upstream []any  `json:"upstream"`
		URL      string `json:"url"`
		Vlan     any    `json:"vlan"`
		WanIP    string `json:"wan_ip"`
		WanType  string `json:"wan_type"`
		Wpa3     bool   `json:"wpa3"`
	} `json:"data"`
	Meta struct {
		Code       int       `json:"code"`
		ServerTime time.Time `json:"server_time"`
	} `json:"meta"`
}
package googleapigeesecurityaction


type GoogleApigeeSecurityActionConditionConfig struct {
	// A list of accessTokens. Limit 1000 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#access_tokens GoogleApigeeSecurityAction#access_tokens}
	AccessTokens *[]*string `field:"optional" json:"accessTokens" yaml:"accessTokens"`
	// A list of API keys. Limit 1000 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#api_keys GoogleApigeeSecurityAction#api_keys}
	ApiKeys *[]*string `field:"optional" json:"apiKeys" yaml:"apiKeys"`
	// A list of API Products. Limit 1000 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#api_products GoogleApigeeSecurityAction#api_products}
	ApiProducts *[]*string `field:"optional" json:"apiProducts" yaml:"apiProducts"`
	// A list of ASN numbers to act on, e.g. 23. https://en.wikipedia.org/wiki/Autonomous_system_(Internet) This uses int64 instead of uint32 because of https://linter.aip.dev/141/forbidden-types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#asns GoogleApigeeSecurityAction#asns}
	Asns *[]*string `field:"optional" json:"asns" yaml:"asns"`
	// A list of Bot Reasons.
	//
	// Current options: Flooder, Brute Guessor, Static Content Scraper,
	// OAuth Abuser, Robot Abuser, TorListRule, Advanced Anomaly Detection, Advanced API Scraper,
	// Search Engine Crawlers, Public Clouds, Public Cloud AWS, Public Cloud Azure, and Public Cloud Google.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#bot_reasons GoogleApigeeSecurityAction#bot_reasons}
	BotReasons *[]*string `field:"optional" json:"botReasons" yaml:"botReasons"`
	// A list of developer apps. Limit 1000 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#developer_apps GoogleApigeeSecurityAction#developer_apps}
	DeveloperApps *[]*string `field:"optional" json:"developerApps" yaml:"developerApps"`
	// A list of developers. Limit 1000 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#developers GoogleApigeeSecurityAction#developers}
	Developers *[]*string `field:"optional" json:"developers" yaml:"developers"`
	// Act only on particular HTTP methods.
	//
	// E.g. A read-only API can block POST/PUT/DELETE methods.
	// Accepted values are: GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE and PATCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#http_methods GoogleApigeeSecurityAction#http_methods}
	HttpMethods *[]*string `field:"optional" json:"httpMethods" yaml:"httpMethods"`
	// A list of IP addresses. This could be either IPv4 or IPv6. Limited to 100 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#ip_address_ranges GoogleApigeeSecurityAction#ip_address_ranges}
	IpAddressRanges *[]*string `field:"optional" json:"ipAddressRanges" yaml:"ipAddressRanges"`
	// A list of countries/region codes to act on, e.g. US. This follows https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#region_codes GoogleApigeeSecurityAction#region_codes}
	RegionCodes *[]*string `field:"optional" json:"regionCodes" yaml:"regionCodes"`
	// A list of user agents to deny. We look for exact matches. Limit 50 per action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#user_agents GoogleApigeeSecurityAction#user_agents}
	UserAgents *[]*string `field:"optional" json:"userAgents" yaml:"userAgents"`
}


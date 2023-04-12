package provider


type OktaProviderConfig struct {
	// Bearer token granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#access_token OktaProvider#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#alias OktaProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// API Token granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#api_token OktaProvider#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Use exponential back off strategy for rate limits.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#backoff OktaProvider#backoff}
	Backoff interface{} `field:"optional" json:"backoff" yaml:"backoff"`
	// The Okta url. (Use 'oktapreview.com' for Okta testing).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#base_url OktaProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// API Token granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#client_id OktaProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Alternate HTTP proxy of scheme://hostname or scheme://hostname:port format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#http_proxy OktaProvider#http_proxy}
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// providers log level. Minimum is 1 (TRACE), and maximum is 5 (ERROR).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#log_level OktaProvider#log_level}
	LogLevel *float64 `field:"optional" json:"logLevel" yaml:"logLevel"`
	// (Experimental) sets what percentage of capacity the provider can use of the total rate limit capacity while making calls to the Okta management API endpoints.
	//
	// Okta API operates in one minute buckets. See Okta Management API Rate Limits: https://developer.okta.com/docs/reference/rl-global-mgmt/
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#max_api_capacity OktaProvider#max_api_capacity}
	MaxApiCapacity *float64 `field:"optional" json:"maxApiCapacity" yaml:"maxApiCapacity"`
	// maximum number of retries to attempt before erroring out.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#max_retries OktaProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// maximum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#max_wait_seconds OktaProvider#max_wait_seconds}
	MaxWaitSeconds *float64 `field:"optional" json:"maxWaitSeconds" yaml:"maxWaitSeconds"`
	// minimum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#min_wait_seconds OktaProvider#min_wait_seconds}
	MinWaitSeconds *float64 `field:"optional" json:"minWaitSeconds" yaml:"minWaitSeconds"`
	// The organization to manage in Okta.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#org_name OktaProvider#org_name}
	OrgName *string `field:"optional" json:"orgName" yaml:"orgName"`
	// Number of concurrent requests to make within a resource where bulk operations are not possible. Take note of https://developer.okta.com/docs/api/getting_started/rate-limits.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#parallelism OktaProvider#parallelism}
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// API Token granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#private_key OktaProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// API Token Id granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#private_key_id OktaProvider#private_key_id}
	PrivateKeyId *string `field:"optional" json:"privateKeyId" yaml:"privateKeyId"`
	// Timeout for single request (in seconds) which is made to Okta, the default is `0` (means no limit is set).
	//
	// The maximum value can be `300`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#request_timeout OktaProvider#request_timeout}
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// API Token granting privileges to Okta API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta#scopes OktaProvider#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}


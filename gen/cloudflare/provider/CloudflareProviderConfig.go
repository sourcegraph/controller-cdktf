package provider


type CloudflareProviderConfig struct {
	// Configure API client to always use that account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#account_id CloudflareProvider#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#alias CloudflareProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Configure the base path used by the API client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_base_path CloudflareProvider#api_base_path}
	ApiBasePath *string `field:"optional" json:"apiBasePath" yaml:"apiBasePath"`
	// Whether to print logs from the API client (using the default log library logger).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_client_logging CloudflareProvider#api_client_logging}
	ApiClientLogging interface{} `field:"optional" json:"apiClientLogging" yaml:"apiClientLogging"`
	// Configure the hostname used by the API client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_hostname CloudflareProvider#api_hostname}
	ApiHostname *string `field:"optional" json:"apiHostname" yaml:"apiHostname"`
	// The API key for operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_key CloudflareProvider#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The API Token for operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_token CloudflareProvider#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// A special Cloudflare API key good for a restricted set of endpoints.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#api_user_service_key CloudflareProvider#api_user_service_key}
	ApiUserServiceKey *string `field:"optional" json:"apiUserServiceKey" yaml:"apiUserServiceKey"`
	// A registered Cloudflare email address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#email CloudflareProvider#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Maximum backoff period in seconds after failed API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#max_backoff CloudflareProvider#max_backoff}
	MaxBackoff *float64 `field:"optional" json:"maxBackoff" yaml:"maxBackoff"`
	// Minimum backoff period in seconds after failed API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#min_backoff CloudflareProvider#min_backoff}
	MinBackoff *float64 `field:"optional" json:"minBackoff" yaml:"minBackoff"`
	// Maximum number of retries to perform when an API request fails.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#retries CloudflareProvider#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// RPS limit to apply when making calls to the API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare#rps CloudflareProvider#rps}
	Rps *float64 `field:"optional" json:"rps" yaml:"rps"`
}


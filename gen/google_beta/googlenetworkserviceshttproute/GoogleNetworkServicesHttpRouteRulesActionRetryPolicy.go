package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesActionRetryPolicy struct {
	// Specifies the allowed number of retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#num_retries GoogleNetworkServicesHttpRoute#num_retries}
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Specifies a non-zero timeout per retry attempt.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#per_try_timeout GoogleNetworkServicesHttpRoute#per_try_timeout}
	PerTryTimeout *string `field:"optional" json:"perTryTimeout" yaml:"perTryTimeout"`
	// Specifies one or more conditions when this retry policy applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#retry_conditions GoogleNetworkServicesHttpRoute#retry_conditions}
	RetryConditions *[]*string `field:"optional" json:"retryConditions" yaml:"retryConditions"`
}


package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#cors_policy GoogleNetworkServicesHttpRoute#cors_policy}
	CorsPolicy *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#destinations GoogleNetworkServicesHttpRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#fault_injection_policy GoogleNetworkServicesHttpRoute#fault_injection_policy}
	FaultInjectionPolicy *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#redirect GoogleNetworkServicesHttpRoute#redirect}
	Redirect *GoogleNetworkServicesHttpRouteRulesActionRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// request_header_modifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#request_header_modifier GoogleNetworkServicesHttpRoute#request_header_modifier}
	RequestHeaderModifier *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier `field:"optional" json:"requestHeaderModifier" yaml:"requestHeaderModifier"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#request_mirror_policy GoogleNetworkServicesHttpRoute#request_mirror_policy}
	RequestMirrorPolicy *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// response_header_modifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#response_header_modifier GoogleNetworkServicesHttpRoute#response_header_modifier}
	ResponseHeaderModifier *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier `field:"optional" json:"responseHeaderModifier" yaml:"responseHeaderModifier"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#retry_policy GoogleNetworkServicesHttpRoute#retry_policy}
	RetryPolicy *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Specifies the timeout for selected route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#timeout GoogleNetworkServicesHttpRoute#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#url_rewrite GoogleNetworkServicesHttpRoute#url_rewrite}
	UrlRewrite *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}


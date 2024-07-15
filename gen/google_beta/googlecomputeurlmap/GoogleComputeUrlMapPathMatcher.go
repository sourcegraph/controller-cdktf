package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcher struct {
	// The name to which this PathMatcher is referred by the HostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#name GoogleComputeUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// default_custom_error_response_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#default_custom_error_response_policy GoogleComputeUrlMap#default_custom_error_response_policy}
	DefaultCustomErrorResponsePolicy *GoogleComputeUrlMapPathMatcherDefaultCustomErrorResponsePolicy `field:"optional" json:"defaultCustomErrorResponsePolicy" yaml:"defaultCustomErrorResponsePolicy"`
	// default_route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#default_route_action GoogleComputeUrlMap#default_route_action}
	DefaultRouteAction *GoogleComputeUrlMapPathMatcherDefaultRouteAction `field:"optional" json:"defaultRouteAction" yaml:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#default_service GoogleComputeUrlMap#default_service}
	DefaultService *string `field:"optional" json:"defaultService" yaml:"defaultService"`
	// default_url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#default_url_redirect GoogleComputeUrlMap#default_url_redirect}
	DefaultUrlRedirect *GoogleComputeUrlMapPathMatcherDefaultUrlRedirect `field:"optional" json:"defaultUrlRedirect" yaml:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#description GoogleComputeUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#header_action GoogleComputeUrlMap#header_action}
	HeaderAction *GoogleComputeUrlMapPathMatcherHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// path_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#path_rule GoogleComputeUrlMap#path_rule}
	PathRule interface{} `field:"optional" json:"pathRule" yaml:"pathRule"`
	// route_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#route_rules GoogleComputeUrlMap#route_rules}
	RouteRules interface{} `field:"optional" json:"routeRules" yaml:"routeRules"`
}


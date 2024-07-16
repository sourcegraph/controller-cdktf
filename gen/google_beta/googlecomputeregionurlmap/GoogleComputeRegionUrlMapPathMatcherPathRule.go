package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherPathRule struct {
	// The list of path patterns to match.
	//
	// Each must start with / and the only place a
	// \* is allowed is at the end following a /. The string fed to the path matcher
	// does not include any text after the first ? or #, and those chars are not
	// allowed here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#paths GoogleComputeRegionUrlMap#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#route_action GoogleComputeRegionUrlMap#route_action}
	RouteAction *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteAction `field:"optional" json:"routeAction" yaml:"routeAction"`
	// The region backend service resource to which traffic is directed if this rule is matched.
	//
	// If routeAction is additionally specified,
	// advanced routing actions like URL Rewrites, etc. take effect prior to sending
	// the request to the backend. However, if service is specified, routeAction cannot
	// contain any weightedBackendService s. Conversely, if routeAction specifies any
	// weightedBackendServices, service must not be specified. Only one of urlRedirect,
	// service or routeAction.weightedBackendService must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#service GoogleComputeRegionUrlMap#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#url_redirect GoogleComputeRegionUrlMap#url_redirect}
	UrlRedirect *GoogleComputeRegionUrlMapPathMatcherPathRuleUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
}


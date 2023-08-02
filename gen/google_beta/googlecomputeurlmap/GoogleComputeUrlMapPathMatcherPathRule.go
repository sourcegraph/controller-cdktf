package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRule struct {
	// The list of path patterns to match.
	//
	// Each must start with / and the only place a
	// \* is allowed is at the end following a /. The string fed to the path matcher
	// does not include any text after the first ? or #, and those chars are not
	// allowed here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_url_map#paths GoogleComputeUrlMap#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_url_map#route_action GoogleComputeUrlMap#route_action}
	RouteAction *GoogleComputeUrlMapPathMatcherPathRuleRouteAction `field:"optional" json:"routeAction" yaml:"routeAction"`
	// The backend service or backend bucket to use if any of the given paths match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_url_map#service GoogleComputeUrlMap#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_url_map#url_redirect GoogleComputeUrlMap#url_redirect}
	UrlRedirect *GoogleComputeUrlMapPathMatcherPathRuleUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
}


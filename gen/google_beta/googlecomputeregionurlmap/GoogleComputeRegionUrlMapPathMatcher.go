package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcher struct {
	// The name to which this PathMatcher is referred by the HostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#name GoogleComputeRegionUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// default_route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#default_route_action GoogleComputeRegionUrlMap#default_route_action}
	DefaultRouteAction *GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction `field:"optional" json:"defaultRouteAction" yaml:"defaultRouteAction"`
	// A reference to a RegionBackendService resource.
	//
	// This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#default_service GoogleComputeRegionUrlMap#default_service}
	DefaultService *string `field:"optional" json:"defaultService" yaml:"defaultService"`
	// default_url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#default_url_redirect GoogleComputeRegionUrlMap#default_url_redirect}
	DefaultUrlRedirect *GoogleComputeRegionUrlMapPathMatcherDefaultUrlRedirect `field:"optional" json:"defaultUrlRedirect" yaml:"defaultUrlRedirect"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#description GoogleComputeRegionUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#header_action GoogleComputeRegionUrlMap#header_action}
	HeaderAction *GoogleComputeRegionUrlMapPathMatcherHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// path_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#path_rule GoogleComputeRegionUrlMap#path_rule}
	PathRule interface{} `field:"optional" json:"pathRule" yaml:"pathRule"`
	// route_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#route_rules GoogleComputeRegionUrlMap#route_rules}
	RouteRules interface{} `field:"optional" json:"routeRules" yaml:"routeRules"`
}


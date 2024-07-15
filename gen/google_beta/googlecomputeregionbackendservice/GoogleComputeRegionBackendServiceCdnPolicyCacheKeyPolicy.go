package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy struct {
	// If true requests to different hosts will be cached separately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#include_host GoogleComputeRegionBackendService#include_host}
	IncludeHost interface{} `field:"optional" json:"includeHost" yaml:"includeHost"`
	// Names of cookies to include in cache keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#include_named_cookies GoogleComputeRegionBackendService#include_named_cookies}
	IncludeNamedCookies *[]*string `field:"optional" json:"includeNamedCookies" yaml:"includeNamedCookies"`
	// If true, http and https requests will be cached separately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#include_protocol GoogleComputeRegionBackendService#include_protocol}
	IncludeProtocol interface{} `field:"optional" json:"includeProtocol" yaml:"includeProtocol"`
	// If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist.
	//
	// If neither is set, the entire query
	// string will be included.
	//
	// If false, the query string will be excluded from the cache
	// key entirely.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#include_query_string GoogleComputeRegionBackendService#include_query_string}
	IncludeQueryString interface{} `field:"optional" json:"includeQueryString" yaml:"includeQueryString"`
	// Names of query string parameters to exclude in cache keys.
	//
	// All other parameters will be included. Either specify
	// query_string_whitelist or query_string_blacklist, not both.
	// '&' and '=' will be percent encoded and not treated as
	// delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#query_string_blacklist GoogleComputeRegionBackendService#query_string_blacklist}
	QueryStringBlacklist *[]*string `field:"optional" json:"queryStringBlacklist" yaml:"queryStringBlacklist"`
	// Names of query string parameters to include in cache keys.
	//
	// All other parameters will be excluded. Either specify
	// query_string_whitelist or query_string_blacklist, not both.
	// '&' and '=' will be percent encoded and not treated as
	// delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#query_string_whitelist GoogleComputeRegionBackendService#query_string_whitelist}
	QueryStringWhitelist *[]*string `field:"optional" json:"queryStringWhitelist" yaml:"queryStringWhitelist"`
}


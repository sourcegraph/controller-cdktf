package googlecomputebackendbucket


type GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy struct {
	// Allows HTTP request headers (by name) to be used in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket#include_http_headers GoogleComputeBackendBucket#include_http_headers}
	IncludeHttpHeaders *[]*string `field:"optional" json:"includeHttpHeaders" yaml:"includeHttpHeaders"`
	// Names of query string parameters to include in cache keys.
	//
	// Default parameters are always included. '&' and '=' will
	// be percent encoded and not treated as delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket#query_string_whitelist GoogleComputeBackendBucket#query_string_whitelist}
	QueryStringWhitelist *[]*string `field:"optional" json:"queryStringWhitelist" yaml:"queryStringWhitelist"`
}


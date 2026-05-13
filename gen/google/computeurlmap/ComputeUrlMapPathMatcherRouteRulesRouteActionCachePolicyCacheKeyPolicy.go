package computeurlmap


type ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy struct {
	// Names of query string parameters to exclude in cache keys.
	//
	// All other
	// parameters will be included. Either specify excludedQueryParameters
	// or includedQueryParameters, not both. '&' and '=' will be percent
	// encoded and not treated as delimiters. Note: This field applies to
	// routes that use backend services. Attempting to set it on a route that
	// points exclusively to Backend Buckets will result in a configuration
	// error. For routes that point to a Backend Bucket, use
	// includedQueryParameters to define which parameters should be part of
	// the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#excluded_query_parameters ComputeUrlMap#excluded_query_parameters}
	ExcludedQueryParameters *[]*string `field:"optional" json:"excludedQueryParameters" yaml:"excludedQueryParameters"`
	// Allows HTTP cookies (by name) to be used in the cache key.
	//
	// The
	// name=value pair will be used in the cache key Cloud CDN generates.
	// Note: This setting is only applicable to routes that use a Backend
	// Service. It does not affect requests served by a Backend Bucket.
	// Attempting to set it on a route that points exclusively to Backend
	// Buckets will result in a configuration error. Up to 5 cookie names can
	// be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#included_cookie_names ComputeUrlMap#included_cookie_names}
	IncludedCookieNames *[]*string `field:"optional" json:"includedCookieNames" yaml:"includedCookieNames"`
	// Allows HTTP request headers (by name) to be used in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#included_header_names ComputeUrlMap#included_header_names}
	IncludedHeaderNames *[]*string `field:"optional" json:"includedHeaderNames" yaml:"includedHeaderNames"`
	// Names of query string parameters to include in cache keys.
	//
	// All other
	// parameters will be excluded. Either specify includedQueryParameters
	// or excludedQueryParameters, not both. '&' and '=' will be percent
	// encoded and not treated as delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#included_query_parameters ComputeUrlMap#included_query_parameters}
	IncludedQueryParameters *[]*string `field:"optional" json:"includedQueryParameters" yaml:"includedQueryParameters"`
	// If true, requests to different hosts will be cached separately.
	//
	// Note:
	// This setting is only applicable to routes that use a Backend Service.
	// It does not affect requests served by a Backend Bucket, as the host is
	// never included in a Backend Bucket's cache key. Attempting to set it on
	// a route that points exclusively to Backend Buckets will result in a
	// configuration error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#include_host ComputeUrlMap#include_host}
	IncludeHost interface{} `field:"optional" json:"includeHost" yaml:"includeHost"`
	// If true, http and https requests will be cached separately.
	//
	// Note: This
	// setting is only applicable to routes that use a Backend Service. It
	// does not affect requests served by a Backend Bucket, as the protocol is
	// never included in a Backend Bucket's cache key. Attempting to set on a
	// route that points exclusively to Backend Buckets will result in a
	// configuration error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#include_protocol ComputeUrlMap#include_protocol}
	IncludeProtocol interface{} `field:"optional" json:"includeProtocol" yaml:"includeProtocol"`
	// If true, include query string parameters in the cache key according to includedQueryParameters and excludedQueryParameters.
	//
	// If neither is
	// set, the entire query string will be included. If false, the query
	// string will be excluded from the cache key entirely. Note: This field
	// applies to routes that use backend services. Attempting to set it on a
	// route that points exclusively to Backend Buckets will result in a
	// configuration error. For routes that point to a Backend Bucket, use
	// includedQueryParameters to define which parameters should be part of
	// the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_url_map#include_query_string ComputeUrlMap#include_query_string}
	IncludeQueryString interface{} `field:"optional" json:"includeQueryString" yaml:"includeQueryString"`
}


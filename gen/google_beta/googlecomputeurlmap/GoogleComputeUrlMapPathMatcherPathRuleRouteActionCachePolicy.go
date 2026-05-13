package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy struct {
	// Bypass the cache when the specified request headers are matched by name, e.g. Pragma or Authorization headers. Values are case-insensitive. Up to 5 header names can be specified. The cache is bypassed for all cacheMode values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#cache_bypass_request_header_names GoogleComputeUrlMap#cache_bypass_request_header_names}
	CacheBypassRequestHeaderNames *[]*string `field:"optional" json:"cacheBypassRequestHeaderNames" yaml:"cacheBypassRequestHeaderNames"`
	// cache_key_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#cache_key_policy GoogleComputeUrlMap#cache_key_policy}
	CacheKeyPolicy *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy `field:"optional" json:"cacheKeyPolicy" yaml:"cacheKeyPolicy"`
	// Specifies the cache setting for all responses from this route.
	//
	// If not
	// specified, Cloud CDN uses CACHE_ALL_STATIC mode. Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#cache_mode GoogleComputeUrlMap#cache_mode}
	CacheMode *string `field:"optional" json:"cacheMode" yaml:"cacheMode"`
	// client_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#client_ttl GoogleComputeUrlMap#client_ttl}
	ClientTtl *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl `field:"optional" json:"clientTtl" yaml:"clientTtl"`
	// default_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#default_ttl GoogleComputeUrlMap#default_ttl}
	DefaultTtl *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// max_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#max_ttl GoogleComputeUrlMap#max_ttl}
	MaxTtl *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	//
	// This can reduce the
	// load on your origin and improve end-user experience by reducing response
	// latency. When the cacheMode is set to CACHE_ALL_STATIC or
	// USE_ORIGIN_HEADERS, negative caching applies to responses with the
	// specified response code that lack any Cache-Control, Expires, or
	// Pragma: no-cache directives. When the cacheMode is set to
	// FORCE_CACHE_ALL, negative caching applies to all responses with the
	// specified response code, and overrides any caching headers. By default,
	// Cloud CDN applies the following TTLs to these HTTP status codes:
	// * 300 (Multiple Choice), 301, 308 (Permanent Redirects): 10m
	// * 404 (Not Found), 410 (Gone), 451 (Unavailable For Legal Reasons): 120s
	// * 405 (Method Not Found), 501 (Not Implemented): 60s
	// These defaults can be overridden in negativeCachingPolicy. If not
	// specified, Cloud CDN applies negative caching by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#negative_caching GoogleComputeUrlMap#negative_caching}
	NegativeCaching interface{} `field:"optional" json:"negativeCaching" yaml:"negativeCaching"`
	// negative_caching_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#negative_caching_policy GoogleComputeUrlMap#negative_caching_policy}
	NegativeCachingPolicy interface{} `field:"optional" json:"negativeCachingPolicy" yaml:"negativeCachingPolicy"`
	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	//
	// If not specified,
	// Cloud CDN applies request coalescing by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#request_coalescing GoogleComputeUrlMap#request_coalescing}
	RequestCoalescing interface{} `field:"optional" json:"requestCoalescing" yaml:"requestCoalescing"`
	// serve_while_stale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#serve_while_stale GoogleComputeUrlMap#serve_while_stale}
	ServeWhileStale *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale `field:"optional" json:"serveWhileStale" yaml:"serveWhileStale"`
}


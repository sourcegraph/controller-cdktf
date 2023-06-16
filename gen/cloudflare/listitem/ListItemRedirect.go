package listitem


type ListItemRedirect struct {
	// The source url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#source_url ListItemA#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// The target url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#target_url ListItemA#target_url}
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#include_subdomains ListItemA#include_subdomains}
	IncludeSubdomains *string `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#preserve_path_suffix ListItemA#preserve_path_suffix}
	PreservePathSuffix *string `field:"optional" json:"preservePathSuffix" yaml:"preservePathSuffix"`
	// Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#preserve_query_string ListItemA#preserve_query_string}
	PreserveQueryString *string `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// The status code to be used when redirecting a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#status_code ListItemA#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list_item#subpath_matching ListItemA#subpath_matching}
	SubpathMatching *string `field:"optional" json:"subpathMatching" yaml:"subpathMatching"`
}


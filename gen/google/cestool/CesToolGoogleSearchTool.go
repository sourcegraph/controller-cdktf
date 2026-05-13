package cestool


type CesToolGoogleSearchTool struct {
	// The name of the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#name CesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Content will be fetched directly from these URLs for context and grounding.
	//
	// More details: https://cloud.google.com/vertex-ai/generative-ai/docs/url-context.
	// Example: "https://example.com/path.html". A maximum of 20 URLs are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#context_urls CesTool#context_urls}
	ContextUrls *[]*string `field:"optional" json:"contextUrls" yaml:"contextUrls"`
	// Description of the tool's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#description CesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of domains to be excluded from the search results. Example: "example.com". A maximum of 2000 domains can be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#exclude_domains CesTool#exclude_domains}
	ExcludeDomains *[]*string `field:"optional" json:"excludeDomains" yaml:"excludeDomains"`
	// Specifies domain names to guide the search.
	//
	// The model will be instructed to prioritize these domains
	// when formulating queries for google search.
	// This is a best-effort hint and these domains may or may
	// not be exclusively reflected in the final search results.
	// Example: "example.com", "another.site".
	// A maximum of 20 domains can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#preferred_domains CesTool#preferred_domains}
	PreferredDomains *[]*string `field:"optional" json:"preferredDomains" yaml:"preferredDomains"`
}


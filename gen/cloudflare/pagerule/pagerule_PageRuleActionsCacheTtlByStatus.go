package pagerule


type PageRuleActionsCacheTtlByStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#codes PageRule#codes}.
	Codes *string `field:"required" json:"codes" yaml:"codes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#ttl PageRule#ttl}.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}


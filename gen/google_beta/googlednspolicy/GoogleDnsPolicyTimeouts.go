package googlednspolicy


type GoogleDnsPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#create GoogleDnsPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#delete GoogleDnsPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_policy#update GoogleDnsPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


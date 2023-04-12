package googlednsresponsepolicy


type GoogleDnsResponsePolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_response_policy#create GoogleDnsResponsePolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_response_policy#delete GoogleDnsResponsePolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_response_policy#update GoogleDnsResponsePolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


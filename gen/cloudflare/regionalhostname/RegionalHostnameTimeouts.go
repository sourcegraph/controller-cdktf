package regionalhostname


type RegionalHostnameTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#create RegionalHostname#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#update RegionalHostname#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


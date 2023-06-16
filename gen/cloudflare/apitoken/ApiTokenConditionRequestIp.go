package apitoken


type ApiTokenConditionRequestIp struct {
	// List of IP addresses or CIDR notation where the token may be used from.
	//
	// If not specified, the token will be valid for all IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/api_token#in ApiToken#in}
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// List of IP addresses or CIDR notation where the token should not be used from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/api_token#not_in ApiToken#not_in}
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}


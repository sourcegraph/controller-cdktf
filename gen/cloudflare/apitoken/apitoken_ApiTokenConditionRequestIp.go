package apitoken


type ApiTokenConditionRequestIp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#in ApiToken#in}.
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#not_in ApiToken#not_in}.
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}


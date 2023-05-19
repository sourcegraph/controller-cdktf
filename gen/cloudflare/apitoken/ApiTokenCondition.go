package apitoken


type ApiTokenCondition struct {
	// request_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/api_token#request_ip ApiToken#request_ip}
	RequestIp *ApiTokenConditionRequestIp `field:"optional" json:"requestIp" yaml:"requestIp"`
}


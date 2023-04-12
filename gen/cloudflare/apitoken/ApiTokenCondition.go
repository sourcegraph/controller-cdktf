package apitoken


type ApiTokenCondition struct {
	// request_ip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#request_ip ApiToken#request_ip}
	RequestIp *ApiTokenConditionRequestIp `field:"optional" json:"requestIp" yaml:"requestIp"`
}


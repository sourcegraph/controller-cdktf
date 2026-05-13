package apigeesecurityaction


type ApigeeSecurityActionDeny struct {
	// The HTTP response code if the Action = DENY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_security_action#response_code ApigeeSecurityAction#response_code}
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
}


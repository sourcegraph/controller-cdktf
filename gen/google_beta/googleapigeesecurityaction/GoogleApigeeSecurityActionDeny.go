package googleapigeesecurityaction


type GoogleApigeeSecurityActionDeny struct {
	// The HTTP response code if the Action = DENY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_action#response_code GoogleApigeeSecurityAction#response_code}
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
}


package googleiamworkforcepool


type GoogleIamWorkforcePoolAccessRestrictions struct {
	// allowed_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_workforce_pool#allowed_services GoogleIamWorkforcePool#allowed_services}
	AllowedServices interface{} `field:"optional" json:"allowedServices" yaml:"allowedServices"`
	// Disable programmatic sign-in by disabling token issue via the Security Token API endpoint. See [Security Token Service API](https://cloud.google.com/iam/docs/reference/sts/rest).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_workforce_pool#disable_programmatic_signin GoogleIamWorkforcePool#disable_programmatic_signin}
	DisableProgrammaticSignin interface{} `field:"optional" json:"disableProgrammaticSignin" yaml:"disableProgrammaticSignin"`
}


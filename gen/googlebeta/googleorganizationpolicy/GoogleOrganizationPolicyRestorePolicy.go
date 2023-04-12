package googleorganizationpolicy


type GoogleOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_policy#default GoogleOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}


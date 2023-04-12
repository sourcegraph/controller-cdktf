package googleprojectorganizationpolicy


type GoogleProjectOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_organization_policy#default GoogleProjectOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}


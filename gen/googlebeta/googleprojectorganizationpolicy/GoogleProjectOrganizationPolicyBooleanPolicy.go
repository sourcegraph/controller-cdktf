package googleprojectorganizationpolicy


type GoogleProjectOrganizationPolicyBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_organization_policy#enforced GoogleProjectOrganizationPolicy#enforced}
	Enforced interface{} `field:"required" json:"enforced" yaml:"enforced"`
}


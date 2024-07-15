package googlefolderorganizationpolicy


type GoogleFolderOrganizationPolicyBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_organization_policy#enforced GoogleFolderOrganizationPolicy#enforced}
	Enforced interface{} `field:"required" json:"enforced" yaml:"enforced"`
}


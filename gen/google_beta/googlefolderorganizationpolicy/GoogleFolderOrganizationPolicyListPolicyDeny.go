package googlefolderorganizationpolicy


type GoogleFolderOrganizationPolicyListPolicyDeny struct {
	// The policy allows or denies all values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_organization_policy#all GoogleFolderOrganizationPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The policy can define specific values that are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_organization_policy#values GoogleFolderOrganizationPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


package googlefolderorganizationpolicy


type GoogleFolderOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_organization_policy#default GoogleFolderOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

